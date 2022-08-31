package gsftp

import (
	"bytes"
	"io"
	"io/ioutil"
)

// NewReadAtBuffer creates a new reader
func NewReadAtBuffer(r io.ReadCloser) (io.ReaderAt, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = r.Close()
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(buf), nil
}
