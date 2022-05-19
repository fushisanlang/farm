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
	"github.com/gogf/gf/frame/g"
	"os"
)

func createPage(s tcell.Screen, id int) {
	s.Fini()
	tools.CallClear()
	switch id {
	case 1:
		g.Log().Println("创建用户 道号")
		service.GetUserName()
	case 2:
		g.Log().Println("创建用户 道龄")
		service.GetUserAge()
	case 3:
		g.Log().Println("创建用户 道场")
		service.GetFarmName()
	case 4:
		g.Log().Println("创建用户 福地")
		service.GetFarmClassId()
	case 5:
		g.Log().Println("创建用户 灵宠")
		service.GetPetName()
	}
	s, err := tcell.NewScreen()
	tools.CheckErr(err)

	s = service.CreateNewScreen()
	g.Log().Println("返回创建页面")
	CreatePage(s)
}

func CreatePage(s tcell.Screen) {
	g.Log().Println("创建用户")
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
