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

func FarmScreen(s tcell.Screen, id int) {

	s.Clear()
	printAllBox(s)
	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	fieldList := service.GetFieldInfo(id)
	for i := 0; i < 12; i++ {
		j := i % 4
		k := i / 4
		if i < len(fieldList) {
			emitStr(s, 4+(30*j), 3+(8*k), style, "园地编号 : (F"+gconv.String(fieldList[i].FieldId)+") "+tools.NumToChinsesT(fieldList[i].Id)+tools.NumToChinsesD(fieldList[i].FieldId))

			if fieldList[i].IsOpen == 1 {
				if fieldList[i].PlantName != "" {
					emitStr(s, 4+(30*j), 4+(8*k), style, "灵植名称 : "+fieldList[i].PlantName)
					emitStr(s, 4+(30*j), 5+(8*k), style, "目前状态 : "+gconv.String(fieldList[i].DuringTime))
					emitStr(s, 4+(30*j), 6+(8*k), style, "下一阶段 : "+gconv.String(fieldList[i].NeedTime))
				} else {
					emitStr(s, 4+(30*j), 4+(8*k), style, "尚未种植")

				}
			}
		} else {
			emitStr(s, 4+(30*j), 3+(8*k), style, "园地编号 : (F"+gconv.String(i+1)+") "+tools.NumToChinsesT(id)+tools.NumToChinsesD(i+1))

			emitStr(s, 4+(30*j), 4+(8*k), style, "尚未开启")

		}
	}
	farmScreen(s)
	s.Show()
}

func farmScreen(s tcell.Screen) {

	var a model.ScreenInfoMessage
	a.MessageStr = "Ctrl + _ : 编号快捷键查看详细"
	infoMessageScreen(s, []model.ScreenInfoMessage{a, nPress, qPress})
}
