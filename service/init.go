/*
@Time : 2022/4/11 9:14 上午
@Author : fushisanlang
@File : init
@Software: GoLand
*/
package service

var Version string

func init() {
	Version = GetDbVersion()
}
