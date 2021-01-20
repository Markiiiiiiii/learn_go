package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type MysqlCofig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	UserName string `ini:"username"`
	Password string `ini:"password"`
}

//打开ini文件
//读取ini文件
//将每条读入结构体
func loadIni(m interface{}, name string, s interface{}) {
	conf := reflect.ValueOf(m)
	sType := reflect.ValueOf(s)
	sKind := sType.Kind()
	if conf.Elem().FieldByName(name).IsValid() { //判断结构体是否存在name这个成员
		switch sKind {
		case reflect.String:
			conf.Elem().FieldByName(name).SetString(sType.String())
		case reflect.Int:
			conf.Elem().FieldByName(name).SetInt(sType.Int())
		}
	}
}

func main() {
	var config MysqlCofig
	k := reflect.TypeOf(config)
	file, err := os.Open("./mysql.ini")
	if err != nil {
		fmt.Printf("文件读取出错，err：%v\n", err)
		return
	}
	defer file.Close()
	configString := bufio.NewReader(file)
	for {
		line, err := configString.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				arr := strings.Split(line, "=")
				if len(arr) > 1 {
					for i := 0; i < k.NumField(); i++ {
						if k.Field(i).Tag.Get("ini") == arr[0] { //根据Tag匹配对应的数据
							stringtmp := strings.Replace(arr[1], "\n", "", -1) //去除字符串中的换行符
							loadIni(&config, k.Field(i).Name, stringtmp)       //password字段直接赋值不做数字转换
						}
					}
				}
			}
			break
		}
		if err != nil {
			fmt.Println("文件逐行读取出错，err：", err)
			return
		}
		strArr := strings.Split(line, "=") //将字符串分割成为数组
		if len(strArr) > 1 {               //判断数组数量是否大于1，忽略掉ini文件的开头声明部分
			for i := 0; i < k.NumField(); i++ {
				if k.Field(i).Tag.Get("ini") == strArr[0] { //根据Tag匹配对应的数据
					stringtmp := strings.Replace(strArr[1], "\n", "", -1)            //去除字符串中的换行符
					var x interface{}                                                //定义一个空接口用来接收不同类型的值
					if num, err := strconv.ParseInt(stringtmp, 10, 64); err == nil { //判断字符串是否可以转化为INT类似的数字
						x = int(num)
					} else {
						x = stringtmp
					}
					loadIni(&config, k.Field(i).Name, x)
				}
			}
		}
	}
	// loadIni(config)
	fmt.Printf("%T,%T,%T,%T\n", config.Address, config.Port, config.UserName, config.Password)
	fmt.Println(config.Address, config.Port, config.UserName, config.Password)
}
