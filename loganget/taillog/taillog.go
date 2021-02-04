package taillog

import (
	"context"
	"fmt"
	"learn_golang/loganget/kafka"

	"github.com/hpcloud/tail"
)

// var (
// 	tailObj *tail.Tail
// )

//tailtask是一个日志收集的任务
type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
	//为实现退出t.run()
	ctx        context.Context
	cancleFunc context.CancelFunc
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	ctx, cancle := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path:       path,
		topic:      topic,
		ctx:        ctx,
		cancleFunc: cancle,
	}
	tailObj.init() //构建一个初始化方法Init根据路径打开对应的日志
	return
}

func (t TailTask) init() (err error) {
	config := tail.Config{ //初始化一个配置项
		ReOpen:    true,                                 //是否重新打开
		Follow:    true,                                 //是否跟踪文件
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件哪个位置开始读取
		MustExist: false,                                //文件不存在是否报错
		Poll:      true,                                 //
	}
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file faild,err:", err)
		return
	}
	//
	go t.run() //直接去采集日志发送到kafka
	return
}

// func Init(filePath string) (err error) {
// 	config := tail.Config{ //初始化一个配置项
// 		ReOpen:    true,                                 //是否重新打开
// 		Follow:    true,                                 //是否跟踪文件
// 		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件哪个位置开始读取
// 		MustExist: false,                                //文件不存在是否报错
// 		Poll:      true,                                 //
// 	}
// 	tailObj, err = tail.TailFile(filePath, config)
// 	if err != nil {
// 		fmt.Println("tail file faild,err:", err)
// 		return
// 	}
// 	return
// }

// func ReadChan(t *TailTask) <-chan *tail.Line {
// 	return t.instance.Lines
// }

func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Println("tail task is out->", t.path, "->", t.topic)
			return
		case line := <-t.instance.Lines:
			// kafka.SendToKafka(t.topic, line.Text) //函数调用函数
			//先把日志数据发到一个通道里，
			kafka.SendToChan(t.topic, line.Text)
			//在kakfa包中有个单独的goruntine去取日志发到kafka
		}
	}
}
