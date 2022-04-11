/*
@Time : 2022/4/11 2:12 下午
@Author : fushisanlang
@File : infoPage
@Software: GoLand
*/
package page

import (
	"farm/screen"
	"github.com/gdamore/tcell/v2"
	"os"
)

func infoPage(s tcell.Screen) {
	screen.UserInfo(s)
	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventKey:

			switch ev.Key() {
			case tcell.KeyEscape:
				s.Fini()
				os.Exit(0)
			case tcell.KeyRune:
				switch ev.Rune() {
				case 'b':
					s.Sync()
					StartPage(s)
				case 'q':
					s.Fini()
					os.Exit(0)
				}

			}
		}
	}

}
