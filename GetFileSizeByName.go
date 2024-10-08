package file

import (
	"github.com/sam-caldwell/ansi"
	"os"
)

// GetFileSizeByName - Return the int64 file size of a given file (by filename).
func GetFileSizeByName(f string) int64 {
	handle, err := os.Open(f)
	if err != nil {
		ansi.Red().Println(err.Error()).Fatal(1)
	}
	defer func() { _ = handle.Close() }()
	return GetFileSize(handle)
}
