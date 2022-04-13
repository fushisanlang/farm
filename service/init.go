/*
@Time : 2022/4/11 9:14 上午
@Author : fushisanlang
@File : init
@Software: GoLand
*/
package service

import "farm/model"

var UserInfo model.UserInfo
var Version string

func init() {
	Version = "0.0.1"
}
