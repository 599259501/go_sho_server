package main

import (
	"github.com/gin-gonic/gin"
	"router/v1"
	"github.com/joho/godotenv"
	_ "github.com/sirupsen/logrus"
	"logging"
	"utils"
	"fmt"
)

func main(){
	// 初始化.env配置文件
	if err := InitEnvFile();err!=nil {
		fmt.Println("Error loading .env file,err=",err)
		return
	}
	// 初始化日志信息
	logging.InitLogger()
	// 设置mode模式
	ginMode := utils.GetEnv("GIN_MODE", gin.ReleaseMode)
	gin.SetMode(ginMode)
	r := gin.Default()
	// 初始化路由信息
	listenAddr := utils.GetEnv("LISTEN_ADDR", ":8080")
	v1.InitRouter(r)
	r.Run(listenAddr) // listen and serve on 0.0.0.0:8080
}

func InitEnvFile()error{
	err := godotenv.Load()
	return err
}
