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
	"github.com/gdamore/tcell/v2"
	"os"
)

//显示背包可卖物品列表，每页5个，左右键翻页
func buyPage(s tcell.Screen, groupId int, RatioBuy int) {
	var buyList []model.BuyList
	if groupId == 1 {
		buyList = service.BuySeedList
	} else if groupId == 2 {
		buyList = service.BuyPlantList

	} else {
		buyList = service.BuyGroceriesList

	}
	screen.BuyScreen(s, buyList, RatioBuy)
	//page, _, _ = screen.SelectGoodsScreen(s, page, goodsList,ratioSale)

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
					buyStorePage(s, RatioBuy)
				case 'q':
					s.Fini()
					os.Exit(0)
				case '1':
					s.Sync()
					buyInfoPage(s, RatioBuy, buyList[0])
				case '2':
					s.Sync()
					buyInfoPage(s, RatioBuy, buyList[1])
				case '3':
					s.Sync()
					buyInfoPage(s, RatioBuy, buyList[2])
				case '4':
					s.Sync()
					buyInfoPage(s, RatioBuy, buyList[3])
				case '5':
					s.Sync()
					buyInfoPage(s, RatioBuy, buyList[4])
				case '6':
					s.Sync()
					buyInfoPage(s, RatioBuy, buyList[5])
				case '7':
					s.Sync()
					buyInfoPage(s, RatioBuy, buyList[6])
				case '8':
					s.Sync()
					buyInfoPage(s, RatioBuy, buyList[7])

				}

			}

		}
	}
}
