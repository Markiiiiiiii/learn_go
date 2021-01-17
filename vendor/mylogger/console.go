package mylogger

//**
//向终端打印log练习
//**
import (
	"fmt"
	"time"
)

//NewConsoleLog 构造函数
func NewConsoleLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		fmt.Printf("NewConsoleLog函数构造失败，err:%v", err)
	}
	return Logger{Level: level} //返回一个结构体
}

//判断日志类型，类似一个开关，对于传入的级别以下的日志都打印出来
func (l Logger) enable(loglevel LogLevel) bool {
	return l.Level <= loglevel
}

//将输入的日志字符串格式化打印
func (l Logger) logprint(levStr LogLevel, format string, a ...interface{}) {
	now := time.Now()
	format = fmt.Sprintf(format, a...) //将传入的字符串进行格式化输出，输入字符串需带有标识符，用相应的变量去替换
	funcName, fileName, linNum := getInfo(3)
	fmt.Printf("[%s][%s] [%s:%s:%d]%s\n", now.Format("2006-01-02 15:04:05"), printTypestring(levStr), fileName, funcName, linNum, format)
}

//Debug 错误
func (l Logger) Debug(format string, a ...interface{}) {
	if l.enable(DEBUG) {
		l.logprint(DEBUG, format, a...)
	}
}

// Info 详细信息
func (l Logger) Info(format string, a ...interface{}) {
	if l.enable(INFO) {
		l.logprint(INFO, format, a...)
	}
}

//Waring 警告信息
func (l Logger) Waring(format string, a ...interface{}) {
	if l.enable(WARING) {
		l.logprint(WARING, format, a...)
	}
}

//Fatal 致命错误
func (l Logger) Fatal(format string, a ...interface{}) {
	if l.enable(FATAL) {
		l.logprint(FATAL, format, a...)
	}
}

//Error 错误信息
func (l Logger) Error(format string, a ...interface{}) {
	if l.enable(ERROR) {
		l.logprint(ERROR, format, a...)
	}
}
