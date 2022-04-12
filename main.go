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

var version string
var defStyle tcell.Style

func init() {
	version = "0.0.1"
	DbVersion := service.GetVersion()
	g.Log().Println("校验版本")

	if DbVersion != version {
		service.UpdateDbVersion()
	}
}

func main() {

	g.Log().Println("服务启动")
	g.Log().Println("确认用户")
	service.GetUserInfo()
	defStyle = tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	encoding.Register()
	s, err := tcell.NewScreen()
	tools.CheckErr(err)
	err = s.Init()
	tools.CheckErr(err)

	s.SetStyle(defStyle)
	page.VerifySizePage(s)

	page.StartPage(s)
}
