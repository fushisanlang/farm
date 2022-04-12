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

func GetVersion() string {
	Version := gconv.String(Dao.GetConf("version"))
	return Version
}
func GetWeightMin() int {
	WeightMin := gconv.Int(Dao.GetConf("version"))
	return WeightMin
}
func GetHeightMin() int {
	HeightMin := gconv.Int(Dao.GetConf("version"))
	return HeightMin
}
