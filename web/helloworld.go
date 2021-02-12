package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//1.创建一个路由
	r := gin.Default()
	//2.绑定路由规则，执行的函数
	//gin.Context,封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	r.Run(":8000")
}
