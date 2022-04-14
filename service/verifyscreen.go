/*
@Time : 2022/4/11 10:16 上午
@Author : fushisanlang
@File : verifyscreen
@Software: GoLand
*/
package service

import (
	"github.com/gdamore/tcell/v2"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func VerifySize(s tcell.Screen) bool {
	//修改，辅助调整 否则打不开.
	w, h := s.Size()
	g.Log().Println("调整窗口大小到" + gconv.String(w) + "*" + gconv.String(h))

	if w < WeightMin || h < HeightMin {
		return false

	} else {
		return true
	}
}
