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

func buyStorePage(s tcell.Screen, RatioBuy int) {
	service.RefreshUserInfo()
	screen.BuyStoreScreen(s)
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

				case '1':
					s.Sync()
					//buyStorePage(s)
					buyPage(s, 1, RatioBuy)
				case '2':
					s.Sync()
					buyPage(s, 2, RatioBuy)
				case '3':
					s.Sync()
					buyPage(s, 4, RatioBuy)
				case 'q':
					s.Fini()
					os.Exit(0)
				}

			}
		}
	}

}
