/*
@Time : 2022/4/11 9:25 上午
@Author : fushisanlang
@File : GetConf
@Software: GoLand
*/
package Dao

import (
	"farm/tools"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/frame/g"
	_ "github.com/mattn/go-sqlite3"
)

func GetOneFieldInfo(id int, fieldId int) gdb.Record {
	g.Log().Println("拉取土地信息:" + gconv.String(id) + "  " + gconv.String(fieldId))
	sql := `select p.plantname plantime, f.duringtime duringtime, p.timeneed timeneed, f.isopen isopen from field f LEFT JOIN plant p on f.plantid = p.id where f.id = ? and f.fieldid = ?`
	OneFieldInfo, _ := g.DB().GetOne(sql, id, fieldId)
	return OneFieldInfo
}
func GetFieldInfo(id int) gdb.Result {

	g.Log().Println("拉取土地列表:" + gconv.String(id))
	sql := `SELECT f.id id, f.fieldid fieldid, p.plantname name, f.duringtime duringtime, p.timeneed timeneed, f.isopen isopen FROM field f LEFT OUTER JOIN plant p ON f.plantid = p.id where f.id = ?;`
	Fields, _ := g.DB().GetAll(sql, id)
	return Fields

}

func GetFieldInfoAll() gdb.Result {

	g.Log().Println("拉取土地列表")

	sql := `SELECT
	f.id,
	count( CASE WHEN f.isopen = 1 THEN f.id ELSE NULL END ) isopencount,
	count( CASE WHEN f.isopen = 1 AND f.plantid != 0 THEN f.id ELSE NULL END )  plantcount 
FROM
	field f 
GROUP BY
	f.id;`

	FieldInfoAll, err := g.DB().GetAll(sql)
	tools.CheckErr(err)

	return FieldInfoAll

}
func OpenField(id int, fieldId int) (bool, int) {
	//20220420
	//sql := `select (select infoValue from userinfo where infoKey='money')-(select infoValue from userinfo where infoKey='nextfieldneedmoney') isMoneyEnough`
	sql := `select money - nextfieldneed isMoneyEnough from userinfonew;`

	isMoneyEnough, _ := g.DB().GetOne(sql)
	IsMoneyEnough := gconv.Int(isMoneyEnough["isMoneyEnough"])

	if IsMoneyEnough >= 0 {
		//20220420
		//sql2 := `UPDATE "main"."field" SET  "isopen" = 1 WHERE "id" = ? AND "fieldid" = ? ;UPDATE "main"."userinfo" SET "infoValue" = infoValue-(select infoValue from userinfo where infoKey='nextfieldneedmoney') WHERE "infoKey" = 'money';UPDATE "main"."userinfo" SET "infoValue" =  "infoValue" * 2  WHERE "infoKey" = 'nextfieldneedmoney';`
		sql2 := `UPDATE "main"."field" SET  "isopen" = 1 WHERE "id" = ? AND "fieldid" = ? ;update userinfonew set money = money - nextfieldneed ,nextfieldneed =nextfieldneed * 2 where id = 1`
		_, err := g.DB().Exec(sql2, id, fieldId)
		tools.CheckErr(err)
		return true, IsMoneyEnough
	} else {
		return false, IsMoneyEnough
	}
}
