package user

import
(
	"github.com/go-xorm/xorm"
	"common/function"
	"time"
	"fmt"
	"table_struct"
	"strconv"
)

const (
	MINI_PROGRAM_SESSION_TYPE = 0
	H5_SESSION_TYPE = 1
	ANDROID_SESSION_TYPE = 2
)

type UserSessionModel struct{
	DB *xorm.Engine
	DataSource string
}

func NewUserSession()*UserSessionModel{
	return &UserSessionModel{
		DataSource: "",
		DB:nil,
	}
}
/**
	插入session数据
 */
func (model *UserSessionModel)AddUserSession(userId int,sessionType int)(*table_struct.TUserSession,error){
	encryption := model.GetEncryption()
	token,err := model.GetUserToken(userId, encryption)
	if err!=nil{
		return nil,err
	}

	updateTime := time.Now()

	session := &table_struct.TUserSession{
		UserId: userId,
		SessionType: sessionType,
		Encryption: encryption,
		Token:token,
		UpdateTime:updateTime,
	}

	_, err = model.DB.Insert(session)
	return session,err
}
/**
	获取用户的session数据
 */
func (model *UserSessionModel)GetUserSession(userId int,sessionType int)(*table_struct.TUserSession,bool,error){
	session := &table_struct.TUserSession{}
	has,err := model.DB.Where("user_id = ?", userId).And("session_type = ?",sessionType).Get(session)
	return session,has,err
}
/**
	更新用户session数据
 */
func (model *UserSessionModel)UpdateUserSession(userId int,sessionType int)(session *table_struct.TUserSession,err error){
	token,err := model.GetUserToken(userId, model.GetEncryption())
	if err!=nil{
		return nil,err
	}

	session = &table_struct.TUserSession{
		Token:token,
		UpdateTime:time.Now(),
	}

	affected,err := model.DB.Where("user_id = ?",userId).And("session_type = ?",sessionType).Update(session)
	if affected <= 0 || err!=nil{
		return nil,err
	}

	return session,nil
}
/**
	没有就新增数据，有就更新数据
 */
func (model *UserSessionModel)SavedSession(userId int,sessionType int)(session *table_struct.TUserSession,err error){
	session,has,_ := model.GetUserSession(userId, sessionType)
	if !has{
		session,err = model.AddUserSession(userId,sessionType)
	} else {
		session,err = model.UpdateUserSession(userId, sessionType)
	}
	return session,err
}
/**
	获取用户token数据
 */
func (model *UserSessionModel)GetUserToken(userId int, encryStr string)(string,error){
	nowTime := time.Now().Unix()

	str := fmt.Sprintf("%d_%d", userId, nowTime)
	return model.DscStr(str, encryStr)
}
/**
	加密用户token数据
 */
func (model *UserSessionModel)DscStr(str string,encryption string) (string,error){
	return function.AesDecryStr(str,encryption)
}
/**
	获取用户token加密串
 */
func (model *UserSessionModel)GetEncryption()string{
	str, _ := function.GetRandStr(64)
	return str
}

