/*
@Time : 2022/4/12 1:52 下午
@Author : fushisanlang
@File : userinfo
@Software: GoLand
*/
package model

type UserInfo struct {
	UserName string
	UserAge  int

	PetName     string
	FarmName    string
	FarmClassId int
}
type FarmClass struct {
	FarmClassId   int
	FarmClassName string
	FarmClassInfo string
}
type FieldInfo struct {
	Id         int
	FieldId    int
	PlantName  string
	DuringTime int
	IsOpen     int
	NeedTime   int
}
type FieldInfoAll struct {
	Id          int
	IsOpenCount int
	PlantCount  int
}
