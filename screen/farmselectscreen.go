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

func FarmSelectScreen(s tcell.Screen, page int) int {
	var pageCount int
	pageCount = 5
	basePageCount := pageCount
	s.Clear()
	printAllBox2(s)
	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	fieldList := service.GetFieldInfoAll()

	listLen := len(fieldList)
	pageCount, page = GetPageAndCount(listLen, pageCount, page)

	for i := 0; i < pageCount; i++ {
		emitStr(s, 4, 3+(6*i), style, "编  号 : "+tools.NumToChinsesT(i+(page-1)*basePageCount+1)+
			"  ("+tools.NumToKeyMap(i+1)+") ")
		emitStr(s, 4, 4+(6*i), style, "已开启 : "+
			gconv.String(fieldList[i+(page-1)*basePageCount].IsOpenCount)+"/10")
		emitStr(s, 4, 5+(6*i), style, "已种植 : "+
			gconv.String(fieldList[i+(page-1)*basePageCount].PlantCount)+"/10")
	}
	farmSelectScreen(s)
	infoMessageScreenColumn2(s, []model.ScreenInfoMessage{vPress, qPress})

	s.Show()
	return page
}

func farmSelectScreen(s tcell.Screen) {
	var a model.ScreenInfoMessage
	a.MessageStr = "    数字 : 选择种植园"
	infoMessageScreen(s, []model.ScreenInfoMessage{a, leftPress, rightPress})
}
