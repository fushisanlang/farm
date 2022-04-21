/*
@Time : 2022/4/11 11:24 上午
@Author : fushisanlang
@File : helloworld
@Software: GoLand
*/
package page

import (
	"farm/screen"
	"github.com/gdamore/tcell/v2"
	"os"
)

func farmSelectPage(s tcell.Screen, page int) {
	page = screen.FarmSelectScreen(s, page)
	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventKey:

			switch ev.Key() {
			case tcell.KeyEscape:
				s.Fini()
				os.Exit(0)
			//	FarmPage(s,12)
			case tcell.KeyRight:
				s.Sync()
				farmSelectPage(s, page+1)
			case tcell.KeyLeft:
				s.Sync()
				farmSelectPage(s, page-1)
			case tcell.KeyRune:
				switch ev.Rune() {
				case '1':
					s.Sync()
					FarmPage(s, 1, 1+(page-1)*5)
				case '2':
					s.Sync()
					FarmPage(s, 1, 2+(page-1)*5)
				case '3':
					s.Sync()
					FarmPage(s, 1, 3+(page-1)*5)
				case '4':
					s.Sync()
					FarmPage(s, 1, 4+(page-1)*5)
				case '5':
					s.Sync()
					FarmPage(s, 1, 5+(page-1)*5)
				case 'v':
					s.Sync()
					firstStorePage(s)
				case 'q':
					s.Fini()
					os.Exit(0)
				}

			}
		}
	}

}
