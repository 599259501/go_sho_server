package login

import (
	"errors"
	"fmt"
)

type ILogin interface{
	CheckLogin(loginInfo interface{})(bool,error)
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


type BaseLogin struct{
}
func NewBaseLogin()*BaseLogin{
	return &BaseLogin{}
}
func (login *BaseLogin)CheckLogin(loginInfo interface{})(bool,error){
	return true,nil
}




