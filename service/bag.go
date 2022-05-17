/*
@Time : 2022/5/17 16:48
@Author : fushisanlang
@File : bag
@Software: GoLand
*/
package service

import "farm/Dao"

func GetNullBagId(linkId int, groupId int) int {
	NullBagId := Dao.GetNullBagId(linkId, groupId)
	return NullBagId
}
