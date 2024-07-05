package file

import "github.com/sam-caldwell/extensions"

// HasJsonExtension - return boolean result if JSON extension exists (case-insensitive)
func HasJsonExtension(filename string) bool {
	return GetExtension(filename) == extensions.JSON
}
