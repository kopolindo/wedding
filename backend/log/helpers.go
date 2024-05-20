package log

import (
	"fmt"
	"os"
)

// dirExists checks if directory exists, return true if exists, false otherwise
func dirExists(dirname string) bool {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil && info.IsDir()
}

// ensureDir checks if directory exists, if not it creates it
func ensureDir(dirname string) error {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		err := os.MkdirAll(dirname, os.ModePerm)
		if err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("%s already exists and is not a directory", dirname)
	}
	return nil
}
