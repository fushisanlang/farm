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

func BuyScreen(s tcell.Screen, buyList []model.BuyList, RatioBuy int) {
	///优化显示，根据实际土地状态进行显示
	s.Clear()

	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	printAllBox3(s)

	emitStr(s, 4, 32, style, "持有晶石"+gconv.String(service.UserMoney)+"块")
	for i := 0; i < len(buyList); i++ {
		emitStr(s, 4+37*(i%2), 3+6*(i/2), style, "按键 ："+gconv.String(i+1))
		emitStr(s, 4+37*(i%2), 4+6*(i/2), style, "名称 ："+buyList[i].Name)
		emitStr(s, 4+37*(i%2), 5+6*(i/2), style, "单价 ："+gconv.String(buyList[i].Price)+"块晶石")
	}
	infoMessageScreenColumn2(s, []model.ScreenInfoMessage{bPress, nPress, qPress})
	buyScreen(s)
	s.Show()
}
func buyScreen(s tcell.Screen) {
	var a model.ScreenInfoMessage
	a.MessageStr = "    数字 : 选择种子"
	infoMessageScreen(s, []model.ScreenInfoMessage{a})
}
