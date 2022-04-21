/*
@Time : 2022/4/19 4:22 下午
@Author : fushisanlang
@File : eradicate
@Software: GoLand
*/
package Dao

import (
	"farm/model"
	"farm/tools"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func HarvestPlant(id int, fieldId int) gdb.Record {
	sql := `SELECT
	p.plantname plantname,
	CAST ( p.yield / 2 * ( 1 + ( RANDOM( ) + 9223372036854775808 ) / 2.0 / 9223372036854775808 + 1) AS INTEGER ) yield ,
	f.plantid plantid
FROM
	field f,
	plant p 
WHERE
	f.plantid = p.id 
	AND f.duringtime >= p.timeneed 
	AND f.id = ?
	AND f.fieldid = ?`

	harvest, err := g.DB().GetOne(sql, id, fieldId)
	tools.CheckErr(err)

	return harvest

}
func GetHarvestBagId(Harvest model.Harvest) int {
	sql := `SELECT
CASE
	WHEN
			(SELECT keyid FROM bag WHERE linkid = ? AND groupid = 2 ) ISNULL 
	THEN
			(SELECT keyid FROM bag WHERE linkid = 0 AND countnum = 0 AND groupid = 0 ORDER BY keyid LIMIT 1 ) 
	ELSE 
			(SELECT keyid FROM bag WHERE linkid = ? AND groupid = 2 ) 
	END as id`
	BagId, err := g.DB().GetOne(sql, Harvest.PlantId, Harvest.PlantId)
	tools.CheckErr(err)
	if gconv.String(BagId["id"]) == "NULL" {
		return 0
	} else {
		return gconv.Int(BagId["id"])
	}

}
func HarvestBag(Harvest model.Harvest, bagId int) {
	//加背包
	sql := `update bag set countnum = countnum + ?, groupid = 2, linkid = ? WHERE keyid =  ?`

	_, err := g.DB().Exec(sql, Harvest.Yield, Harvest.PlantId, bagId)
	tools.CheckErr(err)
}
func HarvestField(id int, fieldId int) {
	//清理土地
	sql := `
UPDATE field 
SET duringtime = ( SELECT p.timeneed * 0.2 FROM field f, plant p WHERE f.plantid = p.id AND f.id = ? AND f.fieldid = ? ) 
WHERE
	id = ? 
	AND fieldid =?
`
	_, err := g.DB().Exec(sql, id, fieldId, id, fieldId)
	tools.CheckErr(err)

}
func HarvestEx(Harvest model.Harvest) {
	sql := `
UPDATE userinfo 
SET ex =( ( SELECT ex FROM plant WHERE id = ? ) * ?  + ex ) % ( ( level + 1 ) * 100 ),
level = level + ( ( SELECT ex FROM plant WHERE id = ? ) * ? + ex ) / ( ( level + 1 ) * 100 ) 
WHERE
	id = 1;
`
	_, err := g.DB().Exec(sql, Harvest.PlantId, Harvest.Yield, Harvest.PlantId, Harvest.Yield)
	tools.CheckErr(err)

}
