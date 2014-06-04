package goutil

import (
	"os"
)

// IsFile returns true if the file exists and is not a directory.
func IsFile(path string) bool {
	exists, fi := fileExists(path)
	if exists && !fi.IsDir() {
		return true
	}
	return false
}

// FileExists returns true if the file exists.
func FileExists(path string) bool {
	exists, _ := fileExists(path)
	return exists
}

func fileExists(path string) (bool, os.FileInfo) {
	fi, err := os.Stat(path)
	os.IsNotExist(err)
	if err != nil && os.IsNotExist(err) {
		return false, nil
	}
	return true, fi
}
