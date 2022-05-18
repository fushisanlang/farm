/*
@Time : 2022/4/8 3:37 下午
@Author : fushisanlang
@File : hello.go
@Software: GoLand
*/
package screen

import (
	"farm/model"
	"farm/service"
	"github.com/gdamore/tcell/v2"
	"strings"
	"time"
)

func FeildSummaryScreen(s tcell.Screen) {
	for service.RefreserCode == 0 {
		feildSummaryScreen(s)
		time.Sleep(1 * time.Second)
	}
}

func feildSummaryScreen(s tcell.Screen) {
	FiledSummaryList := service.FiledSummary()
	s.Clear()

	j := 0
	for i := 1; i <= 40; i++ {
		id := 4 * (i - j - 1)
		if i%4 != 0 {
			//1
			//未开启土地显示空进度条，两端灰色
			if FiledSummaryList[id].IsOpen == 0 {
				emitStr(s, 2, i, styleGray, "["+strings.Repeat(" ", 30)+"]")
				//未种植土地显示空进度条，两端蓝色
			} else if FiledSummaryList[id].NTime == 0 {
				emitStr(s, 2, i, styleBlue, "["+strings.Repeat(" ", 30)+"]")
				//成熟土地显示｜  绿色
			} else if FiledSummaryList[id].DTime == FiledSummaryList[id].NTime {
				emitStr(s, 2, i, styleGreen, "["+strings.Repeat("|", 30)+"]")

				//成长中土地，显示｜ 蓝色
			} else {
				a := FiledSummaryList[id].DTime * 30 / FiledSummaryList[id].NTime
				b := 30 - a
				strRepeat := strings.Repeat("|", a) + strings.Repeat(" ", b)
				emitStr(s, 2, i, styleBlue, "["+strRepeat+"]")
			}

			//2
			if FiledSummaryList[id+1].IsOpen == 0 {
				emitStr(s, 37, i, styleGray, "["+strings.Repeat(" ", 30)+"]")
				//未种植土地显示空进度条，两端蓝色
			} else if FiledSummaryList[id+1].NTime == 0 {
				emitStr(s, 37, i, styleBlue, "["+strings.Repeat(" ", 30)+"]")
				//成熟土地显示｜  绿色
			} else if FiledSummaryList[id+1].DTime == FiledSummaryList[id+1].NTime {
				emitStr(s, 37, i, styleGreen, "["+strings.Repeat("|", 30)+"]")

				//成长中土地，显示｜ 蓝色
			} else {
				a := FiledSummaryList[id+1].DTime * 30 / FiledSummaryList[id+1].NTime
				b := 30 - a
				strRepeat := strings.Repeat("|", a) + strings.Repeat(" ", b)
				emitStr(s, 37, i, styleBlue, "["+strRepeat+"]")
			}

			if FiledSummaryList[id+2].IsOpen == 0 {
				emitStr(s, 72, i, styleGray, "["+strings.Repeat(" ", 30)+"]")
				//未种植土地显示空进度条，两端蓝色
			} else if FiledSummaryList[id+2].NTime == 0 {
				emitStr(s, 72, i, styleBlue, "["+strings.Repeat(" ", 30)+"]")
				//成熟土地显示｜  绿色
			} else if FiledSummaryList[id+2].DTime == FiledSummaryList[id+2].NTime {
				emitStr(s, 72, i, styleGreen, "["+strings.Repeat("|", 30)+"]")

				//成长中土地，显示｜ 蓝色
			} else {
				a := FiledSummaryList[id+2].DTime * 30 / FiledSummaryList[id+2].NTime
				b := 30 - a
				strRepeat := strings.Repeat("|", a) + strings.Repeat(" ", b)
				emitStr(s, 72, i, styleBlue, "["+strRepeat+"]")
			}

			if FiledSummaryList[id+3].IsOpen == 0 {
				emitStr(s, 107, i, styleGray, "["+strings.Repeat(" ", 30)+"]")
				//未种植土地显示空进度条，两端蓝色
			} else if FiledSummaryList[id+3].NTime == 0 {
				emitStr(s, 107, i, styleBlue, "["+strings.Repeat(" ", 30)+"]")
				//成熟土地显示｜  绿色
			} else if FiledSummaryList[id+3].DTime == FiledSummaryList[id+3].NTime {
				emitStr(s, 107, i, styleGreen, "["+strings.Repeat("|", 30)+"]")

				//成长中土地，显示｜ 蓝色
			} else {
				a := FiledSummaryList[id+3].DTime * 30 / FiledSummaryList[id+3].NTime
				b := 30 - a
				strRepeat := strings.Repeat("|", a) + strings.Repeat(" ", b)
				emitStr(s, 107, i, styleBlue, "["+strRepeat+"]")
			}

		} else {
			j++

		}

	}
	infoMessageList := []model.ScreenInfoMessage{nPress, qPress}
	infoMessageScreen(s, infoMessageList)
	s.Show()
}
