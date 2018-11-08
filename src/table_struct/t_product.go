package table_struct

type TProduct struct {
	Id           int    `xorm:"not null pk autoincr INT(11)"`
	Name         string `xorm:"not null VARCHAR(255)"`
	ProductDescc string `xorm:"not null MEDIUMTEXT"`
	State        int    `xorm:"not null default 0 INT(11)"`
}
