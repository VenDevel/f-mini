package utils

import (
	"io/ioutil"
	"os"
)

func init() {

}

func ReadFileAllByPath(path string) ([]byte, error) {
	fi, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return fd, err
}
