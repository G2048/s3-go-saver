package synchronizer

import (
	"errors"
	"io"
	"os"
)

type Opener interface {
	io.ReadWriteCloser
	Open() (*os.File, error)
}

func checkFileIsOpen[T any, R any](fn *func(T) (R, error)) func(T) (R, error) {
	return func(args T) (R, error) {
		result, err := (*fn)(args)
		if errors.Is(err, os.ErrClosed) {
			panic("You Must call the File.Open() first")
		}
		return result, err
	}
}

type File struct {
	Name string
	Body []byte
	fd   *os.File
}

func (f *File) Open() (*os.File, error) {
	var err error
	f.fd, err = os.OpenFile(f.Name, os.O_RDWR|os.O_CREATE, 0777)
	return f.fd, err
}
func (f *File) Close() error {
	return f.fd.Close()
}
func (f *File) Read(p []byte) (n int, err error) {
	ptr := f.fd.Read
	return checkFileIsOpen(&ptr)(p)
}
func (f *File) Write(p []byte) (n int, err error) {
	ptr := f.fd.Write
	n, err = checkFileIsOpen(&ptr)(p)
	err = f.fd.Sync()
	// f.fd.Seek(int64(len(p)), 1)
	return n, err
}

func CreateFile(f Opener) error {
	_, err := f.Open()
	defer f.Close()
	if err != nil {
		panic(err)
	}
	return nil
}
