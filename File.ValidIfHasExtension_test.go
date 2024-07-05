package file

import (
	"fmt"
	"github.com/sam-caldwell/errors/v2"
	"strings"
	"testing"
)

func TestFile_SetWithExtension(t *testing.T) {

	type TestData struct {
		fileName   string
		extensions []string
		error      error
	}
	testData := []TestData{
		{"testFile.iso", []string{".iso"}, nil},
		{"testFile.iso", []string{".iso", ".img", ".disk"}, nil},
		{"testFile.img", []string{".iso", ".img", ".disk"}, nil},
		{"testFile.disk", []string{".iso", ".img", ".disk"}, nil},
		{"testFile.pdf", []string{".pdf", ".doc"}, nil},
		{"testFile.doc", []string{".pdf", ".doc"}, nil},
		{"testFile.iso", []string{".pdf", ".doc"}, fmt.Errorf(errors.MissingExtension)},
		{"testFile", []string{".img", ".iso"}, fmt.Errorf(errors.MissingExtension)},
		{"testFile", []string{".pdf", ".doc"}, fmt.Errorf(errors.MissingExtension)},
		{"test", []string{".iso"}, fmt.Errorf(errors.MissingExtension)},
	}

	testFunc := func(index int, fileName string, extensions []string, expectedError error) {

		t.Run(fmt.Sprintf("test %d (name:%s, extensions:[%v],expectedError:%v)",
			index, fileName, strings.Join(extensions, ","), expectedError), func(t *testing.T) {

			var f File

			t.Logf("test with '%s' and %v ", fileName, extensions)

			err := f.ValidIfHasExtension(fileName, extensions)

			t.Logf("test with '%s' and %v err: %v", fileName, extensions, err)

			if expectedError == nil {
				if err != nil {
					t.Fatalf("%d error mismatch. file '%s' err '%v' expect no error", index, fileName, err)
				} else {
					return
				}
			}
			if (err != nil) && (err.Error() != expectedError.Error()) {
				t.Fatalf("%d error mismatch. file '%s' err '%v' expectedError '%v'",
					index, fileName, err, expectedError)
			}
			if err == nil {
				t.Fatalf("%d expected error not found. file '%s' expectedError '%v'",
					index, fileName, expectedError)
			}
		})
	}

	//Run all tests
	for index, row := range testData {
		testFunc(index, row.fileName, row.extensions, row.error)
	}
}
