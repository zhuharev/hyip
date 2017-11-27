package boltutils

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

func ungzipData(data []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer r.Close()
	data, err = ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func gzipData(data []byte) ([]byte, error) {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	_, err := w.Write(data)
	if err != nil {
		return nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (db *DB) IterateGzipped(bucketName []byte, fn func(k, v []byte) error) error {
	err := db.Iterate(bucketName, func(k, v []byte) error {
		data, err := ungzipData(v)
		if err != nil {
			return err
		}
		return fn(k, data)
	})
	return err
}
