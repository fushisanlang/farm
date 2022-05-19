/*
@Time : 2022/4/8 2:55 下午
@Author : fushisanlang
@File : main.go
@Software: GoLand
*/

package main

import (
	_ "farm/cron"
	"farm/page"
	"farm/service"
	"farm/tools"
	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
	"github.com/gogf/gf/frame/g"
)

func init() {
	g.Log().Println("服务启动")
	g.Log().Println("校验数据库版本")

	service.VerifyVersion()
}

func main() {

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

	isUserExist := service.VerifyUserInfoWithErase()
	if isUserExist == false {
		service.VerifyUserInfo()
		g.Log().Println("用户数据不全，重新创建用户信息")
		page.CreatePage(s)

	}
	g.Log().Println("校验窗口大小")
	page.VerifySizePage(s)
	g.Log().Println("校验离线时间")
	page.VerifyOfflineTimePage(s)
	g.Log().Println("访问起始页")
	page.StartPage(s)
}
