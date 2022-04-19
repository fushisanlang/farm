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

func VerifyEradicateScreen(s tcell.Screen, id int, fieldId int) {
	s.Clear()

	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	fieldInfo := service.GetOneFieldInfo(id, fieldId)
	emitStr(s, 3, 5, style, "园地编号 : "+tools.NumToChinsesT(id)+tools.NumToChinsesD(fieldId))

	emitStr(s, 3, 6, style, "灵植名称 : "+fieldInfo.PlantName)
	emitStr(s, 3, 7, style, "灵植状态 : "+tools.GetPlantTime(fieldInfo.DuringTime, fieldInfo.NeedTime))

	emitStr(s, 3, 8, style, "是否铲除？ [Ctrl+y:铲除/Ctrl+n:取消]")

	infoMessageScreenColumn2(s, []model.ScreenInfoMessage{bPress, nPress, qPress})
	s.Show()
	return
}
