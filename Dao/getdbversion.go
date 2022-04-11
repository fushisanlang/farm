/*
@Time : 2022/4/11 9:25 上午
@Author : fushisanlang
@File : GetConf
@Software: GoLand
*/
package Dao

import (
	//"database/sql"
	//"farm/tools"
	"database/sql"
	"farm/tools"

	_ "github.com/mattn/go-sqlite3"
)

var DbPath string

func GetDbVersion() string {

	db, err := sql.Open("sqlite3", DbPath)
	defer db.Close()
	rows, err := db.Query(`SELECT confValue FROM conf where confKey = 'version'`)
	tools.CheckErr(err)
	var confValue string
	for rows.Next() {

		err = rows.Scan(&confValue)
		tools.CheckErr(err)

	}
	return confValue

}
