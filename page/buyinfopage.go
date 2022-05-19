/*
@Time : 2022/4/11 11:24 上午
@Author : fushisanlang
@File : helloworld
@Software: GoLand
*/
package page

import (
	"farm/model"
	"farm/screen"
	"farm/service"
	"farm/tools"
	"github.com/gdamore/tcell/v2"
	"os"
)

func buyInfoPage(s tcell.Screen, ratioBuy int, buyInfo model.BuyList) {
	service.RefreshUserInfo()
	screen.BuyInfoScreen(s, ratioBuy, buyInfo)
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
				case 'i':
					s.Fini()

					tools.CallClear()
					service.BuyGoods(buyInfo, ratioBuy)
					tools.CallClear()
					s = service.CreateNewScreen()

					buyStorePage(s, ratioBuy)
				case 'b':
					s.Sync()
					buyStorePage(s, ratioBuy)

				case 'q':
					s.Fini()
					os.Exit(0)
				}

			}
		}
	}

}
