package systemutils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pjmd89/goutils/systemutils/debugmode"
)

type Logs struct {
	System *Log
	Access *Log
}
type Log struct {
	path      string
	fileName  string
	extension string
	logger    *log.Logger
}

func NewLog(filePath string) (r *Log) {
	p, err := filepath.Abs(filePath)
	if err != nil {
		log.Fatal("log not loaded: " + err.Error())
	}
	path := filepath.Dir(p)
	fileName := filepath.Base(p)
	extension := filepath.Ext(fileName)
	fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))
	err = os.MkdirAll(path, 0666)
	if err != nil {
		log.Fatal("log not loaded: " + err.Error())
		return nil
	}
	r = &Log{
		path:      path,
		fileName:  fileName,
		extension: extension,
	}
	return
}
func logFile(path string, fileName string, extension string) (r *os.File) {
	if !debugmode.Enabled {
		year, month, day := time.Now().Date()
		filePath := fmt.Sprintf(path+"/"+fileName+"-%v-%v-%v"+extension, day, month.String(), year)
		file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal("log not loaded: " + err.Error())
		}
		r = file
	} else {
		r = nil
	}
	return
}
func (o *Log) Error() (r *log.Logger) {
	var file *os.File
	if !debugmode.Enabled {
		file = logFile(o.path, o.fileName, o.extension)
	} else {
		r = &log.Logger{}
		file = os.Stderr
	}
	r = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
	r.SetOutput(file)
	return
}
func (o *Log) Warning() (r *log.Logger) {
	var file *os.File
	if !debugmode.Enabled {
		file = logFile(o.path, o.fileName, o.extension)
	} else {
		r = &log.Logger{}
		file = os.Stderr
	}
	r = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Llongfile)
	r.SetOutput(file)
	return
}
func (o *Log) Info() (r *log.Logger) {
	var file *os.File
	if !debugmode.Enabled {
		file = logFile(o.path, o.fileName, o.extension)
	} else {
		r = &log.Logger{}
		file = os.Stdout
	}
	r = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
	r.SetOutput(file)
	return
}
func (o *Log) Fatal() (r *log.Logger) {
	var file *os.File
	if !debugmode.Enabled {
		file = logFile(o.path, o.fileName, o.extension)
	} else {
		r = &log.Logger{}
		file = os.Stderr
	}
	r = log.New(file, "FATAL: ", log.Ldate|log.Ltime|log.Llongfile)
	r.SetOutput(file)
	return
}
