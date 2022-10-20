package main

import (
	"fmt"
	"server/tools"
)

func Test() {
	tools.DB = tools.DatabaseConn()
	reportcourt := []int{}
	processcourt := []int{}
	var reportlogs []tools.FyFaillogs
	var processlogs []tools.FyFaillogs
	datetime := []string{"%2022-01%", "%2022-02%", "%2022-03%", "%2022-04%", "%2022-05%", "%2022-06%", "%2022-07%", "%2022-08%", "%2022-09%", "%2022-10%", "%2022-11%", "%2022-12%"}
	for _, i := range datetime {
		tools.DB.Where("reporttime LIKE ?", i).Find(&reportlogs)
		tools.DB.Where("processtime LIKE ?", i).Find(&processlogs)
		reportcourt = append(reportcourt, len(reportlogs))
		processcourt = append(processcourt, len(processlogs))
	}
	fmt.Println(reportcourt, processcourt)

}

// 用于前端法院列表的获取
func ceshi() {
	db := tools.DatabaseConn()
	var courtData []tools.FyCourtInfor
	var zoneData []tools.FyZoneInfor
	db.Find(&courtData)
	db.Find(&zoneData)
	fmt.Println(courtData)
	fmt.Println(zoneData)
	courtinfortable := make(map[string]map[string]map[string]interface{})
	for _, i := range zoneData {
		courtinfortable[i.ZoneName] = map[string]map[string]interface{}{}
		for _, v := range courtData {
			if i.ZoneID == v.CourtZone {
				for _, x := range courtinfortable {
					fmt.Println(x)
				}
			}
		}
	}
	courtinfortable["hz"] = map[string]map[string]interface{}{}
	for _, v := range courtinfortable {
		v["sdfa"] = map[string]interface{}{}
		for _, i := range v {
			i["hjh"] = 77
			fmt.Println(i)
		}
	}
	fmt.Println(courtinfortable)
}
