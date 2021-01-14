package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//**
//向文件输出log练习
//**
type Filelogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

//构造函数
func NewFilelogger(levelStr, fPath, fName string, maxSize int64) *Filelogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &Filelogger{
		Level:       level,
		filePath:    fPath,
		fileName:    fName,
		maxFileSize: maxSize,
	}
	err = fl.initFile()
	if err != nil {
		panic(err)
	}
	return fl
}
func (f *Filelogger) initFile() error {
	fullFilePath := path.Join(f.filePath, f.fileName)
	//打开记录正常日志的日志文件
	fileObj, err := os.OpenFile(fullFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Open LOG file faild,err:%v\n", err)
		return err
	}
	//打开专用于记录error类日志的日志文件
	errFileObj, err := os.OpenFile(fullFilePath+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Open ERRORLOG file faild,err:%v\n", err)
		return err
	}
	//日志文件都已打开
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

// 判断日志类型，类似一个开关，对于传入的级别以下的日志都打印出来
func (f *Filelogger) enable(loglevel LogLevel) bool {
	return f.Level <= loglevel
}

//将输入的日志字符串格式化打印
func (f *Filelogger) logprint(levStr LogLevel, format string, a ...interface{}) {
	now := time.Now()
	format = fmt.Sprintf(format, a...) //将传入的字符串进行格式化输出，输入字符串需带有标识符，用相应的变量去替换
	funcName, fileName, linNum := getInfo(3)
	fmt.Fprintf(f.fileObj, "[%s][%s] [%s:%s:%d]%s\n", now.Format("2006-01-02 15:04:05"), printTypestring(levStr), fileName, funcName, linNum, format)
	// 如果级别大于等于ERROR级别则写入error类日志文件中
	if levStr >= ERROR {
		fmt.Fprintf(f.errFileObj, "[%s][%s] [%s:%s:%d]%s\n", now.Format("2006-01-02 15:04:05"), printTypestring(levStr), fileName, funcName, linNum, format)
	}
}

//关闭log文件
func (f *Filelogger) closeFile() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

//Debug 错误
func (f *Filelogger) Debug(format string, a ...interface{}) {
	f.logprint(DEBUG, format, a...)
}

// Info 详细信息
func (f *Filelogger) Info(format string, a ...interface{}) {
	f.logprint(INFO, format, a...)
}

//Waring 警告信息
func (f *Filelogger) Waring(format string, a ...interface{}) {
	f.logprint(WARING, format, a...)
}

//Fatal 致命错误
func (f *Filelogger) Fatal(format string, a ...interface{}) {
	f.logprint(FATAL, format, a...)
}

//Error 错误信息
func (f *Filelogger) Error(format string, a ...interface{}) {
	f.logprint(ERROR, format, a...)
}
