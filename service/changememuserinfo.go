/*
@Time : 2022/5/17 15:35
@Author : fushisanlang
@File : changeuserinfo
@Software: GoLand
*/
package service

import (
	"farm/Dao"
	"github.com/gogf/gf/util/gconv"
)

func ChangeMemUserInfoMoney() {
	UserMoney = gconv.Int(Dao.GetUserInfo("money"))

}
