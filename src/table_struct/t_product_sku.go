package table_struct

type TProductSku struct {
	Id        int    `xorm:"not null pk autoincr INT(11)"`
	ProductId int    `xorm:"not null INT(11)"`
	Price     string `xorm:"not null DECIMAL(10,2)"`
	Stock     int    `xorm:"not null default 0 INT(11)"`
	AttrStr   string `xorm:"not null VARCHAR(255)"`
}
