package boltutils

import (
	"bytes"
	"fmt"

	"github.com/boltdb/bolt"
)

var (
	// ErrNotFound not found error
	ErrNotFound = fmt.Errorf("not found")
	// ErrBreak break's iterator
	ErrBreak = fmt.Errorf("break")
)

func makePutFunc(bucketName interface{}, key, value []byte) func(tx *bolt.Tx) error {
	switch typ := bucketName.(type) {
	case [][]byte:
		return makePutPathFunc(typ, key, value)
	case []byte:
		return func(tx *bolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists(typ)
			if err != nil {
				return err
			}
			return bucket.Put(key, value)
		}
	}
	panic("unknow path format")
}

// Put value into database
func Put(db *bolt.DB, bucketName interface{}, key, value []byte) error {
	return db.Update(makePutFunc(bucketName, key, value))
}

func makePutPathFunc(bucketNames [][]byte, key, value []byte) func(tx *bolt.Tx) error {
	if len(bucketNames) == 1 {
		return makePutFunc(bucketNames[0], key, value)
	}
	return func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucketNames[0])
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}
		if err != nil {
			return err
		}
		for _, path := range bucketNames[1:] {
			bucket, err = bucket.CreateBucketIfNotExists(path)
			if bucket == nil {
				return bolt.ErrBucketNotFound
			}
			if err != nil {
				return err
			}
		}
		return bucket.Put(key, value)
	}
}

// PutPath value into database
func PutPath(db *bolt.DB, bucketNames [][]byte, key, value []byte) error {
	return db.Update(makePutPathFunc(bucketNames, key, value))
}

func makeGetFunc(bucketName, key []byte, target *[]byte) func(tx *bolt.Tx) error {
	return func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}
		bts := bucket.Get(key)
		if bts == nil {
			return ErrNotFound
		}
		*target = bts
		return nil
	}
}

// Get value from database
func Get(db *bolt.DB, bucketName, key []byte) ([]byte, error) {
	var res []byte
	err := db.View(makeGetFunc(bucketName, key, &res))
	return res, err
}

func makeGetPathFunc(bucketNames [][]byte, key []byte, target *[]byte) func(tx *bolt.Tx) error {
	if len(bucketNames) == 1 {
		return makeGetFunc(bucketNames[0], key, target)
	}
	return func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketNames[0])
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}
		for _, bucketName := range bucketNames[1:] {
			bucket = bucket.Bucket(bucketName)
			if bucket == nil {
				return bolt.ErrBucketNotFound
			}
		}
		bts := bucket.Get(key)
		if bts == nil {
			return ErrNotFound
		}
		*target = bts
		return nil
	}
}

// GetPath value from database bucket/subbucket/.../valuebucket
func GetPath(db *bolt.DB, bucketNames [][]byte, key []byte) ([]byte, error) {
	var res []byte
	err := db.View(makeGetPathFunc(bucketNames, key, &res))
	return res, err
}

func makeIterateFunc(bucketName []byte, fn func(k, v []byte) error, compressors ...Compressor) func(tx *bolt.Tx) error {
	return func(tx *bolt.Tx) (err error) {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}

		c := bucket.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			if len(compressors) > 0 && compressors[0] != nil && len(v) > 0 {
				v, err = compressors[0].Decompress(v)
				if err != nil {
					return err
				}
			}
			err = fn(k, v)
			if err != nil {
				if err == ErrBreak {
					return nil
				}
				return err
			}
		}

		return nil
	}
}

// Iterate over database
func Iterate(db *bolt.DB, bucketName []byte, fn func(k, v []byte) error, compressors ...Compressor) error {
	return db.View(makeIterateFunc(bucketName, fn, compressors...))
}

func makeIteratePrefixFunc(bucketName, prefix []byte, fn func(k, v []byte) error) func(tx *bolt.Tx) error {
	return func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}

		c := bucket.Cursor()

		for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
			err := fn(k, v)
			if err != nil {
				if err == ErrBreak {
					return nil
				}
				return err
			}
		}

		return nil
	}
}

// IteratePrefix over database
func IteratePrefix(db *bolt.DB, bucketName, prefix []byte, fn func(k, v []byte) error) error {
	return db.View(makeIteratePrefixFunc(bucketName, prefix, fn))
}

func makeCreateBucketFunc(bucketName []byte) func(tx *bolt.Tx) error {
	return func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return err
		}
		return nil
	}
}

// CreateBucket create bucket if not exists
func CreateBucket(db *bolt.DB, bucketName []byte) error {
	return db.Update(makeCreateBucketFunc(bucketName))
}

func makeCreateBucketPathFunc(bucketPath [][]byte) func(tx *bolt.Tx) error {
	if len(bucketPath) == 1 {
		return makeCreateBucketFunc(bucketPath[0])
	}
	return func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucketPath[0])
		if err != nil {
			return err
		}
		for _, path := range bucketPath[1:] {
			b, err = b.CreateBucketIfNotExists(path)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func CreateBucketPath(db *bolt.DB, bucketPath [][]byte) error {
	return db.Update(makeCreateBucketPathFunc(bucketPath))
}

// DB wrapper for bolt.DB
type DB struct {
	*bolt.DB
	Compressor Compressor
}

// New return pointer of DB
func New(opts ...Opt) (db *DB, err error) {
	db = new(DB)
	for _, opt := range opts {
		err = opt(db)
		if err != nil {
			return
		}
	}
	return
}

// Put value into db
func (db *DB) Put(bucketName interface{}, key, value []byte) (err error) {
	if db.Compressor != nil {
		value, err = db.Compressor.Compress(value)
		if err != nil {
			return
		}
	}
	return Put(db.DB, bucketName, key, value)
}

// PutPath value into db
func (db *DB) PutPath(bucketName [][]byte, key, value []byte) error {
	return PutPath(db.DB, bucketName, key, value)
}

// Get value from db
func (db *DB) Get(bucketName, key []byte) ([]byte, error) {
	data, err := Get(db.DB, bucketName, key)
	if err != nil {
		return nil, err
	}

	if db.Compressor != nil {
		data, err = db.Compressor.Decompress(data)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

// GetPath value from db
func (db *DB) GetPath(bucketName [][]byte, key []byte) ([]byte, error) {
	return GetPath(db.DB, bucketName, key)
}

// Iterate over db
func (db *DB) Iterate(bucketName []byte, fn func(k, v []byte) error) error {
	return Iterate(db.DB, bucketName, fn, db.Compressor)
}

// IteratePrefix over db
func (db *DB) IteratePrefix(bucketName, prefix []byte, fn func(k, v []byte) error) error {
	return IteratePrefix(db.DB, bucketName, prefix, fn)
}

// CreateBucket create buckt if it not exists
func (db *DB) CreateBucket(bucketName []byte) error {
	return CreateBucket(db.DB, bucketName)
}

func (db *DB) CreateBucketPath(bucketPath [][]byte) error {
	return CreateBucketPath(db.DB, bucketPath)
}
