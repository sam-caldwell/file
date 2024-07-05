package file

import (
	"os"
	"sync"
)

// File - This is a wrapper around file-access logic.
//
//	The goal is to create an abstraction layer so changes to golang will not require
//	significant effort to update all projects.
type File struct {
	lock   sync.Mutex
	handle *os.File
}
