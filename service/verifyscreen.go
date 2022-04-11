/*
@Time : 2022/4/11 10:16 上午
@Author : fushisanlang
@File : verifyscreen
@Software: GoLand
*/
package service

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/gogf/gf/util/gconv"
)

func VerifySize(s tcell.Screen) {
	WeightMin := 100
	HeightMin := 50
	w, h := s.Size()
	fmt.Println(w, h)
	if w < WeightMin || h < HeightMin {
		s.Fini()
		fmt.Println("请将窗口大小调整为" + gconv.String(WeightMin) + "*" + gconv.String(HeightMin) +
			"，目前窗口大小为" + gconv.String(w) + "*" + gconv.String(h))
		time.Sleep(1 * time.Second)
		os.Exit(1)

	}
}
