/*
@Time : 2022/4/20 9:44 上午
@Author : fushisanlang
@File : harvestscreen
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

func HarvestScreen(s tcell.Screen, id int, fieldId int) {
	s.Clear()

	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)

	HarvestInfo := service.Harvest(id, fieldId)
	emitStr(s, 3, 5, style, "园地编号 : "+tools.NumToChinsesT(id)+tools.NumToChinsesD(fieldId))

	emitStr(s, 3, 6, style, "收获"+HarvestInfo.PlantName+"共"+
		gconv.String(HarvestInfo.Yield)+"份,已经放入背包。")

	infoMessageScreen(s, []model.ScreenInfoMessage{bPress, nPress, qPress})
	s.Show()
}
