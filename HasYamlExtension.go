package file

import "github.com/sam-caldwell/extensions"

// HasYamlExtension - return boolean result if a YAML extension exists
func HasYamlExtension(filename string) bool {
	return extensions.HasOneOf(filename, extensions.YAML, extensions.YML)
}
