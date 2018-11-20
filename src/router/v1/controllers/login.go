package controllers

import(
	"github.com/gin-gonic/gin"
	"baselogic/login"
	"baselogic"
	"table_struct"
)
func DoLogin(cxt * gin.Context){
	code := cxt.DefaultPostForm("code", "")
	if code == ""{
		baselogic.JResponse(cxt, login.PARAM_ERROR_CODE,nil, "code参数不能为空")
		return
	}
	isLogin,err,extraInfo := login.GLoginManager.ProcessLoginByType(map[string]interface{}{
		"code":code,
	})
	if !isLogin && err !=nil{
		baselogic.JResponse(cxt, -100,nil, err.Error())
		return
	}
	userSession := &table_struct.TUserSession{}
	// 如果是第一次登陆的用户，就注册用户信息
	userSession = extraInfo.(*table_struct.TUserSession)
	// 登陆成功就刷新状态
	loginHandler := login.GLoginManager.GetLoginMethod(login.MINI_PROGRAM)
	loginHandler.AfterLogin(cxt, map[string]interface{}{
		"session": extraInfo,
	})

	baselogic.JResponse(cxt, login.SUCCESS_CODE,login.MiniProgramSessionInfo{
		UserId: userSession.UserId,
		AccessToken: userSession.Token,
	}, "ok")
}