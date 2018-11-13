package model

import (
	"github.com/go-xorm/xorm"
	"errors"
	"fmt"
 _ "github.com/go-sql-driver/mysql"
)

func init(){
	fmt.Println("begin init dbmanager")
	GDBManager = NewDBManager()

	engin,err := xorm.NewEngine("mysql", "root:599259501c++@localhost:3306/db_business?charset=utf8")
	if err!=nil{
		fmt.Println("newEngin has err=",err)
	} else {
		GDBManager.RegisterDBEngin(BUSINESS_DATASOURCE, engin)
	}

	ginEngin,err := xorm.NewEngine("mysql", "root:599259501c++@/db_gin?charset=utf8")
	if err!=nil{
		fmt.Println("newEngin has err=",err)
	} else {
		GDBManager.RegisterDBEngin(USER_DATASOURCE, ginEngin)
	}
}

var GDBManager *DBManager

type DBManager struct{
	DBMap map[string]*xorm.Engine
}
func NewDBManager()*DBManager{
	return &DBManager{
		DBMap: make(map[string]*xorm.Engine),
	}
}
func (manager *DBManager)GetDBEngin(dbName string)(*xorm.Engine,error){
	if _,ok:=manager.DBMap[dbName];!ok{
		return nil,errors.New(fmt.Sprintf("not found name=%s engin", dbName))
	}
	return manager.DBMap[dbName],nil
}

func (manager *DBManager)RegisterDBEngin(dbName string, engin *xorm.Engine)error{
	if engin == nil{
		return errors.New("engin is nil")
	}

	if _,ok := manager.DBMap[dbName];ok{
		return errors.New(fmt.Sprintf("%s is exists,not register", dbName))
	}

	manager.DBMap[dbName] = engin
	return nil
}
