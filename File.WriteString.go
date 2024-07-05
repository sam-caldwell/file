package file

import (
	"fmt"
	"github.com/sam-caldwell/errors"
)

// WriteString - Write string to the file
//
//	The goal is to create an abstraction layer so changes to golang will not require
//	significant effort to update all projects.
func (fp *File) WriteString(content string) (int, error) {
	if fp.handle == nil {
		return 0, fmt.Errorf(errors.NotInitialized)
	}
	n, err := fp.handle.WriteString(content)
	return n, err
}
