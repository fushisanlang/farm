/*
@Time : 2022/4/11 9:25 上午
@Author : fushisanlang
@File : GetConf
@Software: GoLand
*/
package Dao

import (
	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/frame/g"
	_ "github.com/mattn/go-sqlite3"
)

func GetConf(confKey string) string {
	g.Log().Println("查询配置项" + confKey)
	//20220420
	//ConfValue, _ := g.DB().GetOne("select confValue from `conf` where  confKey = ?", confKey)
	var sql string
	switch confKey {
	case "version":
		sql = "select version confValue from `confnew` where  id = 1"
	case "WeightMin":
		sql = "select WeightMin confValue from `confnew` where  id = 1"

	case "HeightMin":
		sql = "select HeightMin confValue from `confnew` where  id = 1"

	}
	ConfValue, _ := g.DB().GetOne(sql)

	g.Log().Println("获得配置项" + confKey + gconv.String(ConfValue["confValue"]))

	return gconv.String(ConfValue["confValue"])

}
