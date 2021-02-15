package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	//1.创建一个路由
	r := gin.Default()
	//2.绑定路由规则，执行的函数
	//gin.Context,封装了request和response
	// r.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "hello world")
	// })

	// //1. api 传参数
	// r.GET("/user/:name/*action", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	action := c.Param("action")
	// 	c.String(http.StatusOK, name+" is "+action)
	// })

	// //2. URL传参数
	// r.GET("/test", func(c *gin.Context) {
	// 	name := c.DefaultQuery("name", "zhangsan")
	// 	c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	// })

	// //3.1表单传参
	// r.POST("/form", func(c *gin.Context) {
	// 	//表单默认值
	// 	type1 := c.DefaultPostForm("type", "alert")
	// 	//接受其他的参数
	// 	username := c.PostForm("username")
	// 	password := c.PostForm("password")
	// 	hobbys := c.PostFormArray("hobby")
	// 	c.String(http.StatusOK, fmt.Sprintf("type is %s  username %s password %s hobby %v", type1, username, password, hobbys))
	// })

	// //3.2表单传单个文件
	// r.POST("/upload", func(c *gin.Context) {
	// 	file, _ := c.FormFile("file")
	// 	log.Println(file.Filename)
	// 	//传到根目录，名字就用本身
	// 	c.SaveUploadedFile(file, file.Filename)
	// 	c.String(200, fmt.Sprintf("'%s' upload", file.Filename))
	// })

	// //3.3表单传多个文件
	// //限制表单上传文件大小 8M，不设置则默认为32m
	// r.MaxMultipartMemory = 8 << 20
	// r.POST("/upload", func(c *gin.Context) {
	// 	form, err := c.MultipartForm()
	// 	if err != nil {
	// 		c.String(http.StatusBadRequest, fmt.Sprintf("get err %s ", err.Error()))
	// 		return
	// 	}
	// 	//获取所有的文件
	// 	fiels := form.File["files"]
	// 	// 遍历所有图片
	// 	for _, file := range fiels {
	// 		//逐个存储
	// 		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
	// 			c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s ", err.Error()))
	// 			return
	// 		}
	// 	}
	// 	c.String(200, fmt.Sprintf(" '%d' upload", len(fiels)))
	// })

	// 4.路由组
	//4.1路由组1.处理GET请求
	v1 := r.Group("/v1")
	{
		// 花括号{}是书写规范
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}
	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}
	r.Run(":8000")
}
func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}
