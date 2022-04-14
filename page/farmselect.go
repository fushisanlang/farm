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
			case tcell.KeyRune:
				switch ev.Rune() {
				case 'n':
					s.Sync()
					FarmPage(s)

				case 'q':
					s.Fini()
					os.Exit(0)
				}

			}
		}
	}

}