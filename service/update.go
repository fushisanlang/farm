/*
@Time : 2022/4/11 10:10 上午
@Author : fushisanlang
@File : update
@Software: GoLand
*/
package service

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
)

func updateDbVersion() {
	//判断版本，依次执行升级
	g.Log().Println("升级")
	fmt.Println("updateDataBase")
}
