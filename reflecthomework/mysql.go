package main

import "fmt"

type MysqlCofig struct {
	Address  string `ini:"address"`
	Port     int    `ini"port"`
	UserName string `ini:"username"`
	Password string `ini:"password"`
}

//打开ini文件
//将每条读入结构体
func loadIni(file interface{}) {

}
func main() {
	var config MysqlCofig
	loadIni(config)
	fmt.Println(config.Address, config.Port, config.UserName, config.Password)
}
