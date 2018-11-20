package login

import
(
	"github.com/gin-gonic/gin"
	"baselogic/wx_helper"
	"errors"
	"model/user"
	MODEL "model"
	"github.com/sirupsen/logrus"
	"stores"
	"table_struct"
	"fmt"
	"time"
	"utils"
	"strconv"
)

type MiniProgramLoginInfo struct{
	Code string
}

type MiniProgramSessionInfo struct{
	UserId int `json:"user_id"`
	AccessToken string `json:"access_token"`
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
	fmt.Println("sessionInfo=", sessionInfo)
	userSession, hasSession, err := model.GetUserSession(sessionInfo.UserId, MODEL.MINI_PROGRAM_SESSION_TYPE)
	if !hasSession || err!=nil{
		return false,err
	}
	sessionExpires,_ := strconv.ParseInt(utils.GetEnv("SESSION_EXPIRES", "0"),10,64)
	if userSession.Token != sessionInfo.AccessToken || (time.Now().Unix()-userSession.UpdateTime.Unix()) > sessionExpires{
		return false,errors.New("access_token校验不通过")
	}
	return true,nil
}
func (loginService *MiniProgramLogin) DoLogin(loginInfo interface{})(bool,error,interface{}){
	/*loginBody := loginInfo.(MiniProgramLoginInfo)
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
	return true,nil*/

	// 小程序登陆是根据code验证它的可靠性
	loginBody := loginInfo.(MiniProgramLoginInfo)

	helper := wx_helper.NewWxHelper()
	sessionInfo, err :=helper.GetWxMiniSession(loginBody.Code)
	if err!=nil{
		logrus.Info("GetWxMiniSession() has err,err=",err)
		return false, err,nil
	}

	// code 验证成功之后就可以插入session表了
	// 根据openId 找到user_id
	userModel := user.NewUserModel()
	userInfo, isExists, err := userModel.FindUserByOpenId(sessionInfo.OpenId)
	if err!=nil{
		logrus.Info("FindUserByOpenId():has err=",err, ",openId=", sessionInfo.OpenId)
		return false,err,nil
	}

	tSession := &table_struct.TUserSession{}
	if !isExists{
		// 如果用户信息不存在就注册&加上session信息
		userStore := stores.NewUserStore()
		tSession,err = userStore.RegisterMiniProgramUser(sessionInfo, true)
		if err!=nil{
			logrus.Info("RegisterMiniProgramUser() has err=",err)
			return false,err,nil
		}
	}

	model := user.NewUserSession()
	tSession, err = model.SavedSession(userInfo.Id, MODEL.MINI_PROGRAM_SESSION_TYPE)
	fmt.Println("tSession=", tSession)
	if err!=nil{
		logrus.Info("DoLogin(): addUserSession has err=", err)
		return false,err,nil
	}

	return true,nil,tSession
}
func (loginService *MiniProgramLogin) AfterLogin(cxt  *gin.Context,params map[string]interface{}){
}
