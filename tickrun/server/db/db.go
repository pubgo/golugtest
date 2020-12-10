package db

import (
	"github.com/pubgo/golugin/golug_db"
	"xorm.io/xorm"
)

func GetDb() *xorm.Engine {
	return golug_db.GetClient()
}
