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
	"github.com/gogf/gf/util/gconv"
	_ "github.com/mattn/go-sqlite3"
)

func GetGoodsList() gdb.Result {
	sql := `
SELECT
	b.keyid keyid,
	p.plantname name,
	b.countnum countnum,
	p.price price,
	b.groupid groupid
FROM
	bag b,
	plant p 
WHERE
	b.linkid = p.id 
	AND b.groupid = 2 UNION ALL
SELECT
	b.keyid keyid,
	d.drugname name,
	b.countnum countnum,
	d.price price,
	b.groupid groupid
	FROM
	bag b,
	drug d 
WHERE
	b.linkid = d.id 
	AND b.groupid = 3;
`
	GoodsList, _ := g.DB().GetAll(sql)

	g.Log().Println("GoodsList " + gconv.String(GoodsList))

	return GoodsList

}
