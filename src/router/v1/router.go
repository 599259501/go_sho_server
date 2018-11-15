package v1

import
(
	"github.com/gin-gonic/gin"
	"router/v1/controllers"
	"middares"
)

func InitRouter(router *gin.Engine){
	v1 := router.Group("/v1")
	{
		v1.POST("/mini_login", controllers.DoLogin)

		v1.Use(middares.CheckMiniProgramLoginInfo()).POST("/get_home_data", controllers.GetHomePageData)
	}
}
