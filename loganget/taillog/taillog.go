package tail

import (
	"fmt"

	"github.com/hpcloud/tail"
)

var (
	tailObj *tail.Tail
)

func Init(filePath string) (err error) {
	config := tail.Config{ //初始化一个配置项
		ReOpen:    true,                                 //是否重新打开
		Follow:    true,                                 //是否跟踪文件
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件哪个位置开始读取
		MustExist: false,                                //文件不存在是否报错
		Poll:      true,                                 //
	}
	tailObj, err = tail.TailFile(filePath, config)
	if err != nil {
		fmt.Println("tail file faild,err:", err)
		return
	}
	return
}
func ReadChan() <-chan *tail.Line {
	return tailObj.Lines
}
