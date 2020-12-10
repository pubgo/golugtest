package db

import (
	"github.com/pubgo/golug/golug_plugin/plugins/golug_db"
	"xorm.io/xorm"
)

func GetDb() *xorm.Engine {
	return golug_db.GetClient()
}
