package user

import
(
	"github.com/go-xorm/xorm"
	"common/function"
	"time"
	"fmt"
	"table_struct"
	"model"
)

type UserSessionModel struct{
	DB *xorm.Engine
	DataBaseName string
}

func NewUserSession()*UserSessionModel{
	db,err := model.GDBManager.GetDBEngin(model.USER_DATASOURCE)
	if err!=nil{
		fmt.Println("GetDBEngin() has err=",err)
	}

	return &UserSessionModel{
		DataBaseName: "",
		DB:db,
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
		UserId: userId,
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
	session,has,err := model.GetUserSession(userId, sessionType)

	if err!=nil{
		return session,err
	}
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
	return model.EncryStr(str, encryStr)
}
/**
	加密用户token数据
 */
func (model *UserSessionModel)EncryStr(str string,encryption string) (string,error){
	return function.AesEncryStr(str,encryption)
}
/**
	获取用户token加密串
 */
func (model *UserSessionModel)GetEncryption()string{
	str, _ := function.GetRandStr(32)
	return str
}

