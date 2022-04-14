/*
@Time : 2022/4/14 2:38 下午
@Author : fushisanlang
@File : init
@Software: GoLand
*/
package screen

import (
	"farm/model"
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

var bPress, iPress, nPress, qPress model.ScreenInfoMessage

func init() {
	bPress.MessageStr = "按 b 打开背包"
	bPress.MessageStatus = 3
	iPress.MessageStr = "按 i 打开信息页"
	iPress.MessageStatus = 3
	nPress.MessageStr = "按 n 进入主页面"
	nPress.MessageStatus = 3
	qPress.MessageStr = "按 q 退出程序"
	qPress.MessageStatus = 3

}
func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}

func emitStrCenter(s tcell.Screen, w int, h int, style tcell.Style, printString string, status int) {
	//status golang中string底层是通过byte数组实现的。中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码正好是utf-8。
	//居中计算长度是需要*2/3，取半数计算缩进要/2，结果就是中文需要/3，英语需要/2。
	a := len(printString)
	emitStr(s, w/2-a/status, h, style, printString)

}
func printAllBox(s tcell.Screen) {
	printBox(s, 3, 2)
	printBox(s, 3, 10)
	printBox(s, 3, 18)
	printBox(s, 33, 2)
	printBox(s, 33, 10)
	printBox(s, 33, 18)
	printBox(s, 63, 2)
	printBox(s, 63, 10)
	printBox(s, 63, 18)
	printBox(s, 93, 2)
	printBox(s, 93, 10)
	printBox(s, 93, 18)
}
func printAllBox2(s tcell.Screen) {
	printBox2(s, 3, 2)
	printBox2(s, 3, 8)
	printBox2(s, 3, 14)
	printBox2(s, 3, 20)
	printBox2(s, 3, 26)
	printBox2(s, 43, 2)
	printBox2(s, 43, 8)
	printBox2(s, 43, 14)
	printBox2(s, 43, 20)
	printBox2(s, 43, 26)
}

func printBox2(s tcell.Screen, x int, y int) {

	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	emitStr(s, x, y, style, "+-----------------------------+")
	emitStr(s, x, y+4, style, "+-----------------------------+")
	emitStr(s, x, y+1, style, "|编号")
	emitStr(s, x, y+2, style, "|")
	emitStr(s, x, y+3, style, "|")
	emitStr(s, x+30, y+1, style, "|")
	emitStr(s, x+30, y+2, style, "|")
	emitStr(s, x+30, y+3, style, "|")

}
func printBox(s tcell.Screen, x int, y int) {
	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	emitStr(s, x, y, style, "+------------------------+")
	emitStr(s, x, y+5, style, "+------------------------+")
	emitStr(s, x, y+1, style, "|")
	emitStr(s, x, y+2, style, "|")
	emitStr(s, x, y+3, style, "|")
	emitStr(s, x, y+4, style, "|")
	emitStr(s, x+25, y+1, style, "|")
	emitStr(s, x+25, y+2, style, "|")
	emitStr(s, x+25, y+3, style, "|")
	emitStr(s, x+25, y+4, style, "|")

}

func infoMessageScreen(s tcell.Screen, messageList []model.ScreenInfoMessage) {
	_, h := s.Size()
	lenMessageList := len(messageList)
	for i := 0; i < lenMessageList; i++ {

		emitStr(s, 2, h-(lenMessageList-i)+1, tcell.StyleDefault, messageList[i].MessageStr)
	}
}
