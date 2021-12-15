package file

import (
	"os"
)

func MkdirIfNotExist(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		Mkdir(path)
	}
	return err
}

func Mkdir(path string) error {
	return os.Mkdir(path, os.ModePerm)

}
