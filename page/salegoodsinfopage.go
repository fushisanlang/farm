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

func saleGoodsInfoPage(s tcell.Screen, KeyId int, ratioSale int, goodsInfo model.GoodsList) {
	service.RefreshUserInfo()
	screen.SaleGoodsInfoScreen(s, KeyId, ratioSale, goodsInfo)
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
					service.SaleGoods(goodsInfo, ratioSale)
					tools.CallClear()
					s = service.CreateNewScreen()

					saleStorePage(s, 1, ratioSale)
				case 'b':
					s.Sync()
					saleStorePage(s, 1, ratioSale)

				case 'q':
					s.Fini()
					os.Exit(0)
				}

			}
		}
	}

}
