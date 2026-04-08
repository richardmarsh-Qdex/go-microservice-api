package logger

import (
	"log"
	"os"
)

var L = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

func Infof(format string, v ...interface{}) {
	L.Printf("[INFO] "+format, v...)
}

func Errorf(format string, v ...interface{}) {
	L.Printf("[ERROR] "+format, v...)
}
