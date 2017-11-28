// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"encoding/json"

	"github.com/zhuharev/boltutils"
	"github.com/zhuharev/hyip/pkg/setting"
)

var (
	boltDB *boltutils.DB

	kvBucketName = []byte(`kv`)
)

func initBolt() (err error) {

	if setting.Dev {
		boltDB, err = boltutils.New(boltutils.OpenPath("data/kv.bolt"))
		if err != nil {
			return err
		}
	} else {
		boltDB, err = boltutils.New(boltutils.OpenPath("/storage/kv.bolt"))
		if err != nil {
			return err
		}
	}

	err = boltDB.CreateBucket(kvBucketName)
	if err != nil {
		return err
	}

	return nil
}

func marshalValue(value interface{}) ([]byte, error) {
	switch val := value.(type) {
	case []byte:
		return val, nil
	default:
		return json.Marshal(value)
	}
}

func SaveValue(key string, value interface{}) error {
	bts, err := marshalValue(value)
	if err != nil {
		return err
	}
	return boltDB.Put(kvBucketName, []byte(key), bts)
}

func GetValue(key string) ([]byte, error) {
	return boltDB.Get(kvBucketName, []byte(key))
}

func UnmarshalValue(key string, res interface{}) error {
	bts, err := GetValue(key)
	if err != nil {
		return err
	}
	return json.Unmarshal(bts, res)
}
