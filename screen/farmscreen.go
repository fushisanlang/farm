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

func FarmScreen(s tcell.Screen, page int, id int) (int, int) {
	var pageCount int
	pageCount = 4
	s.Clear()
	printAllBox(s)
	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	fieldList := service.GetFieldInfo(id)

	listLen := len(fieldList)
	pageCount, page = GetPageAndCount(listLen, pageCount, page)

	for i := 0; i < pageCount; i++ {

		emitStr(s, 4, 3+(8*i), style, "编  号 : "+tools.NumToChinsesT(id)+
			tools.NumToChinsesD(i+(page-1)*pageCount+1)+"("+tools.NumToKeyMap(i+1)+")")
		// 0 + ( 3 -1) *
		if fieldList[i+(page-1)*pageCount].IsOpen == 1 {
			if fieldList[i].PlantName != "" {
				emitStr(s, 4, 4+(8*i), style, "灵植名称 : "+fieldList[i].PlantName)
				emitStr(s, 4, 5+(8*i), style, "目前状态 : "+gconv.String(fieldList[i].DuringTime))
				emitStr(s, 4, 6+(8*i), style, "成熟时间 : "+gconv.String(fieldList[i].NeedTime-fieldList[i].DuringTime))
			} else {
				emitStr(s, 4, 4+(8*i), style, "尚未种植")

			}

		} else {
			emitStr(s, 4, 4+(8*i), style, "尚未开启")

		}
	}
	farmScreen(s)
	s.Show()
	return page, pageCount
}

func farmScreen(s tcell.Screen) {

	var a model.ScreenInfoMessage
	a.MessageStr = "    数字 : 查看详细"
	infoMessageScreen(s, []model.ScreenInfoMessage{a, leftPress, rightPress})
	infoMessageScreenColumn2(s, []model.ScreenInfoMessage{nPress, qPress})
}
