package file

import (
	"crypto/rand"
	"fmt"
	"os"
	"strings"
)

// CreateRandomImage - creates a file at the specified path with the specified size filled with random data
//
//	     func main() {
//		        // Example usage of createRandomImage function
//	   	if err:=createRandomImage("random.img", 10*size.GB);err!=nil {
//	               ansi.Red().Println(err.Error()).Fatal(exit.GeneralError).Reset()
//	           }
//	     }
func CreateRandomImage(path string, size uint64) (err error) {
	const (
		bufferSize uint64 = 4096
	)
	if strings.TrimSpace(path) == "" {
		return fmt.Errorf("path/file cannot be empty")
	}
	var file *os.File
	if file, err = os.Create(path); err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	buffer := make([]byte, bufferSize)
	for written := uint64(0); written < size; {
		bytesToWrite := bufferSize
		if remaining := size - written; remaining < bufferSize {
			bytesToWrite = remaining
		}
		if _, err = rand.Read(buffer[:bytesToWrite]); err != nil {
			return fmt.Errorf("failed to generate random data: %v", err)
		}
		if _, err = file.Write(buffer[:bytesToWrite]); err != nil {
			return fmt.Errorf("failed to write to file: %v", err)
		}
		written += bytesToWrite
	}
	return err
}
