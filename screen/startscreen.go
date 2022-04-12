/*
@Time : 2022/4/8 3:37 下午
@Author : fushisanlang
@File : hello.go
@Software: GoLand
*/
package screen

import (
	"farm/service"
	"github.com/gdamore/tcell/v2"
)

func StartScreen(s tcell.Screen) {
	w, h := s.Size()
	version := service.GetVersion()
	s.Clear()
	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	emitStrCenter(s, w, h/2, style, "药宗模拟器", 3)
	emitStrCenter(s, w, h/2+1, style, version, 2)
	infoMessageScreen(s)
	s.Show()
}
