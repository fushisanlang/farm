/*
@Time : 2022/4/8 3:37 下午
@Author : fushisanlang
@File : hello.go
@Software: GoLand
*/
package screen

import (
	"farm/model"
	"farm/service"
	"github.com/gdamore/tcell/v2"
	"github.com/gogf/gf/util/gconv"
)

func FirstStoreScreen(s tcell.Screen) (int, int) {
	///优化显示，根据实际土地状态进行显示
	s.Clear()
	service.ChangeRatioAuto()
	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	RatioSale, RatioBuy := gconv.Int(service.RatioSale), gconv.Int(service.RatioBuy)
	emitStr(s, 3, 5, style, "当前购买价格为标准价格的 : "+gconv.String(RatioBuy)+"%")
	emitStr(s, 3, 6, style, "当前售卖价格为标准价格的 : "+gconv.String(RatioSale)+"%")
	emitStr(s, 3, 7, style, "每次访问商店，价格都会变化。")

	firstStoreScreen(s)
	infoMessageScreenColumn2(s, []model.ScreenInfoMessage{nPress, qPress})
	s.Show()
	return RatioBuy, RatioSale
}
func firstStoreScreen(s tcell.Screen) {
	var a, b model.ScreenInfoMessage
	a.MessageStr = "    1 : 购买"
	b.MessageStr = "    2 : 售卖"
	infoMessageScreen(s, []model.ScreenInfoMessage{a, b})
}
