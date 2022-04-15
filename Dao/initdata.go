/*
@Time : 2022/4/15 2:30 下午
@Author : fushisanlang
@File : erasehisinfo
@Software: GoLand
*/
package Dao

import (
	"farm/tools"
	"github.com/gogf/gf/frame/g"
)

func InitData() {
	g.Log().Println("InitData")
	eraseUserTable()
	eraseFieldTable()
	eraseBagTable()
	giveMoney()
	giveSeed()
	setUserLevel()
	openField()
}
func eraseUserTable() {
	sql := "UPDATE userinfo SET infoValue = ''"
	_, err := g.DB().Exec(sql)
	tools.CheckErr(err)
}
func eraseFieldTable() {
	sql := "UPDATE field SET isopen = 0 and plantid =0 and duringtime =0"
	_, err := g.DB().Exec(sql)
	tools.CheckErr(err)
}
func eraseBagTable() {
	sql := "UPDATE bag SET linkid = 0 and groupid =0 and countnum =0"
	_, err := g.DB().Exec(sql)
	tools.CheckErr(err)
}

func setUserLevel() {
	sql := `UPDATE "userinfo" SET  "infoValue" = 0 WHERE "infoKey" = 'ex' or "infoKey" = 'level'  ;`
	_, err := g.DB().Exec(sql)
	tools.CheckErr(err)
}

func giveMoney() {
	sql := `UPDATE "userinfo" SET  "infoValue" = '1000' WHERE "infoKey" = 'money';`
	_, err := g.DB().Exec(sql)
	tools.CheckErr(err)
}

func giveSeed() {
	sql := `UPDATE "bag" SET  "linkid" = 1, "groupid" = 1, "countnum" = 3 WHERE "id" = 1 AND "bagid" = 1 ;`
	_, err := g.DB().Exec(sql)
	tools.CheckErr(err)
}
func openField() {
	sql := `update field set "isopen" = 1 where "id"=1 and "fieldid" in (1,2,3)`
	_, err := g.DB().Exec(sql)
	tools.CheckErr(err)
}
