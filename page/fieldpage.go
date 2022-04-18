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

func FieldPage(s tcell.Screen, screenId int, fieldId int, page int) {

	fieldInfo := screen.FieldScreen(s, screenId, fieldId)
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

				case 'i': //开启
					s.Sync()
					if fieldInfo.IsOpen == 0 {
						openPage(s, screenId, fieldId, page)
					}
				case 'u': //种植
					s.Sync()
					if fieldInfo.PlantName == "" {
						plantPage(s, screenId, fieldId, page)
					}
				case 'o':
					s.Sync()
					WaitPage(s, "o")
				case 'p':
					s.Sync()
					WaitPage(s, "p")
				case 'b':
					s.Sync()
					FarmPage(s, page, screenId)
				case 'q':
					s.Fini()
					os.Exit(0)
				}

			}
		}
	}

}