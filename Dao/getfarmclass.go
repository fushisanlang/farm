/*
@Time : 2022/4/11 9:25 上午
@Author : fushisanlang
@File : GetConf
@Software: GoLand
*/
package Dao

import (
	"github.com/gogf/gf/database/gdb"

	"github.com/gogf/gf/frame/g"
	_ "github.com/mattn/go-sqlite3"
)

func GetFarmClass() gdb.Result {

	ConfValue, _ := g.DB().GetAll("select * from farmclass")
	return ConfValue

}
