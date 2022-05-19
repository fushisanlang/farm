/*
@Time : 2022/4/19 3:12 下午
@Author : fushisanlang
@File : plant
@Software: GoLand
*/
package service

import "farm/Dao"

func DuringTimeAutoAdd() {
	Dao.DuringTimeAutoAdd()
}

func DuringTimeAdd(duringTime int) {
	Dao.DuringTimeAdd(duringTime)
}
