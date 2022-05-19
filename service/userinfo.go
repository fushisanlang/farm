/*
@Time : 2022/5/19 13:42
@Author : fushisanlang
@File : userinfo
@Software: GoLand
*/
package service

import (
	"farm/Dao"
	"github.com/gogf/gf/util/gconv"
	"time"
)

func VerifyUserInfoWithErase() bool {
	dbUserName := Dao.GetUserInfo("username")       // 道号
	dbUserAge := Dao.GetUserInfo("userage")         //道龄
	dbPetName := Dao.GetUserInfo("petname")         //灵宠
	dbFarmName := Dao.GetUserInfo("farmname")       // 道场名
	dbFarmClassId := Dao.GetUserInfo("farmclassid") //道场类型

	if dbUserName != "" && dbUserAge != "" && dbPetName != "" && dbFarmName != "" && dbFarmClassId != "" {

		return true
	} else {
		initData()

		return false
	}

}

func VerifyUserInfo() bool {
	dbUserName := Dao.GetUserInfo("username")       // 道号
	dbUserAge := Dao.GetUserInfo("userage")         //道龄
	dbPetName := Dao.GetUserInfo("petname")         //灵宠
	dbFarmName := Dao.GetUserInfo("farmname")       // 道场名
	dbFarmClassId := Dao.GetUserInfo("farmclassid") //道场类型
	if dbFarmClassId == "" {
		dbFarmClassId = "-1"

	}
	if dbUserAge == "" {
		dbUserAge = "-1"

	}
	UserInfo.UserName = dbUserName
	UserInfo.UserAge = gconv.Int(dbUserAge)
	UserInfo.PetName = dbPetName
	UserInfo.FarmName = dbFarmName
	UserInfo.FarmClassId = gconv.Int(dbFarmClassId)
	if dbUserName != "" && dbUserAge != "-1" && dbPetName != "" && dbFarmName != "" && dbFarmClassId != "-1" {

		return true
	} else {

		return false
	}

}

func VerifyOfflineTime() int {
	nowTime := gconv.Int(time.Now().Unix())
	lastTime := gconv.Int(Dao.GetUserInfo("lasttime"))
	return nowTime - lastTime
}
func RefreshUserInfo() {
	refreshUserMoney()
}
func refreshUserMoney() {
	UserMoney = gconv.Int(Dao.GetUserInfo("money"))

}
func RefreshUserLastTime() {
	now := gconv.Int(time.Now().Unix())
	Dao.RefreshUserLastTime(now)
}
