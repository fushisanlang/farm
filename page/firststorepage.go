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
	"github.com/gogf/gf/util/gconv"
	"os"
)

func firstStorePage(s tcell.Screen) {

	RatioBuy, RatioSale := screen.FirstStoreScreen(s)
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

				case '1':
					s.Sync()
					//buyStorePage(s)
					WaitPage(s, "buyStorePage"+gconv.String(RatioBuy))
				case '2':
					s.Sync()
					saleStorePage(s, 1, RatioSale)

				case 'q':
					s.Fini()
					os.Exit(0)
				}

			}
		}
	}

}
