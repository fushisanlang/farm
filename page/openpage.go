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

func openPage(s tcell.Screen, screenId int, fieldId int, page int) {

	screen.OpenScreen(s, screenId, fieldId)
	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventKey:

			switch ev.Key() {
			case tcell.KeyEscape:
				s.Fini()
				os.Exit(0)

			case tcell.KeyRune:
				switch ev.Rune() {
				case 'n':
					s.Sync()
					farmSelectPage(s, 1)
				case 'b':
					s.Sync()
					FieldPage(s, screenId, fieldId, page)
				case 'q':
					s.Fini()
					os.Exit(0)
				}

			}
		}
	}

}
