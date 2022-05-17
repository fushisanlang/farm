/*
@Time : 2022/4/8 3:37 下午
@Author : fushisanlang
@File : hello.go
@Software: GoLand
*/
package screen

import (
	"farm/model"
	"github.com/gdamore/tcell/v2"
)

func ErrorScreen(s tcell.Screen, errorString string) {
	w, h := s.Size()
	s.Clear()
	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	emitStrCenter(s, w, h/2, style, "errorString", 3)
	infoMessageList := []model.ScreenInfoMessage{nPress, qPress}
	infoMessageScreen(s, infoMessageList)
	s.Show()
}
