package file

import (
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFile_GetExtension(t *testing.T) {
	t.Run("test uninitialized state", func(t *testing.T) {
		var testFile File
		testFile.handle = nil
		if testFile.GetExtension() != "" {
			t.Fatalf("file extension should be '' if not initialized")
		}
	})
	t.Run("test with an actual file", func(t *testing.T) {
		var (
			testFile         File
			name             string
			extension        = ".myExtension"
			expectedFileName = filepath.Join("/tmp/", uuid.New().String()) + extension
		)
		t.Run("test initialized state", func(t *testing.T) {
			t.Cleanup(func() {
				_ = testFile.handle.Close()
				_ = os.Remove(name)
			})
			t.Run("test create test file", func(t *testing.T) {
				var err error
				if err = testFile.Open(expectedFileName, os.O_CREATE|os.O_RDWR, 0600); err != nil {
					t.Fatal(err)
				}
				if testFile.handle == nil {
					t.Fatal("handle is nil")
				}
				name = testFile.handle.Name()
				if name == "" {
					t.Fatal("name is empty")
				}
			})
			t.Run("test get the extension", func(t *testing.T) {
				if ext := testFile.GetExtension(); ext != strings.ToLower(extension) {
					t.Fatalf("extension is wrong.  Got: '%s', wanted: '%s'", ext, extension)
				}
			})
		})
	})
}
