package baselogic

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func JResponse(cxt *gin.Context,code int,data interface{},msg string)*gin.Context{
	cxt.JSON(http.StatusOK,gin.H{
		"msg": msg,
		"code": code,
		"data": data,
	})
	cxt.Abort()
	return cxt
}
