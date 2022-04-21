/*
@Time : 2022/4/19 4:21 下午
@Author : fushisanlang
@File : eradicate
@Software: GoLand
*/
package service

import (
	"farm/Dao"
	"farm/model"
	"github.com/gogf/gf/util/gconv"
)

func Harvest(id int, fieldId int) (bool, model.Harvest) {
	var HarvestInfo model.Harvest
	for true {
		Harvest := Dao.HarvestPlant(id, fieldId)

		HarvestInfo = model.Harvest{
			PlantId:   gconv.Int(Harvest["plantid"]),
			PlantName: gconv.String(Harvest["plantname"]),
			Yield:     gconv.Int(Harvest["yield"]),
		}
		if HarvestInfo.Yield != 0 {
			break
		}
	}
	bagId := Dao.GetHarvestBagId(HarvestInfo)
	if bagId == 0 {
		return false, HarvestInfo
	} else {
		Dao.HarvestBag(HarvestInfo, bagId)
		Dao.HarvestField(id, fieldId)
		Dao.HarvestEx(HarvestInfo)
	}
	return true, HarvestInfo
}
