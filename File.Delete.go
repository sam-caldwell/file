package file

import (
	"fmt"
	"github.com/sam-caldwell/errors"
	"os"
)

// Delete - Delete the current file (if set)
//
//	The goal is to create an abstraction layer so changes to golang will not require
//	significant effort to update all projects.
func (fp *File) Delete() error {
	if fp.handle == nil {
		return fmt.Errorf(errors.NotInitialized)
	}
	fp.lock.Lock()
	defer fp.lock.Unlock()
	name := fp.handle.Name()
	_ = fp.handle.Close()
	return os.Remove(name)
}
