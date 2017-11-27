# Bolt utils [![Build Status](https://travis-ci.org/zhuharev/boltutils.svg?branch=master)](https://travis-ci.org/zhuharev/boltutils) [![Go Report Card](https://goreportcard.com/badge/github.com/zhuharev/boltutils)](https://goreportcard.com/report/github.com/zhuharev/boltutils) [![Coverage Status](https://coveralls.io/repos/github/zhuharev/boltutils/badge.svg?branch=master)](https://coveralls.io/github/zhuharev/boltutils?branch=master) [![GoDoc](https://godoc.org/github.com/zhuharev/boltutils?status.svg)](http://godoc.org/github.com/zhuharev/boltutils)

Bolt db wrapper for facilitating the use it low-level API.

## Usage

```go
import (
  "github.com/zhuharev/boltutils"
)

...

db, _ := boltutils.Open("path/to/db.bolt", 0666, nil)
db.Put([]byte("key"), []byte("value"))
value, _ := db.Get([]byte("key"))
// value = []byte("value")
```

## gzip in the box

```go
db.PutGzip([]byte("key"), []byte("value"))
value, _ := db.GetGzipped([]byte("key"))
// value = []byte("value")
```

## json support

```go
db.PutJSON(anyStruct)
err = db.GetJSON([]byte("bucket"), []byte("key"), &anyStruct)
```
