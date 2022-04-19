/*
@Time : 2022/4/8 3:37 下午
@Author : fushisanlang
@File : hello.go
@Software: GoLand
*/
package screen

import (
	"farm/model"
	"farm/tools"
	"github.com/gdamore/tcell/v2"
	"github.com/gogf/gf/util/gconv"
)

func SelectSeedScreen(s tcell.Screen, page int, seedList []model.SeedList) (int, int) {
	var pageCount int
	pageCount = 5
	basePageCount := pageCount
	s.Clear()

	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	listLen := len(seedList)
	pageCount, page = GetPageAndCount(listLen, pageCount, page)
	//2,1
	for i := 0; i < pageCount; i++ {
		//
		emitStr(s, 3, 5+(8*i), style, "种子名称 ： "+seedList[i+(page-1)*basePageCount].PlantName) // i + (2-1) * 3
		emitStr(s, 30, 5+(8*i), style, "选取按键 ： "+tools.NumToKeyMap((i+1)%10))
		emitStr(s, 3, 6+(8*i), style, "种子数量 ： "+gconv.String(seedList[i+(page-1)*basePageCount].CountNum))
		emitStr(s, 30, 6+(8*i), style, "成熟时间 ： "+gconv.String(seedList[i+(page-1)*basePageCount].TimeNeed)+"昼夜")
		emitStr(s, 3, 7+(8*i), style, "成品价格 ： "+gconv.String(seedList[i+(page-1)*basePageCount].Price)+"灵石")
		emitStr(s, 30, 7+(8*i), style, "提供经验 ： "+gconv.String(seedList[i+(page-1)*basePageCount].Ex))

	}
	selectSeedScreen(s)
	infoMessageScreenColumn2(s, []model.ScreenInfoMessage{bPress, nPress, qPress})
	s.Show()
	return page, pageCount
}

func selectSeedScreen(s tcell.Screen) {
	var a model.ScreenInfoMessage
	a.MessageStr = "    数字 : 选择种子"
	infoMessageScreen(s, []model.ScreenInfoMessage{a, leftPress, rightPress})
}
