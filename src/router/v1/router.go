package v1

import "github.com/gin-gonic/gin"

func InitRouter(router *gin.Engine){
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
	}
}
