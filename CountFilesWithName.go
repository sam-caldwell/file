package file

import (
	"os"
	"path/filepath"
)

// CountFilesWithName - Walk a given path and count the number of files found.
//
//	Return an error if file system issues occur.
func CountFilesWithName(path string, fileName string) (int, error) {
	count := 0
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == fileName {
			count++
		}
		return nil
	})
	return count, err
}
