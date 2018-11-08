package table_struct

import (
	"time"
)

type TUserSession struct {
	Id          int       `xorm:"not null pk autoincr INT(11)"`
	UserId      int       `xorm:"not null default 0 comment('用户id') index INT(11)"`
	SessionType int       `xorm:"not null default 0 comment('登陆类型,0-小程序登陆') TINYINT(4)"`
	Encryption  string    `xorm:"not null default '' comment('加密串') VARCHAR(255)"`
	Token       string    `xorm:"not null default '' comment('经过加密串加密的访问秘钥，一般3小时失效') VARCHAR(255)"`
	UpdateTime  time.Time `xorm:"not null default '0000-00-00 00:00:00' TIMESTAMP"`
}
