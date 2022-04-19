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
)

func EradicateScreen(s tcell.Screen, id int, fieldId int) {
	s.Clear()

	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)

	service.Eradicate(id, fieldId)
	emitStr(s, 3, 5, style, "园地编号 : "+tools.NumToChinsesT(id)+tools.NumToChinsesD(fieldId))

	emitStr(s, 3, 6, style, "铲除成功")

	infoMessageScreen(s, []model.ScreenInfoMessage{bPress, nPress, qPress})
	s.Show()
}
