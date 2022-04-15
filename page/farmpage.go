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

func FarmPage(s tcell.Screen, screenId int) {

	screen.FarmScreen(s, screenId)
	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventKey:

			switch ev.Key() {
			case tcell.KeyEscape:
				s.Fini()
				os.Exit(0)
			case tcell.KeyCtrlQ:
				s.Sync()
				FieldPage(s, screenId, 1)
			case tcell.KeyCtrlW:
				s.Sync()
				FieldPage(s, screenId, 2)
			case tcell.KeyCtrlE:
				s.Sync()
				FieldPage(s, screenId, 3)
			case tcell.KeyCtrlR:
				s.Sync()
				FieldPage(s, screenId, 4)
			case tcell.KeyCtrlA:
				s.Sync()
				FieldPage(s, screenId, 5)
			case tcell.KeyCtrlS:
				s.Sync()
				FieldPage(s, screenId, 6)
			case tcell.KeyCtrlD:
				s.Sync()
				FieldPage(s, screenId, 7)
			case tcell.KeyCtrlF:
				s.Sync()
				FieldPage(s, screenId, 8)
			case tcell.KeyCtrlZ:
				s.Sync()
				FieldPage(s, screenId, 9)
			case tcell.KeyCtrlX:
				s.Sync()
				FieldPage(s, screenId, 10)
			case tcell.KeyCtrlC:
				s.Sync()
				FieldPage(s, screenId, 11)
			case tcell.KeyCtrlV:
				s.Sync()
				FieldPage(s, screenId, 12)
			case tcell.KeyRune:
				switch ev.Rune() {
				case 'n':
					s.Sync()
					FarmSelectPage(s)

				case 'q':
					s.Fini()
					os.Exit(0)
				}

			}
		}
	}

}
