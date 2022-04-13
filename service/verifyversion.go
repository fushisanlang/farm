/*
@Time : 2022/4/11 10:16 上午
@Author : fushisanlang
@File : verifyscreen
@Software: GoLand
*/
package service

import (
	"github.com/gogf/gf/frame/g"
)

func VerifyVersion() {

	DbVersion := GetConf("version")
	g.Log().Println("校验版本")

	if DbVersion != Version {
		updateDbVersion()
	}
}
