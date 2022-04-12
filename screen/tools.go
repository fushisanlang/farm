/*
@Time : 2022/4/8 3:44 下午
@Author : fushisanlang
@File : tools
@Software: GoLand
*/
package screen

import (
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

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
	//居中计算长度是需要*2/3，取半数计算缩紧要/2，结果就是中文需要/3，英语需要/2。
	a := len(printString)
	emitStr(s, w/2-a/status, h, style, printString)

}
