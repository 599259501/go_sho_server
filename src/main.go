package main

import (
	"github.com/gin-gonic/gin"
	"router/v1"
	"github.com/joho/godotenv"
	"log"
)

func main(){
	// 初始化.env配置文件
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file,err=",err)
		return
	}
	// 初始化路由信息
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	v1.InitRouter(r)
	// 初始化数据库链接

	// 注册登录handler

	r.Run() // listen and serve on 0.0.0.0:8080
}
