package boltutils

import (
	"bytes"
	"io/ioutil"

	"github.com/pierrec/lz4"
)

type Compressor interface {
	Compress(data []byte) ([]byte, error)
	Decompress(data []byte) ([]byte, error)
}

type CompressorType int

const (
	NoopCompressor CompressorType = iota
	GzipCompressor
	Lz4Compressor
)

type gzipCompressor struct {
}

func newGzipCompressor() gzipCompressor {
	return gzipCompressor{}
}

func (gc gzipCompressor) Compress(data []byte) ([]byte, error) {
	return gzipData(data)
}

func (gc gzipCompressor) Decompress(data []byte) ([]byte, error) {
	return ungzipData(data)
}

type lz4Compressor struct {
}

func newLz4Compressor() lz4Compressor {
	return lz4Compressor{}
}

func (gc lz4Compressor) Compress(data []byte) ([]byte, error) {
	var b bytes.Buffer
	w := lz4.NewWriter(&b)
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

func (gc lz4Compressor) Decompress(compressedData []byte) ([]byte, error) {
	r := lz4.NewReader(bytes.NewReader(compressedData))
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return data, nil
}
