package file

import (
	"github.com/sam-caldwell/extensions"
)

// GetExtension - return the given file's extension
func GetExtension(name string) string {
	return extensions.Get(name)
}
