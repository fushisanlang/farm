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
	emitStr(s, 3, 5, style, "1 zhongzi : ")
	emitStr(s, 3, 6, style, "2 zuowu : ")
	emitStr(s, 3, 7, style, "3 zawu ， ")

	firstStoreScreen(s)
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
