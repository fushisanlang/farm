/*
@Time : 2022/4/11 10:03 上午
@Author : fushisanlang
@File : GetDbVersion
@Software: GoLand
*/
package service

import "farm/Dao"

func GetDbVersion() string {
	DBVersion := Dao.GetDbVersion()
	return DBVersion
}