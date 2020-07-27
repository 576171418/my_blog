package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Engine *xorm.Engine

func init() {
	engine, err := xorm.NewEngine("mysql", "root:sc666666@/myblog?charset=utf8" )

	if err != nil {
		panic(fmt.Sprintf("init database failed: %s", err.Error()))
	}
	Engine = engine
}
