/*
@Time : 2022/4/19 3:10 下午
@Author : fushisanlang
@File : init
@Software: GoLand
*/
package cron

import (
	"farm/service"

	"github.com/gogf/gf/os/gcron"
)

func init() {
	gcron.Add("* * * * * *", duringTimeAutoAdd)

}

func duringTimeAutoAdd() {
	service.DuringTimeAutoAdd()
}
