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

func GetOneFieldInfo(id int, fieldId int) model.FieldInfo {
	oneFieldInfo := Dao.GetOneFieldInfo(id, fieldId)
	var OneFieldInfo model.FieldInfo
	OneFieldInfo.Id = id
	OneFieldInfo.FieldId = fieldId
	OneFieldInfo.IsOpen = gconv.Int(oneFieldInfo["isopen"])
	OneFieldInfo.PlantName = gconv.String(oneFieldInfo["plantime"])
	OneFieldInfo.DuringTime = gconv.Int(oneFieldInfo["duringtime"])
	OneFieldInfo.NeedTime = gconv.Int(oneFieldInfo["timeneed"])

	return OneFieldInfo

}

func GetFieldInfo(id int) []model.FieldInfo {
	fieldInfoList := Dao.GetFieldInfo(id).List()
	var fieldInfo model.FieldInfo
	lenList := len(fieldInfoList)
	FieldInfoList := make([]model.FieldInfo, lenList, lenList)
	for i := 0; i < lenList; i++ {
		fieldInfo.Id = gconv.Int(fieldInfoList[i]["id"])
		fieldInfo.FieldId = gconv.Int(fieldInfoList[i]["fieldid"])
		fieldInfo.PlantName = gconv.String(fieldInfoList[i]["name"])
		fieldInfo.DuringTime = gconv.Int(fieldInfoList[i]["duringtime"])
		fieldInfo.NeedTime = gconv.Int(fieldInfoList[i]["timeneed"])
		fieldInfo.IsOpen = gconv.Int(fieldInfoList[i]["isopen"])
		FieldInfoList[i] = fieldInfo
	}
	return FieldInfoList
}
func GetFieldInfoAll() []model.FieldInfoAll {
	fieldInfoAllList := Dao.GetFieldInfoAll().List()
	var fieldInfo model.FieldInfoAll
	lenList := len(fieldInfoAllList)
	FieldInfoAllList := make([]model.FieldInfoAll, lenList, lenList)
	for i := 0; i < lenList; i++ {
		fieldInfo.Id = gconv.Int(fieldInfoAllList[i]["id"])
		fieldInfo.IsOpenCount = gconv.Int(fieldInfoAllList[i]["isopencount"])
		fieldInfo.PlantCount = gconv.Int(fieldInfoAllList[i]["plantcount"])
		//fieldInfo.IsRipeCount = gconv.Int(fieldInfoAllList[i]["isripecount"])
		//fieldInfo.WillRipeCount = gconv.Int(fieldInfoAllList[i]["willripecount"])

		FieldInfoAllList[i] = fieldInfo
	}
	return FieldInfoAllList
}
func OpenField(id int, fieldId int) (bool, int) {
	return Dao.OpenField(id, fieldId)
}
