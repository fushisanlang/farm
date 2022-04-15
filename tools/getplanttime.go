/*
@Time : 2022/4/15 10:31 上午
@Author : fushisanlang
@File : getplanttime
@Software: GoLand
*/
package tools

import "github.com/gogf/gf/util/gconv"

func GetPlantTime(duringtime int, neettime int) string {
	var PlantStatus string
	if gconv.Float32(duringtime)/gconv.Float32(neettime) >= 1 {
		PlantStatus = "成熟"
	} else if gconv.Float32(duringtime)/gconv.Float32(neettime) > 0.8 {
		PlantStatus = "结果"

	} else if gconv.Float32(duringtime)/gconv.Float32(neettime) > 0.6 {
		PlantStatus = "开花"

	} else if gconv.Float32(duringtime)/gconv.Float32(neettime) > 0.4 {
		PlantStatus = "发芽"

	} else if gconv.Float32(duringtime)/gconv.Float32(neettime) > 0.2 {
		PlantStatus = "生根"

	} else {
		PlantStatus = "播种"

	}

	return PlantStatus

}
