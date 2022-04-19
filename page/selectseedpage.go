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

func selectSeedPage(s tcell.Screen, screenId int, fieldId int, page int) {
	seedList := service.GetSeedList()
	page, pageCount := screen.SelectSeedScreen(s, page, seedList)
	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventKey:

			switch ev.Key() {
			case tcell.KeyEscape:
				s.Fini()
				os.Exit(0)
			case tcell.KeyRight:
				s.Sync()
				selectSeedPage(s, screenId, fieldId, page+1)
			case tcell.KeyLeft:
				s.Sync()
				selectSeedPage(s, screenId, fieldId, page-1)
			case tcell.KeyRune:
				switch ev.Rune() {
				case 'n':
					s.Sync()
					farmSelectPage(s, 1)
				case 'b':
					s.Sync()
					FieldPage(s, screenId, fieldId, page)
				case 'q':
					s.Fini()
					os.Exit(0)
				case '1':
					s.Sync()
					if pageCount > 0 {
						service.PlantSeed(screenId, fieldId, seedList[1+(page-1)-1].Id, seedList[1+(page-1)-1].BagId)
						farmSelectPage(s, 1)
					}
				case '2':
					s.Sync()
					if pageCount > 1 {
						service.PlantSeed(screenId, fieldId, seedList[2+(page-1)-1].Id, seedList[1+(page-1)-1].BagId)
						farmSelectPage(s, 1)
					}
				case '3':
					s.Sync()
					if pageCount > 2 {
						service.PlantSeed(screenId, fieldId, seedList[3+(page-1)-1].Id, seedList[1+(page-1)-1].BagId)
						farmSelectPage(s, 1)
					}
				case '4':
					s.Sync()
					if pageCount > 3 {
						service.PlantSeed(screenId, fieldId, seedList[4+(page-1)-1].Id, seedList[1+(page-1)-1].BagId)
						farmSelectPage(s, 1)
					}
				case '5':
					s.Sync()
					if pageCount > 4 {
						service.PlantSeed(screenId, fieldId, seedList[5+(page-1)-1].Id, seedList[1+(page-1)-1].BagId)
						farmSelectPage(s, 1)
					}
				}

			}
		}
	}

}
