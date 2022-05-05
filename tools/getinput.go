/*
@Time : 2022/4/22 10:24
@Author : fushisanlang
@File : getinput
@Software: GoLand
*/
package tools

import (
	"fmt"
)

func GetInputNum() int {
	CallClear()

	var a int
	for true {
		fmt.Println("输入道龄，按回车确认")

		fmt.Scan(&a)

		if a >= 0 {
			break
		} else {
			CallClear()
			fmt.Println("请输入不小于0的数")
		}
	}
	return a
}
