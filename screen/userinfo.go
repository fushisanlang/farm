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

func UserInfoScreen(s tcell.Screen) {
	w, h := s.Size()
	userInfo := service.GetUserInfo()
	s.Clear()
	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	var sexInfo string
	if userInfo.UserSex == 0 {
		sexInfo = "女"
	} else {
		sexInfo = "男"
	}
	emitStrCenter(s, w, h/2-1, style, userInfo.UserName, 3)
	emitStrCenter(s, w, h/2, style, gconv.String(userInfo.UserAge)+"岁", 3)
	emitStrCenter(s, w, h/2+1, style, sexInfo, 3)
	//emitStrCenter(s, w, h/2+1, style, version, 2)
	infoMessageScreen(s)
	s.Show()
}
