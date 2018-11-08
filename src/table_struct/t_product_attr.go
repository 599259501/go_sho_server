package table_struct

type TProductAttr struct {
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	Name       string `xorm:"not null VARCHAR(255)"`
	CategoryId int    `xorm:"not null comment('分类id') INT(11)"`
}
