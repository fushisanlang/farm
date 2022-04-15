/*
@Time : 2022/4/11 11:24 上午
@Author : fushisanlang
@File : helloworld
@Software: GoLand
*/
package page

import (
	"farm/screen"
	"github.com/gdamore/tcell/v2"
	"os"
)

func FarmSelectPage(s tcell.Screen) {
	screen.FarmSelectScreen(s)
	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventKey:

			switch ev.Key() {
			case tcell.KeyEscape:
				s.Fini()
				os.Exit(0)
			case tcell.KeyCtrlQ:
				s.Sync()
				FarmPage(s, 1)
			case tcell.KeyCtrlW:
				s.Sync()
				FarmPage(s, 2)
			case tcell.KeyCtrlE:
				s.Sync()
				FarmPage(s, 3)
			case tcell.KeyCtrlR:
				s.Sync()
				FarmPage(s, 4)
			case tcell.KeyCtrlA:
				s.Sync()
				FarmPage(s, 5)
			case tcell.KeyCtrlS:
				s.Sync()
				FarmPage(s, 6)
			case tcell.KeyCtrlD:
				s.Sync()
				FarmPage(s, 7)
			case tcell.KeyCtrlF:
				s.Sync()
				FarmPage(s, 8)
			case tcell.KeyCtrlZ:
				s.Sync()
				FarmPage(s, 9)
			case tcell.KeyCtrlX:
				s.Sync()
				FarmPage(s, 10)
			//case tcell.KeyCtrlC:
			//	s.Sync()
			//	FarmPage(s,11)
			//case tcell.KeyCtrlV:
			//	s.Sync()
			//	FarmPage(s,12)

			case tcell.KeyRune:
				switch ev.Rune() {

				case 'q':
					s.Fini()
					os.Exit(0)
				}

			}
		}
	}

}
