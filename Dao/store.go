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
func BuyGoods(GoodsList model.BuyList, goodsCount int, money int) {
	//{"Id":8,"Name":"第八个","Price":10,"GroupId":4} 244 99
	//10*244*99/100
	keyId := GetNullBagId(GoodsList.Id, GoodsList.GroupId)
	g.Log().Println("---")

	g.Log().Println(goodsCount, GoodsList.Id, GoodsList.GroupId, keyId, money)
	g.Log().Println("---")
	//117 0 4 8 1193
	//244 8 4 8 2415
	//{"Id":3,"Name":"第三个","Price":10,"GroupId":4} 111 96
	sql := `update bag set countnum = countnum + ?, linkid = ?,groupid = ?  where keyid = ?;
		update userinfo set money = money -  ?;`
	_, err := g.DB().Exec(sql, goodsCount, GoodsList.Id, GoodsList.GroupId, keyId, money)

	tools.CheckErr(err)

	//autoCleanBag()
}
func SaleGoods(GoodsList model.GoodsList, goodsCount int, ratioSale int) {
	money := GoodsList.Price * ratioSale / 100 * goodsCount
	sql := `update bag set countnum = countnum - ? where keyid = ?;update userinfo set money = money +  ?;`
	_, err := g.DB().Exec(sql, goodsCount, GoodsList.KeyId, money)

	tools.CheckErr(err)

	autoCleanBag()
}

//每小时更新
//种子
func GetBuySeedList() gdb.Result {
	sql := ` SELECT id id,plantname name,price/10*yield price,1 groupid FROM plant where level <= (select level from userinfo ) ORDER BY RANDOM() limit 8;
`
	BuyGoodsList, err := g.DB().GetAll(sql)
	tools.CheckErr(err)
	return BuyGoodsList
}

//植物
func GetBuyPlantList() gdb.Result {
	sql := `  SELECT id id,plantname name,price price,2 groupid FROM plant where level <= (select level from userinfo ) ORDER BY RANDOM() limit 8;
`
	BuyPlantList, err := g.DB().GetAll(sql)
	tools.CheckErr(err)
	return BuyPlantList
}

//杂物。炼丹炉，药引子
func GetBuyGroceriesList() gdb.Result {
	sql := ` SELECT id id,name name,price price,4 groupid FROM groceries where level <= (select level from userinfo ) ORDER BY RANDOM() limit 8;`
	BuyOtherList, err := g.DB().GetAll(sql)
	tools.CheckErr(err)
	return BuyOtherList
}
