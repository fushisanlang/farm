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
func FiledSummary() []model.FiledSummary {
	filedSummaryList := Dao.FiledSummary().List()
	var filedSummary model.FiledSummary
	lenList := len(filedSummaryList)
	FiledSummaryList := make([]model.FiledSummary, lenList, lenList)
	for i := 0; i < lenList; i++ {
		filedSummary.KeyId = gconv.Int(filedSummaryList[i]["keyid"])
		filedSummary.IsOpen = gconv.Int(filedSummaryList[i]["isopen"])
		filedSummary.DTime = gconv.Int(filedSummaryList[i]["duringtime"])
		filedSummary.NTime = gconv.Int(filedSummaryList[i]["timeneed"])
		FiledSummaryList[i] = filedSummary
	}
	return FiledSummaryList

}
