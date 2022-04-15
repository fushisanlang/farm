/*
@Time : 2022/4/12 9:48 上午
@Author : fushisanlang
@File : infoMessage
@Software: GoLand
*/
package screen

import (
	"farm/model"
	"github.com/gdamore/tcell/v2"
)

func createMessageScreen(s tcell.Screen) {
	var onePress, twoPress, threePress, fourPress, fivePress model.ScreenInfoMessage
	onePress.MessageStr = "       1 : 修改道号"
	twoPress.MessageStr = "       2 : 修改道龄"
	threePress.MessageStr = "       3 : 修改道场"
	fourPress.MessageStr = "       4 : 修改福地"
	fivePress.MessageStr = "       5 : 修改灵宠"
	infoMessageList := []model.ScreenInfoMessage{onePress, twoPress, threePress, fourPress, fivePress, qPress}
	infoMessageScreen(s, infoMessageList)

}
