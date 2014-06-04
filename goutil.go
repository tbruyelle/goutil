package goutil

import (
	"os"
)

// FileExists returns true if the file exists.
func FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
