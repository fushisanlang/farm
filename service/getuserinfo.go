/*
@Time : 2022/4/11 10:03 上午
@Author : fushisanlang
@File : GetDbVersion
@Software: GoLand
*/
package service

import (
	"farm/Dao"
	"farm/model"
	"github.com/gogf/gf/util/gconv"
)

func GetUserInfo() model.UserInfo {

	var UserInfo model.UserInfo

	dbUserName := Dao.GetConf("username")
	dbUserSex := Dao.GetConf("usersex")
	dbUserAge := Dao.GetConf("userage")
	if dbUserName != "" && dbUserSex != "" && dbUserAge != "" {
		UserInfo.UserName = dbUserName
		UserInfo.UserSex = gconv.Int(dbUserSex)
		UserInfo.UserAge = gconv.Int(dbUserAge)
	} else {
		CreateUser()
	}
	return UserInfo
}
