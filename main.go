/*
@Time : 2022/4/8 2:55 下午
@Author : fushisanlang
@File : main.go
@Software: GoLand
*/

package main

import (
	"farm/page"
	"farm/service"
	"farm/tools"
	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Log().Println("服务启动")
	g.Log().Println("校验数据库版本")

	service.VerifyVersion()

	g.Log().Println("创建窗口")
	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	encoding.Register()
	s, err := tcell.NewScreen()
	tools.CheckErr(err)
	err = s.Init()
	tools.CheckErr(err)
	s.SetStyle(defStyle)

	g.Log().Println("校验用户基础数据")

	isUserExist := service.VerifyUserInfo()
	if isUserExist == false {
		page.CreatePage(s)

	}

	page.VerifySizePage(s)
	page.StartPage(s)
}
