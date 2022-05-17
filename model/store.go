/*
@Time : 2022/4/22 10:15
@Author : fushisanlang
@File : store
@Software: GoLand
*/
package model

type GoodsList struct {
	KeyId    int
	Name     string
	Countnum int
	Price    int
	GroupId  int
}

type BuyList struct {
	Id      int
	Name    string
	Price   int
	GroupId int
}
