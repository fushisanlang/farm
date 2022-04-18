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

func PlantScreen(s tcell.Screen, id int, fieldId int) {
	s.Clear()

	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	seedList := service.GetSeedList()
	//emitStr(s, 3, 5, style, "园地编号 : "+tools.NumToChinsesT(id)+tools.NumToChinsesD(fieldId))
	//if fieldInfo.IsOpen == 0 {
	//	OpenStatus,IsMoneyEnough :=service.OpenField(id,fieldId)
	//	if  OpenStatus == true {
	//		emitStr(s, 3, 6, style, "开启成功,剩余灵石：" + gconv.String(IsMoneyEnough))
	//	} else {
	//		emitStr(s, 3, 6, style, "开启失败，灵石不足,尚缺少灵石：" + gconv.String(0-IsMoneyEnough))
	//	}
	//} else {
	//	emitStr(s, 3, 6, style, "园地已开启")
	//}
	for i := 0; i < len(seedList); i++ {
		emitStr(s, 3, 5, style, "种子名称 ： "+seedList[i].PlantName)
		emitStr(s, 30, 5, style, "选取按键 ： "+tools.NumToKeyMap((i+1)%10))
		emitStr(s, 3, 6, style, "种子数量 ： "+gconv.String(seedList[i].CountNum))
		emitStr(s, 30, 6, style, "成熟时间 ： "+gconv.String(seedList[i].TimeNeed)+"昼夜")
		emitStr(s, 3, 7, style, "成品价格 ： "+gconv.String(seedList[i].Price)+"灵石")
		emitStr(s, 30, 7, style, "提供经验 ： "+gconv.String(seedList[i].Ex))

	}
	plantScreen(s)
	s.Show()
}

func plantScreen(s tcell.Screen) {

	infoMessageScreen(s, []model.ScreenInfoMessage{bPress, nPress, qPress})
}
