/*
@Time : 2022/4/12 9:48 上午
@Author : fushisanlang
@File : infoMessage
@Software: GoLand
*/
package screen

import "github.com/gdamore/tcell/v2"

func infoMessageScreen(s tcell.Screen) {
	_, h := s.Size()
	emitStr(s, 2, h-4, tcell.StyleDefault, "Press q to quit.")
	emitStr(s, 2, h-3, tcell.StyleDefault, "Press n to next.")
	emitStr(s, 2, h-2, tcell.StyleDefault, "Press b to back.")
	//emitStr(s, 2, h-1, tcell.StyleDefault, "Press  to exit.")

}

func infoMessageOnVerifyScreen(s tcell.Screen) {
	_, h := s.Size()
	emitStr(s, 2, h-2, tcell.StyleDefault, "Press q to quit.")

	//emitStr(s, 2, h-1, tcell.StyleDefault, "Press  to exit.")

}
