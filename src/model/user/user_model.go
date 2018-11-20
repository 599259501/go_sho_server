package user

import
(
	"github.com/go-xorm/xorm"
	"table_struct"
	"fmt"
	"model"
	"time"
	"baselogic/wx_helper"
	"github.com/sirupsen/logrus"
	"errors"
)
type UserModel struct{
	DB *xorm.Engine
	DataBaseName string
}

func NewUserModel()*UserModel{
	db,err := model.GDBManager.GetDBEngin(model.USER_DATASOURCE)
	if err!=nil{
		fmt.Println("GetDBEngin() has err=",err)
	}

	return &UserModel{
		DB:db,
	}
}
// 根据用户唯一标识查找用户信息
func (model *UserModel)FindUserInfo(userKey string)(*table_struct.TUser,bool,error){
	userInfo := &table_struct.TUser{}
	has,err :=model.DB.Where("name=?",userKey).Get(userInfo)
	if err!=nil{
		return nil,has,err
	}

	return userInfo,has,nil
}
// 根据用户openId获取用户信息
func (model *UserModel)FindUserByOpenId(openId string)(*table_struct.TUser,bool,error){
	userInfo := &table_struct.TUser{}
	has,err :=model.DB.Where("open_id=?",openId).Get(userInfo)
	if err!=nil{
		return nil,has,err
	}

	return userInfo,has,nil
}
// 插入用户数据
func (model *UserModel)AddMiniProgramUser(wxSessionInfo wx_helper.WxSessionInfo)(userId int,err error){
	userInfo := table_struct.TUser{
		OpenId: wxSessionInfo.OpenId,
		CreateTime: time.Now(),
	}

	if model.DB == nil{
		logrus.Info("AddMiniProgramUser(): model db is nil")
		return 0,errors.New("model db is nil")
	}
	_,err = model.DB.Insert(&userInfo)

	return userInfo.Id,err
}
