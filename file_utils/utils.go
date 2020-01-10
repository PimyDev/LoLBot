package file_utils

import "os"

func FileExists(filename string) bool {
	exists, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !exists.IsDir()
}
