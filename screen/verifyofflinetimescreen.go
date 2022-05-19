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
	"github.com/gdamore/tcell/v2"
	"github.com/gogf/gf/util/gconv"
)

func VerifyOfflineTimescreen(s tcell.Screen, offlineTime int) {
	w, h := s.Size()
	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)

		//离线超过60。83小时，游戏中度过10*60*365天，不再计算收益
	if offlineTime > 216000 {
		offlineTime = 216000
		emitStrCenter(s, w, h/2-3, style, "道友已经神游超过10甲子，道场已翻天覆地。", 3)
		emitStrCenter(s, w, h/2-2, style, "道友快快去整理一番。", 3)

	} else {
		//215999
		offlineSixtyYear := offlineTime / 21600          //9
		offlineYear := (offlineTime % 21600) / 360       //59
		offlineMonth := (offlineTime % 21600 % 360) / 30 //11
		offlineDay := offlineTime % 21600 % 360 % 30     //29
		a := ""
		if offlineSixtyYear != 0 {
			a = a + gconv.String(offlineSixtyYear) + "甲子"
		}
		if offlineYear != 0 {
			a = a + gconv.String(offlineYear) + "年"
		}
		if offlineMonth != 0 {
			a = a + gconv.String(offlineMonth) + "月"
		}
		if offlineDay != 0 {
			a = a + gconv.String(offlineDay) + "昼夜"
		}

		emitStrCenter(s, w, h/2-3, style, "道友已经神游"+a+"，道场已翻天覆地。", 3)
		emitStrCenter(s, w, h/2-2, style, "快快去整理一番吧。", 3)

	}
	addTime := offlineTime / 2
	service.DuringTimeAdd(addTime)
	//昼夜

	infoMessageScreen(s, []model.ScreenInfoMessage{nPress, qPress})
	s.Show()
}
