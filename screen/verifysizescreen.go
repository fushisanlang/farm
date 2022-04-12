/*
@Time : 2022/4/8 3:37 下午
@Author : fushisanlang
@File : hello.go
@Software: GoLand
*/
package screen

import (
	"github.com/gdamore/tcell/v2"
	"github.com/gogf/gf/util/gconv"
)

func VerifySizeScreen(s tcell.Screen, WeightMin int, HeightMin int) {
	w, h := s.Size()
	s.Clear()
	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	emitStrCenter(s, w, h/2-3, style, "预期窗口大小为", 3)
	emitStrCenter(s, w, h/2-2, style, gconv.String(WeightMin)+"*"+gconv.String(HeightMin), 2)
	emitStrCenter(s, w, h/2-1, style, "目前窗口大小为", 3)
	emitStrCenter(s, w, h/2-0, style, gconv.String(w)+"*"+gconv.String(h), 2)
	emitStrCenter(s, w, h/2+1, style, "请调整窗口大小", 3)
	infoMessageOnVerifyScreen(s)
	s.Show()
}
