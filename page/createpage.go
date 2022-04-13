/*
@Time : 2022/4/11 11:24 上午
@Author : fushisanlang
@File : helloworld
@Software: GoLand
*/
package page

import (
	"farm/screen"
	"farm/service"
	"farm/tools"
	"github.com/gdamore/tcell/v2"
	"os"
)

func createPage(s tcell.Screen, id int) {
	s.Fini()
	tools.CallClear()
	switch id {
	case 1:
		service.GetUserName()
	case 2:
		service.GetUserAge()
	case 3:
		service.GetFarmName()
	case 4:
		service.GetFarmClassId()
	case 5:
		service.GetPetName()
	}
	s, err := tcell.NewScreen()
	tools.CheckErr(err)

	s.Init()
	CreatePage(s)
}

func CreatePage(s tcell.Screen) {
	screen.CreateScreen(s)
	if service.VerifyUserInfo() == false {
		for {
			switch ev := s.PollEvent().(type) {
			case *tcell.EventKey:

				switch ev.Key() {
				case tcell.KeyEscape:
					s.Fini()
					os.Exit(0)
				case tcell.KeyRune:
					switch ev.Rune() {
					case 'n':
						s.Sync()
						StartPage(s)
					case '1': //姓名
						createPage(s, 1)
					case '2': //年龄
						createPage(s, 2)
					case '3': //道场
						createPage(s, 3)
					case '4': //道场种类
						createPage(s, 4)
					case '5': //灵宠名称
						createPage(s, 5)
					case 'q':
						s.Fini()
						os.Exit(0)
					}

				}
			}
		}
	} else {
		StartPage(s)
	}
}
