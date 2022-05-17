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
	"farm/tools"
	"github.com/gdamore/tcell/v2"
	"github.com/gogf/gf/util/gconv"
)

func OpenScreen(s tcell.Screen, id int, fieldId int) {
	s.Clear()

	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	fieldInfo := service.GetOneFieldInfo(id, fieldId)
	emitStr(s, 3, 5, style, "园地编号 : "+tools.NumToChinsesT(id)+tools.NumToChinsesD(fieldId))
	if fieldInfo.IsOpen == 0 {
		OpenStatus, IsMoneyEnough := service.OpenField(id, fieldId)
		if OpenStatus == true {
			service.ChangeMemUserInfoMoney()

			emitStr(s, 3, 6, style, "开启成功,剩余灵石："+gconv.String(IsMoneyEnough))
		} else {
			emitStr(s, 3, 6, style, "开启失败，灵石不足,尚缺少灵石："+gconv.String(0-IsMoneyEnough))
		}
	} else {
		emitStr(s, 3, 6, style, "园地已开启")
	}

	openScreen(s)
	s.Show()
}

func openScreen(s tcell.Screen) {

	infoMessageScreen(s, []model.ScreenInfoMessage{bPress, nPress, qPress})
}
