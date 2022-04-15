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

func FieldPage(s tcell.Screen, screenId int, fieldId int) {

	screen.FieldScreen(s, screenId, fieldId)
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
					FarmSelectPage(s)
				case 'i':
					s.Sync()
					WaitPage(s, "i")
				case 'b':
					s.Sync()
					FarmPage(s, screenId)
				case 'q':
					s.Fini()
					os.Exit(0)
				}

			}
		}
	}

}
