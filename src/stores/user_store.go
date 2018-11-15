package stores

import (
	"model/user"
)
type UserStore struct{}

func NewUserStore()*UserStore{
	return &UserStore{}
}

func (store UserStore)RegisterMiniProgramUser(openId string,isRefreshToken bool){
	userModel := user.NewUserModel()

}