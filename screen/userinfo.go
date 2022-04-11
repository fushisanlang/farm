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

func UserInfo(s tcell.Screen) {
	w, h := s.Size()
	userInfo := service.GetUserInfo()
	s.Clear()
	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	emitStrCenter(s, w, h/2, style, userInfo, 3)
	//emitStrCenter(s, w, h/2+1, style, version, 2)
	infoMessage(s)
	s.Show()
}
