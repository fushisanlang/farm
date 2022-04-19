/*
@Time : 2022/4/19 4:22 下午
@Author : fushisanlang
@File : eradicate
@Software: GoLand
*/
package Dao

import (
	"farm/tools"
	"github.com/gogf/gf/frame/g"
)

func Eradicate(id int, fieldId int) {
	sql := `UPDATE "main"."field" 
SET "plantid" = 0,
"duringtime" = 0.0 
WHERE
	"id" = ? 
	AND "fieldid" = ?`

	_, err := g.DB().Exec(sql, id, fieldId)
	tools.CheckErr(err)
}
