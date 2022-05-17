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

func BuyInfoScreen(s tcell.Screen, ratioBuy int, buyInfo model.BuyList) {
	s.Clear()
	// bagId,ratioSale
	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)

	//service.Sale(id, fieldId)
	emitStr(s, 3, 5, style, "名称 : "+buyInfo.Name)

	emitStr(s, 3, 6, style, "单价 : "+gconv.String(buyInfo.Price*ratioBuy/100))

	buyInfoScreen(s)
	s.Show()
}

func buyInfoScreen(s tcell.Screen) {
	var i model.ScreenInfoMessage
	i.MessageStr = "       i : 购买"
	infoMessageScreen(s, []model.ScreenInfoMessage{i, bPress, qPress, nPress})
}
