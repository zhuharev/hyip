package boltutils

import (
	"os"

	"github.com/boltdb/bolt"
)

type Opt func(db *DB) error

func Compression(ct CompressorType) func(db *DB) error {
	return func(db *DB) error {
		switch ct {
		case GzipCompressor:
			db.Compressor = newGzipCompressor()
		case Lz4Compressor:
			db.Compressor = newLz4Compressor()
		}
		return nil
	}
}

func BoltDB(bdb *bolt.DB) Opt {
	return func(db *DB) error {
		db.DB = bdb
		return nil
	}
}

// Open open database file and return pointer of DB
func Open(path string, mode os.FileMode, options *bolt.Options) Opt {
	return func(db *DB) error {
		bdb, err := bolt.Open(path, mode, options)
		if err != nil {
			return err
		}
		db.DB = bdb
		return nil
	}
}

// OpenPath shortland for Open with single arg
func OpenPath(path string) Opt {
	return Open(path, 0777, nil)
}
