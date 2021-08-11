package model

import "time"

const (
	CONTACT_CATE_USER     = 0x01
	CONTACT_CATE_COMUNITY = 0x02
)

type Contact struct {
	Id       int64     `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	OwnerId  int64     `xorm:"bigint(20)" form:"ownerid" json:"ownerid"`
	DstObj   int64     `xorm:"bigint(20)" form:"dstobj" json:"dstobj"`
	Cate     int       `xorm:"int(11)" form:"cate" json:"cate"`
	Memo     string    `xorm:"varchar(120)" form:"memo" json:"memo"`
	CreateAt time.Time `xorm:"datetime" form:"createat" json:"createat"`
}
