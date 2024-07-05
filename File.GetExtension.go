package file

import (
	"github.com/sam-caldwell/extensions"
)

// GetExtension - return the current file's extension
//
//		 The goal is to create an abstraction layer so changes to golang will not require
//		 significant effort to update all projects.
//
//	     - returns empty string if not initialized
func (fp *File) GetExtension() string {
	if fp.handle == nil {
		return ""
	}
	return extensions.Get(fp.handle.Name())
}
