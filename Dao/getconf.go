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

	ConfValue, _ := g.DB().GetOne("select confValue from `conf` where  confKey = ?", confKey)
	return gconv.String(ConfValue["confValue"])

}
