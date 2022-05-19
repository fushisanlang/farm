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
)

func BuyStoreScreen(s tcell.Screen) {
	///优化显示，根据实际土地状态进行显示
	s.Clear()
	service.ChangeRatioAuto()
	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	printBox4(s, 3, 2)
	printBox4(s, 43, 2)

	printBox4(s, 83, 2)

	emitStr(s, 4, 3, style, "按键信息：1 ")
	emitStr(s, 4, 4, style, "商品种类：种子 ")
	emitStr(s, 44, 3, style, "按键信息：2 ")
	emitStr(s, 44, 4, style, "商品种类：作物 ")
	emitStr(s, 84, 3, style, "按键信息：3 ")
	emitStr(s, 84, 4, style, "商品种类：杂物 ")

	buyStoreScreen(s)
	infoMessageScreenColumn2(s, []model.ScreenInfoMessage{nPress, qPress})
	s.Show()

}
func buyStoreScreen(s tcell.Screen) {
	var a, b, c model.ScreenInfoMessage
	a.MessageStr = "    1 : 购买种子"
	b.MessageStr = "    2 : 购买作物"
	c.MessageStr = "    3 : 购买杂物"

	infoMessageScreen(s, []model.ScreenInfoMessage{a, b, c})
}
