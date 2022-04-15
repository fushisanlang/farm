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

func FarmSelectScreen(s tcell.Screen) {
	s.Clear()
	printAllBox2(s)
	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	fieldList := service.GetFieldInfoAll()
	for i := 0; i < len(fieldList); i++ {
		j := i % 2
		k := i / 2

		emitStr(s, 4+(40*j), 3+(6*k), style, "编  号 : "+tools.NumToChinsesT(fieldList[i].Id)+"  ("+tools.NumToKeyMap(i+1)+") ")
		emitStr(s, 4+(40*j), 4+(6*k), style, "已开启 : "+gconv.String(fieldList[i].IsOpenCount)+"/10")
		emitStr(s, 4+(40*j), 5+(6*k), style, "已种植 : "+gconv.String(fieldList[i].PlantCount)+"/10")

	}
	farmSelectScreen(s)

	s.Show()
}

func farmSelectScreen(s tcell.Screen) {
	var a model.ScreenInfoMessage
	a.MessageStr = "Ctrl + _ : 编号快捷键选择种植园"
	infoMessageScreen(s, []model.ScreenInfoMessage{a, qPress})
}
