/*
@Time : 2022/4/13 10:48 上午
@Author : fushisanlang
@File : test
@Software: GoLand
*/
package service

import (
	"farm/Dao"
	"farm/model"
	"farm/tools"
	"fmt"
	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/util/gconv"
)

func GetPetName() {
	g.Log().Println("读取灵宠")
	for true {
		fmt.Println("输入灵宠，按回车确认")

		fmt.Scan(&UserInfo.PetName)
		if UserInfo.PetName != "" {
			Dao.UpdateUserInfo("petname", UserInfo.PetName)
			break
		} else {
			tools.CallClear()
			fmt.Println("请勿输入空值")
		}
	}

}

func GetUserName() {
	g.Log().Println("读取道号")
	for true {
		fmt.Println("输入道号，按回车确认")

		fmt.Scan(&UserInfo.UserName)
		if UserInfo.UserName != "" {
			Dao.UpdateUserInfo("username", UserInfo.UserName)
			break
		} else {
			tools.CallClear()
			fmt.Println("请勿输入空值")
		}
	}

}

func GetUserAge() {
	g.Log().Println("读取道龄")
	for true {
		fmt.Println("输入道龄，按回车确认")

		fmt.Scan(&UserInfo.UserAge)

		if UserInfo.UserAge > 0 {
			Dao.UpdateUserInfo("userage", gconv.String(UserInfo.UserAge))
			break
		} else {
			tools.CallClear()
			fmt.Println("请输入大于0的数")
		}
	}

}
func GetFarmName() {
	g.Log().Println("读取道场")
	for true {
		fmt.Println("输入道场，按回车确认")

		fmt.Scan(&UserInfo.FarmName)
		if UserInfo.FarmName != "" {
			Dao.UpdateUserInfo("farmname", UserInfo.FarmName)

			break
		} else {
			tools.CallClear()
			fmt.Println("请勿输入空值")
		}
	}

}
func GetFarmClass() []model.FarmClass {

	farmClassList := Dao.GetFarmClass().List()
	var farmClass model.FarmClass
	lenList := len(farmClassList)
	FarmClassList := make([]model.FarmClass, lenList, lenList)
	for i := 0; i < lenList; i++ {
		farmClass.FarmClassId = gconv.Int(farmClassList[i]["farmclassid"])
		farmClass.FarmClassName = gconv.String(farmClassList[i]["farmclassname"])
		farmClass.FarmClassInfo = gconv.String(farmClassList[i]["farmclassinfo"])
		FarmClassList[i] = farmClass
	}
	return FarmClassList

}

func GetFarmClassId() {
	g.Log().Println("读取福地")
	farmClassList := GetFarmClass()
	listLen := len(farmClassList)

	for true {
		for i := 0; i < listLen; i++ {
			fmt.Printf(gconv.String(farmClassList[i].FarmClassId) + ":" + farmClassList[i].FarmClassName + "\n")
			fmt.Println(farmClassList[i].FarmClassInfo)
		}
		fmt.Println("输入福地的编号，按回车确认")
		var userInput int
		fmt.Scan(&userInput)

		if userInput > 0 && userInput < listLen {
			UserInfo.FarmClassId = userInput - 1
			Dao.UpdateUserInfo("farmclassid", gconv.String(UserInfo.FarmClassId))

			break
		} else {
			tools.CallClear()
			fmt.Println("请输入正确的福地编号")
		}
	}

}
