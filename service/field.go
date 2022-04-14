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
		fieldInfo.AllDuringTime = gconv.Int(fieldInfoList[i]["allduringtime"])
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

		FieldInfoAllList[i] = fieldInfo
	}
	return FieldInfoAllList
}
