package file

// Exists - Return boolean if file exists
//
//	The goal is to create an abstraction layer so changes to golang will not require
//	significant effort to update all projects.
func (fp *File) Exists(fileName string) bool {
	return fp.Existsp(&fileName)
}
