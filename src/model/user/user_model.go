package user

import
(
	"github.com/go-xorm/xorm"
	"table_struct"
	"fmt"
	"model"
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

func (model *UserModel)FindUserInfo(userKey string)(*table_struct.TUser,bool,error){
	userInfo := &table_struct.TUser{}
	has,err :=model.DB.Where("name=?",userKey).Get(userInfo)
	if err!=nil{
		return nil,has,err
	}

	return userInfo,has,nil
}
