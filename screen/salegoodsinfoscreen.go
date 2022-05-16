/*
@Time : 2022/4/8 3:37 下午
@Author : fushisanlang
@File : hello.go
@Software: GoLand
*/
package screen

import (
	"farm/model"
	"github.com/gdamore/tcell/v2"
	"github.com/gogf/gf/util/gconv"
)

func SaleGoodsInfoScreen(s tcell.Screen, keyid int, ratioSale int, goodsInfo model.GoodsList) {
	s.Clear()
	// bagId,ratioSale
	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)

	//service.Sale(id, fieldId)
	emitStr(s, 3, 5, style, "货物名称 : "+goodsInfo.Name)

	emitStr(s, 3, 6, style, "货物总数 : "+gconv.String(goodsInfo.Countnum))
	emitStr(s, 3, 7, style, "售卖单价 : "+gconv.String(goodsInfo.Price*ratioSale/100))

	saleGoodsInfoScreen(s)
	s.Show()
}

func saleGoodsInfoScreen(s tcell.Screen) {
	var i model.ScreenInfoMessage
	i.MessageStr = "       i : 售卖"
	infoMessageScreen(s, []model.ScreenInfoMessage{i, bPress, qPress, nPress})
}
