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

//显示背包可卖物品列表，每页5个，左右键翻页
func saleStorePage(s tcell.Screen, page int, ratioSale int) {
	service.RefreshUserInfo()
	goodsList := service.GetGoodsList()
	page, pageCount, basePageCount := screen.SelectGoodsScreen(s, page, goodsList, ratioSale)
	//page, _, _ = screen.SelectGoodsScreen(s, page, goodsList,ratioSale)

	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventKey:

			switch ev.Key() {
			case tcell.KeyEscape:
				s.Fini()
				os.Exit(0)
			case tcell.KeyRight:
				s.Sync()
				saleStorePage(s, page+1, ratioSale)
			case tcell.KeyLeft:
				s.Sync()
				saleStorePage(s, page-1, ratioSale)
			case tcell.KeyRune:

				switch ev.Rune() {
				case 'n':
					s.Sync()
					farmSelectPage(s, 1)

				case 'q':
					s.Fini()
					os.Exit(0)
				case '1':
					s.Sync()
					if pageCount > 0 {
						//salePage(s,id,"本页信息") (页面信息，前页信息）
						//根据按键输入，从list拿到相应goods的详细信息
						saleGoodsInfoPage(s, goodsList[1+(page-1)*basePageCount-1].KeyId, ratioSale, goodsList[1+(page-1)*basePageCount-1])
					}
				case '2':
					s.Sync()
					if pageCount > 1 {
						//2            2 + (2-1)*5 -1 =2+5-1=6
						//saleGoodsInfoPage
						saleGoodsInfoPage(s, goodsList[2+(page-1)*basePageCount-1].KeyId, ratioSale, goodsList[2+(page-1)*basePageCount-1])
					}
				case '3':
					s.Sync()
					if pageCount > 1 {
						//2            2 + (2-1)*5 -1 =2+5-1=6
						saleGoodsInfoPage(s, goodsList[3+(page-1)*basePageCount-1].KeyId, ratioSale, goodsList[3+(page-1)*basePageCount-1])
					}
				case '4':
					s.Sync()
					if pageCount > 1 {
						//2            2 + (2-1)*5 -1 =2+5-1=6
						saleGoodsInfoPage(s, goodsList[4+(page-1)*basePageCount-1].KeyId, ratioSale, goodsList[4+(page-1)*basePageCount-1])
					}
				case '5':
					s.Sync()
					if pageCount > 1 {
						//2            2 + (2-1)*5 -1 =2+5-1=6
						saleGoodsInfoPage(s, goodsList[5+(page-1)*basePageCount-1].KeyId, ratioSale, goodsList[5+(page-1)*basePageCount-1])
					}
				}

			}

		}
	}
}
