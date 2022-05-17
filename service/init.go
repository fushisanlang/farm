/*
@Time : 2022/4/11 9:14 上午
@Author : fushisanlang
@File : init
@Software: GoLand
*/
package service

import (
	"farm/Dao"
	"farm/model"
	"github.com/gogf/gf/util/gconv"
)

var UserInfo model.UserInfo
var Version string
var HeightMin, WeightMin int
var DbVersion string
var RatioBuy, RatioSale float32
var BuySeedList, BuyPlantList, BuyGroceriesList []model.BuyList
var UserMoney int

func init() {
	RatioBuy, RatioSale = 1, 1
	Version = "0.0.1"
	DbVersion = GetConf("version")
	WeightMin = gconv.Int(GetConf("WeightMin"))
	HeightMin = gconv.Int(GetConf("HeightMin"))
	UserMoney = gconv.Int(Dao.GetUserInfo("money"))
}
