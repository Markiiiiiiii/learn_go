package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

//自定义日志库
//LogLevel 类型
type LogLevel uint16

//定义一个logger接口
type LogInterface interface {
	Debug(format string, a ...interface{})
	Info(format string, a ...interface{})
	Waring(format string, a ...interface{})
	Fatal(format string, a ...interface{})
	Error(format string, a ...interface{})
}

// 定义日志级别
const (
	UNKNOW LogLevel = iota
	DEBUG
	INFO
	WARING
	FATAL
	ERROR
)

//Logger 日志结构体
type Logger struct {
	Level LogLevel //结构体中定义一个LogLevel类型的
}

//将构造函数接收到的字符串解析成自定义的LogLevel类型
func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "error":
		return ERROR, nil
	case "waring":
		return WARING, nil
	case "fatal":
		return FATAL, nil
	case "info":
		return INFO, nil
	default:
		err := errors.New("无效日志级别")
		return UNKNOW, err
	}
}
func printTypestring(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case ERROR:
		return "ERROR"
	case WARING:
		return "WARING"
	case FATAL:
		return "FATAL"
	case INFO:
		return "INFO"

	}
	return "DEBUG"
}

func getInfo(n int) (funcName, fileName string, lineNum int) {
	pc, file, lineNum, ok := runtime.Caller(n)
	if !ok {
		fmt.Println("runtime.caller() faild ")
		return
	}
	funcName = strings.Split(runtime.FuncForPC(pc).Name(), ".")[1]
	fileName = path.Base(file)
	return
}
