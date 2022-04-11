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
)

var version string
var defStyle tcell.Style

func init() {
	version = "0.0.1"
	DbVersion := service.Version
	if DbVersion != version {
		service.UpdateDbVersion()
	}

}

func main() {

	defStyle = tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	encoding.Register()
	s, err := tcell.NewScreen()
	tools.CheckErr(err)
	err = s.Init()
	tools.CheckErr(err)

	s.SetStyle(defStyle)

	service.VerifySize(s)

	page.StartPage(s)
}
