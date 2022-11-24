package systemutils

import (
	"fmt"
	"path/filepath"
	"strings"
)

type FileInfo struct {
	Abs      string
	Dir      string
	BaseName string
	FileName string
	Ext      string
	Exists   bool
}

func NewPath(path string) (*FileInfo, error) {
	fileInfo := &FileInfo{}
	abs, err := filepath.Abs(path)
	dir := filepath.Dir(abs)
	baseName := filepath.Base(abs)
	ext := filepath.Ext(abs)
	filename := strings.TrimSuffix(baseName, ext)
	fileInfo.Abs = abs
	fileInfo.Dir = dir
	fileInfo.BaseName = baseName
	fileInfo.FileName = filename
	fileInfo.Ext = ext
	fileInfo.Exists = FileExists(abs)

	if err != nil || !fileInfo.Exists {
		return fileInfo, fmt.Errorf("file" + abs + " do not exists")
	}
	return fileInfo, nil
}
