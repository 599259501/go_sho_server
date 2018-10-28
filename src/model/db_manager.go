package model

import (
	"github.com/go-xorm/xorm"
)
type DBManager struct{
	DBMap map[string]*xorm.Engine
}
func NewDBManager()*DBManager{
	return &DBManager{}
}
func (manager *DBManager)GetDBEngin(dbName string)(*xorm.Engine,error){
	if _,ok:=manager.DBMap[dbName];!ok{
		return nil,nil
	}
	return manager.DBMap[dbName],nil
}
