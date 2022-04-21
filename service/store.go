/*
@Time : 2022/4/21 10:10
@Author : fushisanlang
@File : store
@Software: GoLand
*/
package service

import (
	"math/rand"
	"time"
)

func ChangeRatioAuto() {
	changeRatioBuyAuto()
	changeRatioSaleAuto()
}
func changeRatioBuyAuto() {
	rand.Seed(time.Now().UnixNano())
	RatioBuy = 90 + rand.Float32()*20

}
func changeRatioSaleAuto() {
	rand.Seed(time.Now().UnixNano())
	RatioSale = 90 + rand.Float32()*20

}
