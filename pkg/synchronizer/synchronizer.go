package synchronizer

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log/slog"
)

func (f *File) Hash() (string, error) {
	if f.Body == nil {
		panic("File.Body is nil!")
	}
	return fmt.Sprintf("%x", sha256.Sum256(f.Body)), nil
}
func FillBody(f *File) {
	f.Body = f.ReadFull()
}
func (f *File) ReadFull() (fileBody []byte) {
	var buff []byte = make([]byte, 10)
	for {
		n, _ := io.ReadFull(f, buff)
		// fmt.Printf("Num: %d buff: %s\n", n, buff)
		if n == 0 {
			break
		}
		fileBody = append(fileBody, buff...)
	}
	return fileBody
}

func SameFile(f1, f2 *File) bool {
	hash1, err1 := f1.Hash()
	hash2, err2 := f2.Hash()
	if err1 != nil || err2 != nil {
		slog.Warn("Error by hashing file: %s %s", f1.Name, err1.Error())
		slog.Warn("Error by hashing file: %s %s", f2.Name, err2.Error())
		return false
	}
	return hash1 == hash2
}

func Synchronize() {
	// file1 := "file1.txt"
	// file2 := "file2.txt"
	// fmt.Printf("file1: %v\n", file1)
}
