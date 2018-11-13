package controllers

import(
	"github.com/gin-gonic/gin"
	"baselogic/login"
	"baselogic"
)
func DoLogin(cxt * gin.Context){
	userName := cxt.DefaultPostForm("user_name", "")
	password := cxt.DefaultPostForm("password", "") // 这里前端必须用md5加密过密码
	loginType := cxt.DefaultPostForm("login_type", login.MINI_PROGRAM)
	verifyCode := cxt.DefaultPostForm("verify_code", "")

	if userName == "" || password == "" || verifyCode == ""{
		baselogic.JResponse(cxt, login.PARAM_ERROR_CODE,nil, "用户名/密码/验证码不能为空")
		return
	}
	isLogin,err := login.GLoginManager.ProcessLoginByType(userName,password,verifyCode,loginType)
	if !isLogin{
		baselogic.JResponse(cxt, -100,nil, err.Error())
		return
	}

	// 登陆成功就刷新状态
	loginHandler := login.GLoginManager.GetLoginMethod(loginType)
	loginHandler.AfterLogin(cxt, userName, password)

	baselogic.JResponse(cxt, login.SUCCESS_CODE,nil, "ok")
}