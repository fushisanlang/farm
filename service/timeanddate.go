/*
@Time : 2022/5/16 14:18
@Author : fushisanlang
@File : timeanddate
@Software: GoLand
*/
package service

import (
	"time"
)

func TimeAndDate() (string, string) {
	now := time.Now()
	DateNow := now.Format("2006/1/02")
	TimeNow := now.Format("15:04:05")
	return DateNow, TimeNow
}
