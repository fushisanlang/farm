/*
@Time : 2022/4/21 10:10
@Author : fushisanlang
@File : store
@Software: GoLand
*/
package service

import (
	"farm/Dao"
	"farm/model"
	"farm/tools"
	"fmt"
	"math/rand"
	"time"

	"github.com/gogf/gf/util/gconv"
)

func ChangeRatioAuto() {
	changeRatioBuyAuto()
	changeRatioSaleAuto()
}
func changeRatioBuyAuto() {
	rand.Seed(time.Now().UnixNano())
	RatioBuy = 90 + rand.Float32()*20

}
func changeRatioSaleAuto() {
	rand.Seed(time.Now().UnixNano())
	RatioSale = 90 + rand.Float32()*20

}
func GetGoodsList() []model.GoodsList {
	goodsAllList := Dao.GetGoodsList().List()
	var goodsInfo model.GoodsList
	lenList := len(goodsAllList)
	GoodsAllList := make([]model.GoodsList, lenList, lenList)
	for i := 0; i < lenList; i++ {
		goodsInfo.KeyId = gconv.Int(goodsAllList[i]["keyid"])
		goodsInfo.Name = gconv.String(goodsAllList[i]["name"])
		goodsInfo.Countnum = gconv.Int(goodsAllList[i]["countnum"])
		goodsInfo.Price = gconv.Int(goodsAllList[i]["price"])
		goodsInfo.Groupid = gconv.Int(goodsAllList[i]["groupid"])

		GoodsAllList[i] = goodsInfo
	}
	return GoodsAllList
}

func SaleGoods(GoodsList model.GoodsList, ratioSale int) {
	GoodsName := GoodsList.Name
	GoodsCount := GoodsList.Countnum
	GoodsPrice := GoodsList.Price * ratioSale / 100
	tools.CallClear()
	fmt.Println("售卖" + GoodsName + "，单价为" + gconv.String(GoodsPrice) + "晶石，仓库中现有" + gconv.String(GoodsCount) + "个。")
	fmt.Println("打算出售多少？")
	var a int
	for true {
		fmt.Println("请输入出售量，点击回车确认。")

		fmt.Scan(&a)

		if a > 0 && a <= GoodsCount {
			fmt.Println("将卖出" + gconv.String(a) + "个。")
			saleGoods(GoodsList, a, ratioSale)

			break
		} else if a > GoodsCount {
			fmt.Println("共计" + gconv.String(GoodsCount) + "个，将全部卖出。")
			saleGoods(GoodsList, GoodsCount, ratioSale)

			a = GoodsCount
			break
		} else if a == -1 {
			break

		} else {
			fmt.Println("输入异常。输入-1退出售卖页面")
		}
	}

}
func saleGoods(GoodsList model.GoodsList, goodsCount int, ratioSale int) {
	Dao.SaleGoods(GoodsList, goodsCount, ratioSale)
}
