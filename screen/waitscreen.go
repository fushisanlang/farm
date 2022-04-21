/*
@Time : 2022/4/8 3:37 下午
@Author : fushisanlang
@File : hello.go
@Software: GoLand
*/
package screen

import (
	"github.com/gdamore/tcell/v2"
)

func WaitScreen(s tcell.Screen, Waitstring string) {
	s.Clear()

	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	emitStr(s, 3, 5, style, Waitstring)

	emitStr(s, 3, 4, style, Waitstring)
	s.Show()
}
