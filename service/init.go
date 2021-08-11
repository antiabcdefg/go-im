package service

import (
	"go-im/model"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var DbEngin *xorm.Engine

func init() {
	drivername := "mysql"
	DsName := "root:root@(localhost:3306)/chat?charset=utf8"

	err := errors.New("")
	DbEngin, err = xorm.NewEngine(drivername, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}

	DbEngin.ShowSQL(true)
	DbEngin.SetMaxOpenConns(2)
	DbEngin.Sync2(new(model.User), new(model.Contact), new(model.Community))

	fmt.Println("init database ok")
}
