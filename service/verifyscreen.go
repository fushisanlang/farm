/*
@Time : 2022/4/11 10:16 上午
@Author : fushisanlang
@File : verifyscreen
@Software: GoLand
*/
package service

import (
	"github.com/gdamore/tcell/v2"
)

func VerifySize(s tcell.Screen) bool {
	//修改，辅助调整 否则打不开.
	WeightMin := GetWeightMin()
	HeightMin := GetHeightMin()
	w, h := s.Size()

	if w < WeightMin || h < HeightMin {
		return false

	} else {
		return true
	}
}
