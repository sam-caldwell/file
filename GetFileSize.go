package file

import (
	"github.com/sam-caldwell/ansi"
	"os"
)

// GetFileSize - Return the int64 file size of a given file.
func GetFileSize(f *os.File) int64 {
	fileInfo, err := f.Stat()
	if err != nil {
		ansi.Red().Println(err.Error()).Fatal(1)
	}
	return fileInfo.Size()
}
