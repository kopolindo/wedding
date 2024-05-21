package log

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2/log"
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

// runningInDocker checks if application is running in docker container
// returns true if running in docker, false otherwise
func runningInDocker() bool {
	file, err := os.Open("/proc/1/cgroup")
	if err != nil {
		log.Errorf("Error: %s", err.Error())
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "docker") || strings.Contains(line, "docker-") {
			return true
		}
	}

	return false
}
