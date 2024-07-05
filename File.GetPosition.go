package file

import (
	"fmt"
	"github.com/sam-caldwell/errors"
)

// GetPosition - Return current file position (in bytes from origin)
//
//	The goal is to create an abstraction layer so changes to golang will not require
//	significant effort to update all projects.
func (fp *File) GetPosition() (pos int64, err error) {
	if fp.handle == nil {
		return 0, fmt.Errorf(errors.NotInitialized)
	}
	fp.lock.Lock()
	defer fp.lock.Unlock()
	pos, err = fp.handle.Seek(0, SeekFromCurrent)
	return pos, err
}
