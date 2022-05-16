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
)

//清理背包
func autoCleanBag() {
	sql := `update bag set linkid = 0 ,groupid = 0 ,countnum =0 where  linkid = 0  or groupid = 0  or countnum =0;`
	_, err := g.DB().Exec(sql)

	tools.CheckErr(err)

}
