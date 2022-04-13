/*
@Time : 2022/4/12 9:48 上午
@Author : fushisanlang
@File : infoMessage
@Software: GoLand
*/
package screen

import "github.com/gdamore/tcell/v2"

func infoMessageScreen(s tcell.Screen) {
	_, h := s.Size()
	emitStr(s, 2, h-4, tcell.StyleDefault, "Press q to quit.")
	emitStr(s, 2, h-3, tcell.StyleDefault, "Press n to next.")
	emitStr(s, 2, h-2, tcell.StyleDefault, "Press b to back.")
	//emitStr(s, 2, h-1, tcell.StyleDefault, "Press  to exit.")

}
func createMessageScreen(s tcell.Screen) {
	_, h := s.Size()
	//	emitStrCenter(s, w, h/2, style, "道号 : "+service.UserInfo.UserName, 3)
	//	emitStrCenter(s, w, h/2+1, style, "道龄 : "+gconv.String(service.UserInfo.UserAge), 3)
	//	emitStrCenter(s, w, h/2+2, style, "道场 : "+service.UserInfo.FarmName, 3)
	//	emitStrCenter(s, w, h/2+3, style, "福地 : "+className, 3)
	//	emitStrCenter(s, w, h/2+4, style, "灵宠 : "+service.UserInfo.PetName, 3)
	emitStr(s, 2, h-7, tcell.StyleDefault, "按 1 修改道号")
	emitStr(s, 2, h-6, tcell.StyleDefault, "按 2 修改道龄")
	emitStr(s, 2, h-5, tcell.StyleDefault, "按 3 修改道场")
	emitStr(s, 2, h-4, tcell.StyleDefault, "按 4 修改福地")
	emitStr(s, 2, h-3, tcell.StyleDefault, "按 5 修改灵宠")
	emitStr(s, 2, h-2, tcell.StyleDefault, "按 q 退出程序")
	//emitStr(s, 2, h-1, tcell.StyleDefault, "Press  to exit.")

}
func infoMessageOnVerifyScreen(s tcell.Screen) {
	_, h := s.Size()
	emitStr(s, 2, h-2, tcell.StyleDefault, "Press q to quit.")

	//emitStr(s, 2, h-1, tcell.StyleDefault, "Press  to exit.")

}
