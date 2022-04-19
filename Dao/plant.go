/*
@Time : 2022/4/19 3:12 下午
@Author : fushisanlang
@File : plant
@Software: GoLand
*/
package Dao

import (
	"farm/tools"
	"github.com/gogf/gf/frame/g"
)

func DuringTimeAutoAdd() {
	sql := `UPDATE field 
SET duringtime = duringtime + 1 
WHERE
	keyid IN ( SELECT keyid FROM field f, plant p WHERE f.isopen = 1 AND f.plantid = p.id AND f.duringtime < p.timeneed )`

	_, err := g.DB().Exec(sql)
	tools.CheckErr(err)
}
