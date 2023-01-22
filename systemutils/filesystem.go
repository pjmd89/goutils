package systemutils

import (
	"io/fs"
	"os"
)

type FSInterface interface {
	Open(name string) (fs.File, error)
	ReadDir(name string) ([]fs.DirEntry, error)
	ReadFile(name string) ([]byte, error)
}

type FS struct {
}

func NewFS() *FS {
	return &FS{}
}
func (o *FS) Open(name string) (r fs.File, err error) {
	r, err = os.Open(name)
	return
}
func (o *FS) ReadDir(name string) (r []fs.DirEntry, err error) {
	r, err = os.ReadDir(name)
	return
}
func (o *FS) ReadFile(name string) (r []byte, err error) {
	r, err = os.ReadFile(name)
	return
}
