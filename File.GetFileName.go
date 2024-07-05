package file

// GetFileName - Return the internal file/path state
//
//	The goal is to create an abstraction layer so changes to golang will not require
//	significant effort to update all projects.
func (fp *File) GetFileName() string {
	if fp.handle == nil {
		return ""
	}
	fp.lock.Lock()
	defer fp.lock.Unlock()
	return fp.handle.Name()
}
