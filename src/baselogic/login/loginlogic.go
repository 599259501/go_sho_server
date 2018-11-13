package login

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

var GLoginManager *LoginManager

func init(){
	GLoginManager = NewLoginManager()
	// 注册小程序登陆处理handler
	GLoginManager.RegisterLoginMethod(MINI_PROGRAM, NewMiniProgramLogin())
}

type ILogin interface{
	// 检测登陆态
	CheckLogin(loginInfo interface{})(bool,error)
	DoLogin(loginInfo interface{})(bool,error)
	AfterLogin(cxt  *gin.Context,params ...string)
}


type LoginManager struct{
	LoginMethods map[string]ILogin
}
func NewLoginManager()*LoginManager {
	return &LoginManager{
		LoginMethods: map[string]ILogin{},
	}
}
func (manager *LoginManager)RegisterLoginMethod(method string,login ILogin)error{
	if _,ok:=manager.LoginMethods[method];ok{
		return errors.New(fmt.Sprintf("%s has exists", method))
	}

	manager.LoginMethods[method] = login
	return nil
}
func (manager *LoginManager)GetLoginMethod(methodName string)ILogin{
	var method ILogin
	var ok bool
	if method,ok = manager.LoginMethods[methodName];!ok{
		// todo add logging
		return nil
	}
	return method
}
func (manager *LoginManager)ProcessLoginByType(userName,password,verifyCode,loginType string)(bool,error){
	loginHandler := manager.GetLoginMethod(loginType)
	if loginHandler == nil{
		return false,errors.New("not support login")
	}

	isLogin := false
	var err error
	switch loginType {
	case MINI_PROGRAM:
			isLogin, err = loginHandler.DoLogin(MiniProgramLoginInfo{
				UserName:userName,
				Password:password,
				VerifyCode:verifyCode,
			})
	default:
		err = errors.New("not support logintype")
	}

	return isLogin,err
}


type BaseLogin struct{
}
func NewBaseLogin()*BaseLogin{
	return &BaseLogin{}
}
func (login *BaseLogin)CheckLogin(loginInfo interface{})(bool,error){
	return true,nil
}




