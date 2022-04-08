/*
@Time : 2022/4/8 3:44 下午
@Author : fushisanlang
@File : tools
@Software: GoLand
*/
package screen

import (
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}

func infoMessage(s tcell.Screen) {
	_, h := s.Size()
	emitStr(s, 2, h-3, tcell.StyleDefault, "Press ESC to exit.")
	emitStr(s, 2, h-2, tcell.StyleDefault, "Press N or -> to next.")
	//emitStr(s, 2, h-1, tcell.StyleDefault, "Press  to exit.")

}
