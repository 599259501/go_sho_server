package login

import (
	"model/user"
	"fmt"
)
type ICheckLoginHelper interface{
	IsValidUser(userMarket,password string)bool
}

type CheckLoginHelper struct {
}

func NewCheckLoginHelper()*CheckLoginHelper{
	return &CheckLoginHelper{}
}

func (helper *CheckLoginHelper)IsValidUser(userMarket,password string)bool{
	model := user.NewUserModel()

	userInfo, hasInfo, err := model.FindUserInfo(userMarket)
	fmt.Println("userinfo=",userInfo)
	if err!=nil{
		fmt.Println("IsValidUser()has err=",err)
		return false
	}

	if !hasInfo{
		fmt.Println("not found user info")
		return false
	}

	if userInfo.Password != password{
		fmt.Println("IsValidUser(): 密码不正确")
		return false
	}

	return  true
}
