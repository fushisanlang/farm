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

func GetFieldInfo(id int) gdb.Result {

	g.Log().Println("拉取土地列表")
	sql := `SELECT f.id id, f.fieldid fieldid, p.plantname name, f.duringtime duringtime, p.timeneed timeneed, f.isopen isopen FROM field f LEFT OUTER JOIN plant p ON f.plantid = p.id where f.isopen = 1 and f.id = ?;`
	Fields, _ := g.DB().GetAll(sql, id)
	return Fields

}

func GetFieldInfoAll() gdb.Result {

	g.Log().Println("拉取土地列表")

	sql := "select id,count(case when  isopen =1 then id else null end) isopencount ,count(case when  isopen =1 and plantid != 0 then id else null end) plantcount from field GROUP BY id;"
	FieldInfoAll, _ := g.DB().GetAll(sql)
	return FieldInfoAll

}
