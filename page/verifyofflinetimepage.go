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
	"github.com/gdamore/tcell/v2"
	"os"
)

func VerifyOfflineTimePage(s tcell.Screen) {
	service.RefreshUserInfo()
	OfflineTime := service.VerifyOfflineTime()
	if OfflineTime > 600 {

		screen.VerifyOfflineTimescreen(s, OfflineTime)
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
						farmSelectPage(s, 1)

					case 'q':
						s.Fini()
						os.Exit(0)
					}
				}
			}

		}

	}

}
