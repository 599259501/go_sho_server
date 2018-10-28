package table_struct

type T_USER_SESSION struct{
	Id int `xorm:"int pk autoincr 'id'"`
	UserId int `xorm:"int notnull 'user_id'"`
	SessionType int `xorm:"TINYINT notnull 'session_type'"`
	Encryption string `xorm:"varchar(255) notnull 'encryption'"`
	Token string `xorm:"varchar(255) notnull 'token'"`
	UpdateTime string `xorm:"TIMESTAMP notnull 'update_time'"`
}
