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
		fmt.Printf("打开Log文件出错,err:%v\n", err)
		return err
	}
	//打开专用于记录error类日志的日志文件
	errFileObj, err := os.OpenFile(fullFilePath+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开ErrorLog文件出错,err:%v\n", err)
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

//切割日志文件
func (f *Filelogger) splitFile(file *os.File) (*os.File, error) { //传入文件名参数判断是否是正常的log文件还是error log 文件
	//2. rename 备份一下当前文件 xx.log --> xx.log.bak2021011214
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("splitFile时获取文件状态出错,err:%v\n", err)
		return nil, err
	}
	oldLogFile := path.Join(f.filePath, fileInfo.Name())      //当前文件路径
	newLogFile := fmt.Sprintf("%s.bak%s", oldLogFile, nowStr) //新建文件路径
	// 1.关闭当前文件
	file.Close()
	os.Rename(oldLogFile, newLogFile)
	//3.打开一个新的源log文件
	newFileObj, err := os.OpenFile(oldLogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("不能创建一个新的日志文件,err:%v\n", err)
		return nil, err
	}
	//4.将打开的新的日志文件赋值给f.fileObj
	return newFileObj, nil
}

//检查文件大小是否需要切割
func (f *Filelogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("checkSize检查文件获取文件状态出错,err:%v\n", err)
		return false
	} //判断log文件的大小是否大于等于要求切割的文件大小值
	return fileInfo.Size() >= f.maxFileSize
}

//生成日志文件
func (f *Filelogger) logWrite(levStr LogLevel, format string, a ...interface{}) {
	now := time.Now()
	format = fmt.Sprintf(format, a...) //将传入的字符串进行格式化输出，输入字符串需带有标识符，用相应的变量去替换
	funcName, fileName, linNum := getInfo(3)
	if f.checkSize(f.fileObj) {
		//判断是否需要切割文件
		newFile, err := f.splitFile(f.fileObj)
		if err != nil {
			fmt.Printf("无法切割Log文件，err:%v\n", err)
			return
		}
		f.fileObj = newFile
	}
	fmt.Fprintf(f.fileObj, "[%s][%s] [%s:%s:%d]%s\n", now.Format("2006-01-02 15:04:05"), printTypestring(levStr), fileName, funcName, linNum, format)
	// 如果级别大于等于ERROR级别则写入error类日志文件中
	if levStr >= ERROR {
		if f.checkSize(f.errFileObj) {
			newFile, err := f.splitFile(f.errFileObj)
			if err != nil {
				fmt.Printf("无法切割ErrorLog文件，err:%v\n", err)
				return
			}
			f.errFileObj = newFile
		}
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
	if f.enable(DEBUG) {
		f.logWrite(DEBUG, format, a...)
	}
}

// Info 详细信息
func (f *Filelogger) Info(format string, a ...interface{}) {
	if f.enable(INFO) {
		f.logWrite(INFO, format, a...)
	}
}

//Waring 警告信息
func (f *Filelogger) Waring(format string, a ...interface{}) {
	if f.enable(WARING) {
		f.logWrite(WARING, format, a...)
	}
}

//Fatal 致命错误
func (f *Filelogger) Fatal(format string, a ...interface{}) {
	if f.enable(FATAL) {
		f.logWrite(FATAL, format, a...)
	}
}

//Error 错误信息
func (f *Filelogger) Error(format string, a ...interface{}) {
	if f.enable(ERROR) {
		f.logWrite(ERROR, format, a...)
	}
}
