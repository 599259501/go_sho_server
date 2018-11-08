package table_struct

type TProductValue struct {
	Id     int    `xorm:"not null pk autoincr INT(11)"`
	AttrId int    `xorm:"not null INT(11)"`
	Value  string `xorm:"not null VARCHAR(255)"`
}
