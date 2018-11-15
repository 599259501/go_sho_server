package controllers

import
(
	"github.com/gin-gonic/gin"
	"baselogic"
)

func GetHomePageData(cxt * gin.Context){
	page := cxt.GetInt("page")
	pageSize := cxt.GetInt("page_size")

	baselogic.JResponse(cxt, 0, map[string]interface{}{
		"page": page,
		"page_size": pageSize,
	}, "ok")
}
