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
	name string
	body []byte
	fd   *os.File
	fi   *os.FileInfo
}

func (f File) Name() string {
	return f.name
}
func (f *File) Body() []byte {
	return f.body
}
func (f *File) Size() int64 {
	return int64(len(f.body))
}
func (f *File) Open() (*os.File, error) {
	var err error
	f.fd, err = os.OpenFile(f.name, os.O_RDWR|os.O_CREATE, 0777)
	fi, err := os.Stat(f.name)
	f.fi = &fi
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
	return n, err
}
func (f *File) TimeCreated() int64 {
	return (*f.fi).ModTime().Unix()
}

func CreateFile(f Opener) error {
	_, err := f.Open()
	defer f.Close()
	if err != nil {
		panic(err)
	}
	return nil
}
