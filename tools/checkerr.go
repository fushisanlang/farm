/*
@Time : 2022/4/11 9:30 上午
@Author : fushisanlang
@File : CheckErr
@Software: GoLand
*/
package tools

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"os"
)

func CheckErr(err error) {
	if err != nil {
		g.Log().Println(err)
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
