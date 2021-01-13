package mylogger

import (
	"fmt"
	"time"
)

//Newlog 构造函数
func Newlog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{Level: level} //返回一个结构体
}

//判断日志类型，类似一个开关，对于传入的级别以下的日志都打印出来
func (l Logger) enable(loglevel LogLevel) bool {
	return l.Level <= loglevel
}

//将输入的日志字符串格式化
func Logprint(llev LogLevel, format string, a ...interface{}) {
	now := time.Now()
	msg := fmt.Sprintf(format, a...) //将传入的字符串进行格式化输出，输入字符串需带有标识符，用相应的变量去替换
	funcName, fileName, linNum := getInfo(3)
	fmt.Printf("[%s][%s] [%s:%s:%d]%s\n", now.Format("2006-01-02 15:04:05"), printTypestring(llev), fileName, funcName, linNum, msg)
}

//Debug 错误
func (l Logger) Debug(format string, a ...interface{}) {
	if l.enable(DEBUG) {
		Logprint(DEBUG, format, a...)
	}
}

// Info 详细信息
func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		Logprint(INFO, format, a...)
	}
}

//Waring 警告信息
func (l Logger) Waring(msg string) {
	if l.enable(WARING) {
		Logprint(WARING, format, a...)
	}
}

//Fatal 致命错误
func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		Logprint(FATAL, format, a...)
	}
}

//Error 错误信息
func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		Logprint(ERROR, format, a...)
	}
}
