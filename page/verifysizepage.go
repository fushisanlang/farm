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

func VerifySizePage(s tcell.Screen) {
	for service.VerifySize(s) == false {
		screen.VerifySizeScreen(s, service.GetWeightMin(), service.GetHeightMin())
		switch ev := s.PollEvent().(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				s.Fini()
				os.Exit(0)
			case tcell.KeyRune:
				switch ev.Rune() {

				case 'q':
					s.Fini()
					os.Exit(0)
				}

			}

		}
	}
}
