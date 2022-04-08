/*
@Time : 2022/4/8 3:51 下午
@Author : fushisanlang
@File : VerifySize
@Software: GoLand
*/
package screen

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/gogf/gf/util/gconv"
	"os"
	"tcellLearn/service"
	"time"
)

func VerifySize(s tcell.Screen) {
	Conf := service.GetConf()
	WeightMin := gconv.Int(Conf.WeightMin)
	HeightMin := gconv.Int(Conf.HeightMin)
	w, h := s.Size()
	fmt.Println(w, h)
	if w < WeightMin || h < HeightMin {
		s.Fini()
		fmt.Println("请将窗口大小调整为" + gconv.String(WeightMin) + "*" + gconv.String(HeightMin) + "，目前窗口大小为" + gconv.String(w) + "*" + gconv.String(h))
		time.Sleep(10 * time.Second)
		os.Exit(1)

	}
}
