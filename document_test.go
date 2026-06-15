package documenttype_test

import (
	"os"
	"testing"

    "path/filepath"
)

func buildFromSamples() ([]TestFile, error) {
	files := make([]TestFile, 0)

	root := "./samples"
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
			return err
        }
		if !info.IsDir() {
        	files = append(files, TestFile{path, filepath.Ext(info.Name())[1:], true})
		}
        return nil
    })

	if err != nil {
		return files, err
	}
	return files, nil
}

func TestAll(t *testing.T) {
	files, err := buildFromSamples()
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	res, err := CheckFromFile(files)
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	if res == false {
		t.Errorf("This should never happen...")
	}
}