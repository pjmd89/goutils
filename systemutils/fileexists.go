package systemutils

import (
	"os"
	"reflect"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if err == nil {
		return true
	}
	x := reflect.TypeOf(info)
	if x != nil {
		return !info.IsDir()
	}
	return false
}
