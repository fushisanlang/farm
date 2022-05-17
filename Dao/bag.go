/*
@Time : 2022/5/16 16:18
@Author : fushisanlang
@File : bag
@Software: GoLand
*/
package Dao

import (
	"farm/tools"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

//清理背包
func autoCleanBag() {
	sql := ` 
update bag set linkid = 0 ,groupid = 0 ,countnum =0 where  linkid = 0  or groupid = 0  or countnum =0;`
	_, err := g.DB().Exec(sql)

	tools.CheckErr(err)

}
func GetNullBagId(linkId int, groupId int) int {
	sql := `
SELECT
CASE WHEN
  ( SELECT keyid FROM bag WHERE linkid = 8 AND groupid = 1 ORDER BY keyid LIMIT 1 ) IS   NULL 
THEN
	( SELECT keyid FROM bag WHERE linkid = 0 AND groupid = 0 ORDER BY keyid LIMIT 1 ) 
ELSE 
	( SELECT keyid FROM bag WHERE linkid = 8 AND groupid = 1 ORDER BY keyid LIMIT 1 ) 
END keyid
`

	keyId, err := g.DB().GetOne(sql, linkId, groupId, linkId, groupId)
	tools.CheckErr(err)
	g.Log().Println("获取nllbagid开始---")
	g.Log().Println(linkId, groupId)
	g.Log().Println(gconv.Int(keyId["keyid"]))
	g.Log().Println("获取nllbagid结束---")

	return gconv.Int(keyId["keyid"])

}
