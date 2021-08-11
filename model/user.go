package model

import "time"

const (
	SEX_WOMEN  = "W"
	SEX_MEN    = "M"
	SEX_UNKNOW = "U"
)

type User struct {
	Id        int64     `xorm:"pk autoincr bigint(64)" form:"id" json:"id"`
	Name      string    `xorm:"varchar(20)" form:"name" json:"name"`
	Passwd    string    `xorm:"varchar(40)" form:"passwd" json:"-"`
	Telephone string    `xorm:"varchar(20)" form:"telephone" json:"telephone"`
	Avatar    string    `xorm:"varchar(150)" form:"avatar" json:"avatar"`
	Sex       string    `xorm:"varchar(2)" form:"sex" json:"sex"`
	Salt      string    `xorm:"varchar(10)" form:"salt" json:"-"`
	Online    int       `xorm:"int(10)" form:"online" json:"online"`
	Token     string    `xorm:"varchar(40)" form:"token" json:"token"`
	Memo      string    `xorm:"varchar(140)" form:"memo" json:"memo"`
	CreateAt  time.Time `xorm:"datetime" form:"createat" json:"createat"`
}
