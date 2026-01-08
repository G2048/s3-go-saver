package synchronizer

import (
	"fmt"
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	fileName1 := "file_test_1.txt"
	fileName2 := "file_test_2.txt"
	file1 := &File{Name: fileName1}
	file2 := &File{Name: fileName2}
	CreateFile(file1)
	CreateFile(file2)
	file1.Open()
	file2.Open()
	defer file1.Close()
	defer file2.Close()
	defer os.Remove(fileName1)
	defer os.Remove(fileName2)

	fmt.Println("Write info to file1")
	file1.Write([]byte("Record1: Hello World in file1!\n"))
	file1.Write([]byte("Record2: Hello World in file1!\n"))

	fmt.Println("Write info to file2")
	file2.Write([]byte("Record1: Hello World in file2!\n"))
	file2.Write([]byte("Record2: Hello World in file2!\n"))

	fileBody1 := file1.ReadFull()
	fmt.Printf("\nFile body:\n %s\n", fileBody1)

	fileBody2 := file2.ReadFull()
	fmt.Printf("\nFile body:\n %s\n", fileBody2)
}
