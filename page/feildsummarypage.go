/*
@Time : 2022/4/11 11:24 上午
@Author : fushisanlang
@File : helloworld
@Software: GoLand
*/
package page

import (
	"farm/screen"
	"farm/service"
	"github.com/gdamore/tcell/v2"
	"os"
)

func FeildSummaryPage(s tcell.Screen) {
	service.RefreshUserInfo()
	service.RefreserCode = 0
	go screen.FeildSummaryScreen(s)

	for {

		switch ev := s.PollEvent().(type) {
		case *tcell.EventKey:

			switch ev.Key() {
			case tcell.KeyEscape:
				service.RefreserCode = 1
				s.Fini()
				os.Exit(0)
			case tcell.KeyRune:
				switch ev.Rune() {
				case 'n':
					service.RefreserCode = 1
					s.Sync()
					farmSelectPage(s, 1)

				case 'q':
					service.RefreserCode = 1
					s.Fini()
					os.Exit(0)
				}

			}
		}
	}

}
