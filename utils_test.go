package documenttype_test

import (
	"os"
	"errors"
	"fmt"

	"github.com/danieledambrosi-rizzoli/documenttype"
)

type TestFile struct {
	path string
	ext  string
	match  bool
}

type TestByte struct {
	buf []byte
	ext string
	match bool
}

func CheckFromFile(testFiles []TestFile) (bool, error) {
	for _, t := range testFiles {
		buff, err := os.ReadFile(t.path)
		if err != nil {
			return false, err
		}

		typ := documenttype.GetType(buff)
		if (typ.Ext == t.ext) != t.match {
			errstr := fmt.Sprintf("File: %s. Expected %s, found %s", t.path, t.ext, typ.Ext)
			return false, errors.New(errstr)
		}
	}

	return true, nil
}

func CheckFromBuff(testBytes []TestByte) (bool, error) {
	for _, t := range testBytes {
		typ := documenttype.GetType(t.buf)
		if (typ.Ext == t.ext) != t.match {
			errstr := fmt.Sprintf("Buffer: %s. Expected %s, found %s", t.buf, t.ext, typ.Ext)
			return false, errors.New(errstr)
		}
	}
	return true, nil
}