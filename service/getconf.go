/*
@Time : 2022/4/12 2:03 下午
@Author : fushisanlang
@File : getconf
@Software: GoLand
*/
package service

import (
	"farm/Dao"
	"github.com/gogf/gf/util/gconv"
)

func GetConf(confKey string) string {
	ConfValue := gconv.String(Dao.GetConf(confKey))
	return ConfValue
}
