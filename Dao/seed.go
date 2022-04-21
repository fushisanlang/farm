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
	AND p.level <= ( SELECT level from userinfo where id = 1 ) 
	AND b.linkid = p.id
`
	SeedList, _ := g.DB().GetAll(sql)

	g.Log().Println("SeedList " + gconv.String(SeedList))

	return SeedList

}

func plantSeedUpdateField(screenId int, fieldId int, Id int, bagId int) {
	/*
		sql := `UPDATE "main"."field"
				SET
					"plantid" = ( SELECT linkid FROM bag WHERE id = ? AND bagid = ? ),
					"duringtime" = 0,
					"timeneed" = (select timeneed from plant where id =
						(SELECT linkid FROM bag WHERE id = ? AND bagid = ? ))
				WHERE
					"id" = ?
					AND "fieldid" = ?;`

	*/
	sql := `UPDATE "main"."field" 
			SET 
				"plantid" = ( SELECT linkid FROM bag WHERE id = ` + gconv.String(Id) + ` AND bagid =  ` + gconv.String(bagId) + ` ),
				"duringtime" = 0,
				"timeneed" = (select timeneed from plant where id = 
					(SELECT linkid FROM bag WHERE id =  ` + gconv.String(Id) + ` AND bagid =  ` + gconv.String(bagId) + ` ))
			WHERE
				"id" =  ` + gconv.String(screenId) + `
				AND "fieldid" =  ` + gconv.String(fieldId) + `;`
	_, err := g.DB().Exec(sql)
	g.Log().Println(sql)

	tools.CheckErr(err)

}
func plantSeedUpdateUserInfo(Id int, bagId int) {
	sql := `UPDATE userinfo 
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
	_, err := g.DB().Exec(sql, Id, bagId, Id, bagId)
	tools.CheckErr(err)
}

func plantSeedUpdateBag(Id int, bagId int) {
	sql := `UPDATE bag 
			SET 
				countnum = countnum - 1 
			WHERE
				id = ?
				AND bagid = ?;
`
	_, err := g.DB().Exec(sql, Id, bagId)
	tools.CheckErr(err)
}
func plantSeedUpdateBagClean() {
	sql := `update bag 
			set 
				linkid = 0 , 
				groupid = 0,
				countnum = 0  
			where 
				countnum <= 0;
`
	_, err := g.DB().Exec(sql)
	tools.CheckErr(err)
}

func PlantSeed(screenId int, fieldId int, Id int, bagId int) {
	g.Log().Println("screenId :" + gconv.String(screenId))
	g.Log().Println("fieldId :" + gconv.String(fieldId))
	g.Log().Println("Id :" + gconv.String(Id))
	g.Log().Println("bagId :" + gconv.String(bagId))

	plantSeedUpdateField(screenId, fieldId, Id, bagId)
	plantSeedUpdateUserInfo(Id, bagId)
	plantSeedUpdateBag(Id, bagId)
	plantSeedUpdateBagClean()

}
