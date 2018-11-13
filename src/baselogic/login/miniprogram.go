package login

import
(
	"github.com/gin-gonic/gin"
	"baselogic/verify_code"
	"errors"
	"baselogic/wx_helper"
	"model/user"
	MODEL "model"
)

type MiniProgramLoginInfo struct{
	UserName string
	Password string
	VerifyCode string
}

type MiniProgramSessionInfo struct{
	UserId int
	AccessToken string
}

type MiniProgramLogin struct{}
func NewMiniProgramLogin()*MiniProgramLogin{
	return &MiniProgramLogin{}
}

func (loginService *MiniProgramLogin)CheckLogin(loginInfo interface{})(bool,error){
	sessionInfo := loginInfo.(MiniProgramSessionInfo)
	// 这里要做的第一件事情就是校验用户session是否合法，如果不合法就要求用户重新登录
	model := user.NewUserSession()
	// todo 这里要做的事情就是检测用户的
	userSession, hasSession, err := model.GetUserSession(sessionInfo.UserId, MODEL.MINI_PROGRAM_SESSION_TYPE)
	if !hasSession || err!=nil{
		return false,err
	}

	if userSession.Token != sessionInfo.AccessToken {
		return false,errors.New("access_token校验不通过")
	}
	return true,nil
}
func (loginService *MiniProgramLogin) DoLogin(loginInfo interface{})(bool,error){
	loginBody := loginInfo.(MiniProgramLoginInfo)
	// 其次去检测用户对应的密码是否是正确的
	loginHelper := NewCheckLoginHelper()
	if !loginHelper.IsValidUser(loginBody.UserName, loginBody.Password){
		return false,errors.New("用户名账号或者密码错误")
	}
	// 首先先检测用户的验证码是否正确
	verifyTool := verify_code.NewImageVerifyVode()
	if !verifyTool.CheckVerifyCode(loginBody.UserName, loginBody.VerifyCode){
		return false,errors.New("验证码错误")
	}
	return true,nil
}
func (loginService *MiniProgramLogin) AfterLogin(cxt  *gin.Context,params ...string){
	// 这里主要是登陆后刷新用户登陆态，如果是网站登陆接口的 话
}
