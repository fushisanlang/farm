/*
@Time : 2022/4/11 9:25 上午
@Author : fushisanlang
@File : GetConf
@Software: GoLand
*/
package Dao

import (
	"farm/model"
	"farm/tools"
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
	b.groupid groupid,
	'作物' type 
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
	b.groupid groupid,
	'丹药' type 
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

func SaleGoods(GoodsList model.GoodsList, goodsCount int, ratioSale int) {
	money := GoodsList.Price * goodsCount * ratioSale / 100
	sql := `update bag set countnum = countnum - ? where keyid = ?;update userinfo set money = money +  ?;`
	_, err := g.DB().Exec(sql, goodsCount, GoodsList.KeyId, money)

	tools.CheckErr(err)

	autoCleanBag()
}
