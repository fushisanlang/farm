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
		KeyMap = "Ctrl + q"
	case 2:
		KeyMap = "Ctrl + w"
	case 3:
		KeyMap = "Ctrl + e"
	case 4:
		KeyMap = "Ctrl + r"
	case 5:
		KeyMap = "Ctrl + a"
	case 6:
		KeyMap = "Ctrl + s"
	case 7:
		KeyMap = "Ctrl + d"
	case 8:
		KeyMap = "Ctrl + f"
	case 9:
		KeyMap = "Ctrl + z"
	case 10:
		KeyMap = "Ctrl + x"
	case 11:
		KeyMap = "Ctrl + c"
	case 12:
		KeyMap = "Ctrl + v"
	}

	return KeyMap
}
