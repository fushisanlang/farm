/*
@Time : 2022/5/16 17:13
@Author : fushisanlang
@File : createnewscreen
@Software: GoLand
*/
package service

import (
	"farm/tools"
	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
)

func CreateNewScreen() tcell.Screen {
	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	encoding.Register()
	s, err := tcell.NewScreen()
	tools.CheckErr(err)
	err = s.Init()
	tools.CheckErr(err)
	s.SetStyle(defStyle)
	return s
}
