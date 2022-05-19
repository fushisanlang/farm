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

func FarmPage(s tcell.Screen, page int, screenId int) {
	service.RefreshUserInfo()
	page, pageCount := screen.FarmScreen(s, page, screenId)
	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventKey:

			switch ev.Key() {
			case tcell.KeyEscape:
				s.Fini()
				os.Exit(0)
			case tcell.KeyRight:
				FarmPage(s, page+1, screenId)
			case tcell.KeyLeft:
				FarmPage(s, page-1, screenId)

			case tcell.KeyRune:
				switch ev.Rune() {
				case 'n':
					s.Sync()
					farmSelectPage(s, 1)
				case '1':
					FieldPage(s, screenId, 1+(page-1)*pageCount, page)
				case '2':
					FieldPage(s, screenId, 2+(page-1)*pageCount, page)
				case '3':
					FieldPage(s, screenId, 3+(page-1)*pageCount, page)
				case '4':
					FieldPage(s, screenId, 4+(page-1)*pageCount, page)

				case 'q':
					s.Fini()
					os.Exit(0)
				}

			}
		}
	}

}
