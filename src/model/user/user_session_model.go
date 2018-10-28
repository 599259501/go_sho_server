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
func (model *UserSessionModel)AddUserSession(userId int,sessionType int)(int64,error){
	encryption := model.GetEncryption()
	token,err := model.GetUserToken(strconv.Itoa(userId), encryption)
	if err!=nil{
		return 0,err
	}

	updateTime := time.Now().Format("2006-01-02 15:04:05")

	session := &table_struct.T_USER_SESSION{
		UserId: userId,
		SessionType: sessionType,
		Encryption: encryption,
		Token:token,
		UpdateTime:updateTime,
	}

	affected, err := model.DB.Insert(session)
	return affected,err
}
/**
	获取用户的session数据
 */
func (model *UserSessionModel)GetUserSession(userId int,sessionType int)(table_struct.T_USER_SESSION,bool,error){
	session := table_struct.T_USER_SESSION{}
	has,err := model.DB.Where("user_id = ?", userId).And("session_type = ",sessionType).Get(&session)
	return session,has,err
}
/**
	获取用户token数据
 */
func (model *UserSessionModel)GetUserToken(userId string, encryStr string)(string,error){
	nowTime := time.Now().Unix()

	str := fmt.Sprintf("%s_%d", userId, nowTime)
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

