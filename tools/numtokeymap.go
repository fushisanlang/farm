/*
@Time : 2022/4/14 4:39 下午
@Author : fushisanlang
@File : numtochinese
@Software: GoLand
*/
package tools

func NumToKeyMap(num int) string {
	var KeyMap string
	switch num {
	case 1:
		KeyMap = "1"
	case 2:
		KeyMap = "2"
	case 3:
		KeyMap = "3"
	case 4:
		KeyMap = "4"
	case 5:
		KeyMap = "5"
	case 6:
		KeyMap = "6"
	case 7:
		KeyMap = "7"
	case 8:
		KeyMap = "8"
	case 9:
		KeyMap = "9"
	case 10:
		KeyMap = "0"

	}

	return KeyMap
}
