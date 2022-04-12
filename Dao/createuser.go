/*
@Time : 2022/4/12 4:47 下午
@Author : fushisanlang
@File : createuser
@Software: GoLand
*/
package Dao

import (
	"farm/tools"
	"github.com/gogf/gf/frame/g"
)

func CreateUser(userName string, userAge int, userSex int) {

	_, err := g.DB().Exec(`delete from conf where confKey in ('username','userage','usersex')`)
	tools.CheckErr(err)

	_, err = g.DB().Exec(`INSERT INTO "main"."conf"( "confKey", "confValue") VALUES ('username', ?),('userage', ?),('usersex', ?);`, userName, userAge, userSex)
	tools.CheckErr(err)
}
