package mylogger

type Filelogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	maxFileSize int64
}

func NewFileLogger(levelstr LogLevel, fp, fn string, maxSize int64) *Filelogger {
	level, err := parseLogLevel(levelstr)
	if err != nil {
		panic(err)
	}
	return &Filelogger{
		Level:       level,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
}
// 判断日志类型，类似一个开关，对于传入的级别以下的日志都打印出来
func (f Filelogger) enable(loglevel LogLevel) bool {
	return f.Level <= loglevel
//将输入的日志字符串格式化
func Logprint(llev LogLevel, format string, a ...interface{}) {
	now := time.Now()
	msg := fmt.Sprintf(format, a...) //将传入的字符串进行格式化输出，输入字符串需带有标识符，用相应的变量去替换
	funcName, fileName, linNum := getInfo(3)
	fmt.Printf("[%s][%s] [%s:%s:%d]%s\n", now.Format("2006-01-02 15:04:05"), printTypestring(llev), fileName, funcName, linNum, msg)
}