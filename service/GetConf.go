/*
@Time : 2022/4/8 4:20 下午
@Author : fushisanlang
@File : GetConf
@Software: GoLand
*/
package service

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Conf struct {
	WeightMin int
	HeightMin int
	DbPath    string
}

func GetConf() Conf {
	var Conf Conf
	var path string = "./conf/conf.toml"
	if _, err := toml.DecodeFile(path, &Conf); err != nil {
		log.Fatal(err)
	}
	return Conf
}
