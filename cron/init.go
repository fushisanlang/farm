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
	autoRefreshStoreBuyList()
	gcron.Add("* * * * * *", duringTimeAutoAdd)
	gcron.Add("0 0 * * * *", autoRefreshStoreBuyList)

}

func duringTimeAutoAdd() {
	service.DuringTimeAutoAdd()
}

func autoRefreshStoreBuyList() {
	service.GetBuyGroceriesList()
	service.GetBuySeedList()
	service.GetBuyPlantList()
}
