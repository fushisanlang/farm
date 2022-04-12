/*
@Time : 2022/4/12 3:04 下午
@Author : fushisanlang
@File : new
@Software: GoLand
*/
package service

import (
	"farm/Dao"
	"farm/tools"
	"fmt"
)

func CreateUser() {
	var UserAge, UserSex int
	var UserName, UserSexInput string
	//
	for true {
		fmt.Println("输入姓名，按回车确认")

		fmt.Scan(&UserName)
		if UserName != "" {
			break
		} else {
			tools.CallClear()
			fmt.Println("请勿输入空值")
		}
	}
	for true {
		fmt.Println("输入年龄，按回车确认")

		fmt.Scan(&UserAge)
		if UserAge > 0 {
			break
		} else {
			tools.CallClear()
			fmt.Println("请输入大于0的数")
		}
	}
	for true {
		fmt.Println("输入性别，按回车确认")

		fmt.Scan(&UserSexInput)
		if UserSexInput == "男" {
			UserSex = 1
			break
		} else if UserSexInput == "女" {
			UserSex = 0
			break
		} else {
			tools.CallClear()
			fmt.Println("请勿输入男，女之外的值")
		}
	}

	Dao.CreateUser(UserName, UserAge, UserSex)
}
