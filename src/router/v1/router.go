package v1

import
(
	"github.com/gin-gonic/gin"
	"router/v1/controllers"
)

func InitRouter(router *gin.Engine){
	v1 := router.Group("/v1")
	{
		v1.POST("/login", controllers.DoLogin)
	}
}
