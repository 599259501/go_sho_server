package table_struct

import (
	"time"
)

type TUser struct {
	Id         int       `xorm:"not null pk autoincr comment('用户唯一id') INT(11)"`
	Name       string    `xorm:"not null default '' comment('用户昵称') VARCHAR(20)"`
	Password   string 	 `xorm:"not null default '' comment('用户加密密码') VARCHAR(255)"`
	Phone      string    `xorm:"not null VARCHAR(30)"`
	OpenId      string    `xorm:"not null VARCHAR(64)"`
	WxName string `xorm:"not null VARCHAR(255)"`
	WxAvatar string `xorm:"not null VARCHAR(255)"`
	Avatar     string    `xorm:"not null default '' comment('用户头像') VARCHAR(255)"`
	State      int       `xorm:"not null default 0 comment('用户账号状态,0-正常') TINYINT(4)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('账号创建时间') TIMESTAMP"`
}
