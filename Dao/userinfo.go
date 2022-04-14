/*
@Time : 2022/4/13 3:20 下午
@Author : fushisanlang
@File : getuserinfo
@Software: GoLand
*/
package Dao

import (
	"farm/tools"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func GetUserInfo(infoKey string) string {
	g.Log().Println("拉取用户信息 key：" + infoKey)
	InfoValue, _ := g.DB().GetOne("select infoValue from `userinfo` where  infoKey = ?", infoKey)
	return gconv.String(InfoValue["infoValue"])

}

func UpdateUserInfo(infoKey string, infoValue string) {
	g.Log().Println("更新用户信息 key：" + infoKey + ", value" + infoValue)

	_, err := g.DB().Exec(`UPDATE "main"."userinfo" SET "infoValue" = ? WHERE "infoKey" = ?;`, infoValue, infoKey)
	tools.CheckErr(err)

}
