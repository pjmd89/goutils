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
	err = os.Mkdir(path, 0666)
	if err != nil {
		log.Fatal("log not loaded: " + err.Error())
		return nil
	}
	r.path = path
	r.fileName = fileName
	r.extension = extension
	return
}
func logFile(path string, fileName string, extension string) (r *os.File) {
	if !debugmode.Enabled {
		year, month, day := time.Now().Date()
		filePath := fmt.Sprintf(path+"/"+fileName+"-%v-%v-%v."+extension, day, month.String(), year)
		file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal("log not loaded: " + err.Error())
		}
		defer file.Close()
		//r.logger = log.New(file, prefix, log.Ldate|log.Ltime|log.Llongfile)
	} else {
		r = nil
	}
	return
}
func (o *Log) Error(s string, vars ...any) (r *log.Logger) {
	if !debugmode.Enabled {
		file := logFile(o.path, o.fileName, o.extension)
		r = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
	} else {
		r = &log.Logger{}
		r.SetPrefix("ERROR: ")
		r.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	}
	return
}
func (o *Log) Warning(s string, vars ...any) (r *log.Logger) {
	if !debugmode.Enabled {
		file := logFile(o.path, o.fileName, o.extension)
		r = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Llongfile)
	} else {
		r = &log.Logger{}
		r.SetPrefix("WARNING: ")
		r.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	}
	return
}
func (o *Log) Info(s string, vars ...any) (r *log.Logger) {
	if !debugmode.Enabled {
		file := logFile(o.path, o.fileName, o.extension)
		r = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
	} else {
		r = &log.Logger{}
		r.SetPrefix("INFO: ")
		r.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	}
	return
}
func (o *Log) Fatal(s string, vars ...any) (r *log.Logger) {
	if !debugmode.Enabled {
		file := logFile(o.path, o.fileName, o.extension)
		r = log.New(file, "FATAL: ", log.Ldate|log.Ltime|log.Llongfile)
	} else {
		r = &log.Logger{}
		r.SetPrefix("FATAL: ")
		r.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	}
	return
}
