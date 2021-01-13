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
func (l Logger) enable(loglevel LogLevel) bool {
	return l.Level <= loglevel
}
func Logprint(llev LogLevel, msg string) {
	now := time.Now()
	funcName, fileName, linNum := getInfo(3)
	fmt.Printf("[%s][%s] [%s-->%s-->%d]%s\n", now.Format("2006-01-02 15:04:05"), printTypestring(llev), fileName, funcName, linNum, msg)
}

//Debug 错误
func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		Logprint(DEBUG, msg)
	}
}

// Info 详细信息
func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		Logprint(INFO, msg)
	}
}

//Waring 警告信息
func (l Logger) Waring(msg string) {
	if l.enable(WARING) {
		Logprint(WARING, msg)
	}
}

//Fatal 致命错误
func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		Logprint(FATAL, msg)
	}
}

//Error 错误信息
func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		Logprint(ERROR, msg)
	}
}
