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
