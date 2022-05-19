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
	var sql string

	switch infoKey {
	case "username":
		sql = `select username infoValue from userinfo where  id = 1;`
	case "petname":
		sql = `select petname infoValue from userinfo where  id = 1;`
	case "farmname":
		sql = `select farmname infoValue from userinfo where  id = 1;`

	case "farmclassid":
		sql = `select farmclassid infoValue from userinfo where  id = 1;`
	case "money":
		sql = `select money infoValue from userinfo where  id = 1;`
	case "userage":
		sql = `select userage infoValue from userinfo where  id = 1;`
	case "lasttime":
		sql = `select lasttime infoValue from userinfo where  id = 1;`
	}

	InfoValue, _ := g.DB().GetOne(sql)

	return gconv.String(InfoValue["infoValue"])

}

func UpdateUserInfo(infoKey string, infoValue string) {
	g.Log().Println("更新用户信息 key：" + infoKey + ", value" + infoValue)

	var sql string
	switch infoKey {
	case "username":
		sql = `UPDATE "main"."userinfo" SET username= ? WHERE "id" = 1;`
	case "petname":
		sql = `UPDATE "main"."userinfo" SET petname= ? WHERE "id" = 1;`
	case "farmname":
		sql = `UPDATE "main"."userinfo" SET farmname= ? WHERE "id" = 1;`

	case "farmclassid":
		sql = `UPDATE "main"."userinfo" SET farmclassid= ? WHERE "id" = 1;`

	case "userage":
		sql = `UPDATE "main"."userinfo" SET userage= ? WHERE "id" = 1;`

	}

	_, err := g.DB().Exec(sql, infoValue)
	tools.CheckErr(err)

}
func RefreshUserLastTime(nowTime int) {
	sql := `update userinfo set lasttime = ?`
	_, err := g.DB().Exec(sql, nowTime)
	tools.CheckErr(err)
}
