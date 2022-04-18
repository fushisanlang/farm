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

func GetSeedList() gdb.Result {
	sql := `
SELECT
	b.id id,
	p.plantname plantname,
	p.timeneed timeneed,
	b.countnum countnum,
	p.price price,
	p.ex ex 
FROM
	bag b,
	plant p 
WHERE
	b.groupid = 1 
	AND b.countnum > 0 
	AND p.level <= ( SELECT infoValue FROM userinfo WHERE infoKey = 'level' ) 
	AND b.linkid = p.id
`
	SeedList, _ := g.DB().GetAll(sql)
	return SeedList

}
