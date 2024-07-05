package file

import "io"

// GetWriter - Return the file handle as io.Writer
func (fp *File) GetWriter() io.Writer {
	return fp.handle
}
