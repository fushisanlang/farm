/*
@Time : 2022/4/14 4:39 下午
@Author : fushisanlang
@File : numtochinese
@Software: GoLand
*/
package tools

func NumToChinsesT(num int) string {
	var ChinesesT string
	switch num {
	case 1:
		ChinesesT = "甲"
	case 2:
		ChinesesT = "乙"
	case 3:
		ChinesesT = "丙"
	case 4:
		ChinesesT = "丁"
	case 5:
		ChinesesT = "戊"
	case 6:
		ChinesesT = "己"
	case 7:
		ChinesesT = "庚"
	case 8:
		ChinesesT = "辛"
	case 9:
		ChinesesT = "壬"
	case 10:
		ChinesesT = "癸"
	}
	return ChinesesT
}
func NumToChinsesD(num int) string {
	var ChinesesD string
	switch num {
	case 1:
		ChinesesD = "子"
	case 2:
		ChinesesD = "丑"
	case 3:
		ChinesesD = "寅"
	case 4:
		ChinesesD = "卯"
	case 5:
		ChinesesD = "辰"
	case 6:
		ChinesesD = "巳"
	case 7:
		ChinesesD = "午"
	case 8:
		ChinesesD = "未"
	case 9:
		ChinesesD = "申"
	case 10:
		ChinesesD = "酉"
	case 11:
		ChinesesD = "戌"
	case 12:
		ChinesesD = "亥"
	}
	return ChinesesD
}
