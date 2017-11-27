package boltutils

import "encoding/json"

func (db *DB) GetJSON(bucketName, key []byte, v interface{}) (err error) {
	var data []byte
	data, err = db.Get(bucketName, key)
	if err != nil {
		return
	}

	return json.Unmarshal(data, v)
}

func (db *DB) PutJSON(bucketName, key []byte, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return db.Put(bucketName, key, data)
}
