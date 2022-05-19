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

func firstStorePage(s tcell.Screen) {
	service.RefreshUserInfo()
	RatioBuy, RatioSale := screen.FirstStoreScreen(s)
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
					bagNullId := service.GetNullBagId(0, 0)
					if bagNullId != 0 {
						buyStorePage(s, RatioBuy)
					} else {
						errorPage(s, "背包满。")
					}

				case '2':
					s.Sync()
					saleStorePage(s, 1, RatioSale)

				case 'q':
					s.Fini()
					os.Exit(0)
				}

			}
		}
	}

}
