/*
@Time : 2022/4/14 4:12 下午
@Author : fushisanlang
@File : field
@Software: GoLand
*/
package service

import (
	"farm/Dao"
	"farm/model"
	"github.com/gogf/gf/util/gconv"
)

func GetSeedList() []model.SeedList {
	seedAllList := Dao.GetSeedList().List()
	var seedInfo model.SeedList
	lenList := len(seedAllList)
	SeedAllList := make([]model.SeedList, lenList, lenList)
	for i := 0; i < lenList; i++ {
		seedInfo.Id = gconv.Int(seedAllList[i]["id"])
		seedInfo.PlantName = gconv.String(seedAllList[i]["plantname"])
		seedInfo.TimeNeed = gconv.Int(seedAllList[i]["timeneed"])
		seedInfo.CountNum = gconv.Int(seedAllList[i]["countnum"])
		seedInfo.Price = gconv.Int(seedAllList[i]["price"])
		seedInfo.Ex = gconv.Int(seedAllList[i]["ex"])
		SeedAllList[i] = seedInfo
	}
	return SeedAllList
}
