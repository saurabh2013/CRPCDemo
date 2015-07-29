package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

//Log Info
func Info(msg ...interface{}) {
	printLog(INFO, msg...)
}

func Infof(format string, err ...interface{}) {
	printLog(INFO, fmt.Sprintf(format, err...))
}

//Log Errors
func Error(err ...interface{}) {
	printLog(ERROR, err...)
}

func Errorf(format string, err ...interface{}) {
	printLog(ERROR, fmt.Sprintf(format, err...))
}

const (
	ERROR = "Error"
	INFO  = "Info"
)

func printLog(typ string, obj ...interface{}) (b bool) {
	defer func() {
		if obj != nil {
			_, fileName, line, _ := runtime.Caller(3)
			log.Printf("[%s] %v [%s:%d]", typ, fmt.Sprint(obj...), fileName, line)

			b = true
		}
	}()
	return
}

func LogTofilef(format string, obj ...interface{}) {
	f, er := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if er != nil {
		log.Fatalf("error opening file: %v", er)
	}
	defer f.Close()
	log.SetOutput(f)
	 printLog(INFO, fmt.Sprintf(format, obj...))
}
 