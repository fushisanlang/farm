/*
@Time : 2022/4/19 4:21 下午
@Author : fushisanlang
@File : eradicate
@Software: GoLand
*/
package service

import "farm/Dao"

func Eradicate(id int, fieldId int) {
	Dao.Eradicate(id, fieldId)
}
