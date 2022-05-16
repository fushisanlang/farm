/*
@Time : 2022/4/8 3:37 下午
@Author : fushisanlang
@File : hello.go
@Software: GoLand
*/
package screen

import (
	"farm/model"
	"farm/tools"
	"github.com/gdamore/tcell/v2"
	"github.com/gogf/gf/util/gconv"
)

func SelectGoodsScreen(s tcell.Screen, page int, goodsList []model.GoodsList, ratioSale int) (int, int, int) {
	var pageCount int
	pageCount = 5
	basePageCount := pageCount
	s.Clear()

	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	listLen := len(goodsList)
	pageCount, page = GetPageAndCount(listLen, pageCount, page)
	//2,1
	if listLen == 0 {
		emitStr(s, 3, 5, style, "无可售卖物品") // i + (2-1) * 3

	} else {
		for i := 0; i < pageCount; i++ {
			//
			emitStr(s, 3, 5+(8*i), style, "物品名称 ： "+goodsList[i+(page-1)*basePageCount].Name) // i + (2-1) * 3
			emitStr(s, 30, 5+(8*i), style, "选取按键 ： "+tools.NumToKeyMap((i+1)%10))
			emitStr(s, 3, 6+(8*i), style, "物品数量 ： "+gconv.String(goodsList[i+(page-1)*basePageCount].Countnum))
			emitStr(s, 30, 6+(8*i), style, "物品价格 ： "+gconv.String(goodsList[i+(page-1)*basePageCount].Price*ratioSale/100))

		}
	}
	selectGoodsScreen(s)
	infoMessageScreenColumn2(s, []model.ScreenInfoMessage{nPress, qPress})
	s.Show()
	return page, pageCount, basePageCount
}

func selectGoodsScreen(s tcell.Screen) {
	var a model.ScreenInfoMessage
	a.MessageStr = "    数字 : 售卖"
	infoMessageScreen(s, []model.ScreenInfoMessage{a, leftPress, rightPress})
}
