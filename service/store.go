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
	"github.com/gogf/gf/frame/g"
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
		goodsInfo.GroupId = gconv.Int(goodsAllList[i]["groupid"])

		GoodsAllList[i] = goodsInfo
	}
	return GoodsAllList
}

//BuyGoods

func BuyGoods(buyInfo model.BuyList, ratioBuy int) {
	/*
		{"Id":7,"Name":"第七个","Price":10,"GroupId":4} 116 101
		2022-05-17 18:00:06.971 ---
			2022-05-17 18:00:06.971 116 7 4 12 1171
		2022-05-17 18:00:06.971 ---
	*/
	GoodsName := buyInfo.Name
	//	money := GoodsList.Price * goodsCount * ratioSale / 100
	GoodsPrice := buyInfo.Price * ratioBuy / 100
	g.Log().Println("GoodsName:" + GoodsName)
	g.Log().Println("GoodsPrice:" + gconv.String(GoodsPrice))

	tools.CallClear()
	fmt.Println("现有晶石：" + gconv.String(UserMoney))
	fmt.Println("购买" + GoodsName + "，单价为" + gconv.String(GoodsPrice) + "晶石。")
	fmt.Println("可购买" + gconv.String(UserMoney/GoodsPrice) + "个。")
	fmt.Println("打算购买多少？")
	var a int
	for true {
		fmt.Println("请输入购买量，点击回车确认。")

		fmt.Scan(&a)
		//钱够
		if a == 0 {
			break
		} else if a*GoodsPrice <= UserMoney && a > 0 {
			fmt.Println("将购买" + gconv.String(a) + "个。")
			fmt.Println("总花费" + gconv.String(GoodsPrice*a) + "晶石。")
			buyGoods(buyInfo, a, a*GoodsPrice)
			break
			//不够
		} else if a*GoodsPrice > UserMoney {
			fmt.Println("你总可购买量为" + gconv.String(UserMoney/GoodsPrice) + "个，将全部购买。")
			a = UserMoney / GoodsPrice
			fmt.Println("总花费" + gconv.String(GoodsPrice*a) + "晶石。")
			buyGoods(buyInfo, a, a*GoodsPrice)
			break
			//退出
		} else {
			fmt.Println("输入异常。输入0退出购买页面")
		}
	}
	time.Sleep(3 * time.Second)
}
func buyGoods(buyList model.BuyList, goodsCount int, money int) {
	g.Log().Println(buyList, goodsCount, money)
	Dao.BuyGoods(buyList, goodsCount, money)
	ChangeMemUserInfoMoney()
}
func SaleGoods(GoodsList model.GoodsList, ratioSale int) {
	GoodsName := GoodsList.Name
	GoodsCount := GoodsList.Countnum
	GoodsPrice := GoodsList.Price * ratioSale / 100
	tools.CallClear()
	fmt.Println("现有晶石：" + gconv.String(UserMoney))
	fmt.Println("售卖" + GoodsName + "，单价为" + gconv.String(GoodsPrice) + "晶石，仓库中现有" + gconv.String(GoodsCount) + "个。")
	fmt.Println("打算出售多少？")
	var a int
	for true {
		fmt.Println("请输入出售量，点击回车确认。")

		fmt.Scan(&a)
		if a == 0 {
			break
		} else if a > 0 && a <= GoodsCount {
			fmt.Println("将卖出" + gconv.String(a) + "个。")
			fmt.Println("总收入" + gconv.String(GoodsList.Price*ratioSale/100*a) + "晶石。")

			saleGoods(GoodsList, a, ratioSale)

			break
		} else if a > GoodsCount {
			fmt.Println("共计" + gconv.String(GoodsCount) + "个，将全部卖出。")
			a = GoodsCount

			fmt.Println("总收入" + gconv.String(GoodsList.Price*ratioSale/100*a) + "晶石。")

			saleGoods(GoodsList, GoodsCount, ratioSale)

			break
		} else {
			fmt.Println("输入异常。输入0退出售卖页面")
		}
	}
	time.Sleep(3 * time.Second)
}
func saleGoods(GoodsList model.GoodsList, goodsCount int, ratioSale int) {

	Dao.SaleGoods(GoodsList, goodsCount, ratioSale)
	ChangeMemUserInfoMoney()
}

//定时更新商店列表
func GetBuySeedList() {
	buyList := Dao.GetBuySeedList().List()
	var buyInfo model.BuyList
	lenList := len(buyList)
	BuyList := make([]model.BuyList, lenList, lenList)
	for i := 0; i < lenList; i++ {
		buyInfo.Id = gconv.Int(buyList[i]["id"])
		buyInfo.Name = gconv.String(buyList[i]["name"])
		buyInfo.Price = gconv.Int(buyList[i]["price"])
		buyInfo.GroupId = 1

		BuyList[i] = buyInfo
	}
	BuySeedList = BuyList
}

//植物
func GetBuyPlantList() {
	buyList := Dao.GetBuyPlantList().List()
	var buyInfo model.BuyList
	lenList := len(buyList)
	BuyList := make([]model.BuyList, lenList, lenList)
	for i := 0; i < lenList; i++ {
		buyInfo.Id = gconv.Int(buyList[i]["id"])
		buyInfo.Name = gconv.String(buyList[i]["name"])
		buyInfo.Price = gconv.Int(buyList[i]["price"])
		buyInfo.GroupId = 2

		BuyList[i] = buyInfo
	}
	BuyPlantList = BuyList
}

//杂物。炼丹炉，药引子
func GetBuyGroceriesList() {
	buyList := Dao.GetBuyGroceriesList().List()
	var buyInfo model.BuyList
	lenList := len(buyList)
	BuyList := make([]model.BuyList, lenList, lenList)
	for i := 0; i < lenList; i++ {
		buyInfo.Id = gconv.Int(buyList[i]["id"])
		buyInfo.Name = gconv.String(buyList[i]["name"])
		buyInfo.Price = gconv.Int(buyList[i]["price"])
		buyInfo.GroupId = 4

		BuyList[i] = buyInfo
	}
	BuyGroceriesList = BuyList
}
