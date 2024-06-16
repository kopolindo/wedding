package log

import (
	"fmt"
	"os"
	"strings"
)

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

// runningInDocker checks if the system is running inside a Docker build environment.
// returns true if running in docker, false otherwise
func runningInDocker() bool {
	container := os.Getenv("CONTAINER")
	return strings.ToLower(container) == "true"
}
