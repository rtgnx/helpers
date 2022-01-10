package helpers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

func JSONRead(r io.Reader, v interface{}) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, v)
}

func JSONWrite(w io.Writer, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	_, err = w.Write(b)
	return err
}

func JSONFromFile(filePath string, v interface{}) error {
	fd, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer fd.Close()

	return JSONRead(fd, v)
}

func JSONToFile(filePath string, v interface{}, perm os.FileMode) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, b, perm)
}
