package file

import (
	"fmt"
	"github.com/sam-caldwell/errors"
	"testing"
)

func TestFileValid(t *testing.T) {

	testFunc := func(index int, name string, expectedError error) {
		var f File
		err := f.valid(&name)

		if expectedError == nil {
			if err != nil {
				t.Fatalf("%d expected no error but got '%v' on '%s'", index, err, name)
			}
		} else {
			if err == nil {
				t.Fatalf("%d expected error (%v) but got nil on %s", index, expectedError, name)
			} else {
				if err.Error() != expectedError.Error() {
					t.Fatalf("error mismatch. expected '%v' got '%v'", expectedError.Error(), err.Error())
				}
			}
		}
	}

	type TestData struct {
		name string
		err  error
	}
	testData := []TestData{
		{"myFile.txt", nil},
		{"./myFile.txt", nil},
		{"/tmp/myFile.txt", nil},
		{"/tmp/subdir1/myFile.txt", nil},
		{"/tmp/subdir1/subdir2/myFile.txt", nil},
		{"ftp://myserver.myDomain.tld:/tmp/subdir1/subdir2/myFile.txt", nil},
		{"http://myserver.myDomain.tld:/tmp/subdir1/subdir2/myFile.txt", nil},
		{"https://myserver.myDomain.tld:/tmp/subdir1/subdir2/myFile.txt", nil},
		{"nfs://myserver.myDomain.tld:/tmp/subdir1/subdir2/myFile.txt", nil},
		{"cifs://myserver.myDomain.tld:/tmp/subdir1/subdir2/myFile.txt", nil},
		{"http://myserver.myDomain.tld:/tmp/../../subdir1/subdir2/myFile.txt", fmt.Errorf(errors.InvalidFilePath)},
		{"https://myserver.myDomain.tld:/tmp/subdir1/subdir2/../../myFile.txt", fmt.Errorf(errors.InvalidFilePath)},
		{"", fmt.Errorf(errors.InvalidFilePath)},
		{"../myTraversal1.txt", fmt.Errorf(errors.InvalidFilePath)},
		{"/tmp/../myTraversal1.txt", fmt.Errorf(errors.InvalidFilePath)},
		{"/tmp/../../myTraversal1.txt", fmt.Errorf(errors.InvalidFilePath)},
		{"subdir1/../myTraversal*.txt", fmt.Errorf(errors.InvalidFilePath)},
		{"../../../../../etc/passwd", fmt.Errorf(errors.InvalidFilePath)},
		{"../../../../../etc/passwd*", fmt.Errorf(errors.InvalidFilePath)},
	}

	for n, row := range testData {
		testFunc(n, row.name, row.err)
	}
}
