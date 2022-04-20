/*
@Time : 2022/4/11 9:25 上午
@Author : fushisanlang
@File : GetConf
@Software: GoLand
*/
package Dao

import (
	"farm/tools"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	_ "github.com/mattn/go-sqlite3"
)

func GetSeedList() gdb.Result {
	sql := `
SELECT
	b.id id,
	b.bagid bagid,
	p.plantname plantname,
	p.timeneed timeneed,
	b.countnum countnum,
	p.price price,
	p.ex ex 
FROM
	bag b,
	plant p 
WHERE
	b.groupid = 1 
	AND b.countnum > 0 
	AND p.level <= ( SELECT level from userinfonew where id = 1 ) 
	AND b.linkid = p.id
`
	SeedList, _ := g.DB().GetAll(sql)
	return SeedList

}
func PlantSeed(screenId int, fieldId int, Id int, bagId int) {

	sql := `UPDATE "main"."field" 
			SET 
				"plantid" = ( SELECT linkid FROM bag WHERE id = ? AND bagid = ? ),
				"duringtime" = 0,
				"timeneed" = (select timeneed from plant where id = 
					(SELECT linkid FROM bag WHERE id = ? AND bagid = ? ))
			WHERE
				"id" = ?
				AND "fieldid" = ?;
-- -------------------
			UPDATE bag 
			SET 
				countnum = countnum - 1 
			WHERE
				id = ?
				AND bagid = ?; 
-- -------------------
			update bag 
			set 
				linkid = 0 , 
				groupid = 0,
				countnum = 0  
			where 
				countnum <= 0;
-- -------------------
			UPDATE userinfonew 
			SET 
				ex =( 
					( SELECT ex FROM plant WHERE id =( SELECT linkid FROM bag WHERE id = ? AND bagid = ? ) )+ ex ) 
						% 
					( ( level + 1 ) * 100 ),
				level = level +( 
					( SELECT ex FROM plant WHERE id = ( SELECT linkid FROM bag WHERE id = ? AND bagid = ? ) )+ ex ) 
						/
					( ( level + 1 ) * 100 ) 
			WHERE
				id = 1;
`

	g.Log().Println("播种 " + gconv.String(screenId) + " " + gconv.String(fieldId) +
		" " + gconv.String(Id) + " " + gconv.String(bagId) + " ")

	_, err := g.DB().Exec(sql, Id, bagId, Id, bagId, screenId, fieldId, Id, bagId, Id, bagId, Id, bagId)
	tools.CheckErr(err)

}
