package stores

import (
	"model/user"
	"baselogic/wx_helper"
	"model"
	"github.com/sirupsen/logrus"
	"table_struct"
)
type UserStore struct{}

func NewUserStore()*UserStore{
	return &UserStore{}
}

func (store UserStore)RegisterMiniProgramUser(wxSessionInfo wx_helper.WxSessionInfo,isRefreshToken bool)(session *table_struct.TUserSession,err error){
	userModel := user.NewUserModel()

	userId, err := userModel.AddMiniProgramUser(wxSessionInfo)
	if err != nil{
		logrus.Info("AddMiniProgramUser():has err=",err)
		return nil,err
	}

	// 更新session信息
	sessionModel := user.NewUserSession()
	session, err = sessionModel.SavedSession(userId, model.MINI_PROGRAM_SESSION_TYPE)
	return session,err
}