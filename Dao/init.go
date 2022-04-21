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
	//eraseBagTable()
	//giveSeed()
	openField()
}
func eraseUserTable() {
	sql := `UPDATE "main"."userinfo" SET "username" = NULL,"userage" = NULL, "farmname" = NULL, "farmclassid" = NULL, "petname" = NULL, "money" = 1000, "ex" = 0, "level" = 0, "nextfieldneed" = 500 WHERE "id" = 1;`
	_, err := g.DB().Exec(sql)
	tools.CheckErr(err)
}
func eraseFieldTable() {
	sql := "UPDATE field SET isopen = 0 , plantid =0 , duringtime =0,timeneed = 0"
	_, err := g.DB().Exec(sql)
	tools.CheckErr(err)
}
func eraseBagTable() {
	sql := "UPDATE bag SET linkid = 0 , groupid =0 , countnum =0"
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
