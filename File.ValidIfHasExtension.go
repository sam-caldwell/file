package file

import (
	"fmt"
	"github.com/sam-caldwell/errors"
	"github.com/sam-caldwell/extensions"
)

// ValidIfHasExtension - Return nil if the file name is valid and has an allowed extension.
//
//	Note: any error returned indicates a validation issue.  Nil returns indicate valid input
func (fp *File) ValidIfHasExtension(name string, allowedExtensions []string) error {
	if err := fp.valid(&name); err != nil {
		return err
	}
	if !extensions.HasOneOf(name, allowedExtensions...) {
		return fmt.Errorf(errors.MissingExtension)
	}
	return nil
}
