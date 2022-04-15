/*
@Time : 2022/4/8 3:37 下午
@Author : fushisanlang
@File : hello.go
@Software: GoLand
*/
package screen

import (
	"farm/service"
	"github.com/gdamore/tcell/v2"
	"github.com/gogf/gf/util/gconv"
)

func CreateScreen(s tcell.Screen) {
	farmClassList := service.GetFarmClass()
	classId := service.UserInfo.FarmClassId
	var className string
	if classId == -1 {
		className = ""
	} else {
		className = farmClassList[classId].FarmClassName
	}

	w, h := s.Size()
	s.Clear()
	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	emitStrCenter(s, w, h/2, style, "道号 : "+service.UserInfo.UserName, 3)
	if service.UserInfo.UserAge == -1 {
		emitStrCenter(s, w, h/2+1, style, "道龄 : ", 3)
	} else {

		emitStrCenter(s, w, h/2+1, style, "道龄 : "+gconv.String(service.UserInfo.UserAge), 3)

	}
	emitStrCenter(s, w, h/2+2, style, "道场 : "+service.UserInfo.FarmName, 3)
	emitStrCenter(s, w, h/2+3, style, "福地 : "+className, 3)

	emitStrCenter(s, w, h/2+4, style, "灵宠 : "+service.UserInfo.PetName, 3)
	createMessageScreen(s)
	s.Show()
}
