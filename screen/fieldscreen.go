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

func FieldScreen(s tcell.Screen, id int, fieldId int) {
	s.Clear()

	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	fieldInfo := service.GetOneFieldInfo(id, fieldId)
	emitStr(s, 3, 5, style, "园地编号 : "+tools.NumToChinsesT(id)+tools.NumToChinsesD(fieldId))
	if fieldInfo.IsOpen == 0 {
		emitStr(s, 3, 6, style, "园地状态 : 未开启")
	} else if fieldInfo.PlantName == "" {
		emitStr(s, 3, 6, style, "园地状态 : 未种植")
	} else {
		emitStr(s, 3, 6, style, "灵植名称 : "+fieldInfo.PlantName)
		emitStr(s, 3, 7, style, "灵植状态 : "+tools.GetPlantTime(fieldInfo.DuringTime, fieldInfo.NeedTime))
	}
	fieldScreen(s)
	s.Show()
}

func fieldScreen(s tcell.Screen) {
	var i model.ScreenInfoMessage
	i.MessageStr = "       i : 管理"
	infoMessageScreen(s, []model.ScreenInfoMessage{i, bPress, nPress, qPress})
}
