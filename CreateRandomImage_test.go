package file

import (
	"github.com/sam-caldwell/size"
	"os"
	"testing"
)

func TestCreateRandomImage(t *testing.T) {
	t.Run("Test nil file/path", func(t *testing.T) {
		if err := CreateRandomImage("", 10); err == nil {
			t.Fatalf("CreateRandomImage() expected error = %v", err)
		} else {
			if err.Error() != "path/file cannot be empty" {
				t.Fatalf("unexpected error: %v", err)
			}
		}
	})
	t.Run("File creation tests", func(t *testing.T) {
		tests := []struct {
			name     string
			size     uint64
			expected int64
		}{
			{"0 file", 0, 0},
			{"1KB file", 1 * size.KB, 1 * size.KB},
			{"1MB file", 1 * size.MB, 1 * size.MB},
			{"10MB file", 10 * size.MB, 10 * size.MB},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				tmpFile, err := os.CreateTemp("", "random_img_*.img")
				if err != nil {
					t.Fatalf("Failed to create temp file: %v", err)
				}
				defer func() {
					if err := os.Remove(tmpFile.Name()); err != nil {
						t.Fatalf("failed to remove test file (%s)", tmpFile.Name())
					}
				}()

				if err := CreateRandomImage(tmpFile.Name(), tt.size); err != nil {
					t.Fatalf("CreateRandomImage() error = %v", err)
				}

				fileInfo, err := os.Stat(tmpFile.Name())
				if err != nil {
					t.Fatalf("Failed to stat file: %v", err)
				}
				if fileInfo.Size() != tt.expected {
					t.Errorf("Expected file size %d, got %d", tt.expected, fileInfo.Size())
				}
			})
		}
	})

}
