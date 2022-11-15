package methods

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/tools"
	"server/util"
	"time"
)

type OnLineTime struct {
	OnlineYear  int64 `json:"onlineYear"`
	FailYear    int   `json:"failYear"`
	OnlineMouth int64 `json:"onlineMouth"`
	FailMouth   int   `json:"failMouth"`
	OnlineDay   int64 `json:"onlineDay"`
	FailDay     int   `json:"failDay"`
}

func IndexGetFaultLine(c *gin.Context) {
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
	send := [][]int{reportcourt, processcourt}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"data": send,
		"code": 200,
	})
}

// 获取设备在线时长
func IndexGetOnLineTime(c *gin.Context) {
	var onlineDay []tools.FyDeviceTime
	var onlineYear []tools.FyDeviceTime
	var onlineMouth []tools.FyDeviceTime
	var failMouth []tools.FyFaillogs
	var failDay []tools.FyFaillogs
	var failYear []tools.FyFaillogs
	var result OnLineTime
	//获取当日的在线时长
	tools.DB.Where("online_time >= ?", util.GetZeroTime(time.Now())).Find(&onlineDay)
	tools.DB.Where("processtime >= ?", util.GetZeroTime(time.Now())).Find(&failDay)
	fmt.Println(util.GetZeroTime(time.Now()))
	result.FailDay = len(failDay)
	for _, value := range onlineDay {
		result.OnlineDay += (value.OutlineTime.Unix() - value.OnlineTime.Unix()) / 3600
	}
	// 获取当月的在线时长
	tools.DB.Where("online_time >= ?", util.GetFirstDateOfMonth(time.Now())).Find(&onlineYear)
	tools.DB.Where("processtime >= ?", util.GetFirstDateOfMonth(time.Now())).Find(&failMouth)
	result.FailMouth = len(failMouth)
	for _, value := range onlineYear {
		result.OnlineMouth += (value.OutlineTime.Unix() - value.OnlineTime.Unix()) / 3600
	}
	tools.DB.Where("online_time >= ?", util.GetFirstDateOfYear(time.Now())).Find(&onlineMouth)
	tools.DB.Where("processtime >= ?", util.GetFirstDateOfYear(time.Now())).Find(&failYear)
	result.FailYear = len(failYear)
	for _, value := range onlineMouth {
		result.OnlineYear += (value.OutlineTime.Unix() - value.OnlineTime.Unix()) / 3600
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"data": result,
		"code": 200,
	})
}

func GetCourtInfo(c *gin.Context) {
	type result struct {
		Name  string
		Zone  string
		State string
	}
	//获取法院信息
	var courtInfo []tools.FyCourtInfor
	//获取法院区域
	var zoneCourt []tools.FyZoneInfor
	tools.DB.Find(&zoneCourt)
	tools.DB.Find(&courtInfo)
	var courtData []result
	for _, v := range zoneCourt {
		for _, y := range courtInfo {
			if v.ZoneID == y.CourtZone {
				courtData = append(courtData, result{Name: y.CourtName, Zone: v.ZoneName, State: y.CourtState})
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"data": courtData,
		"code": 200,
	})
}

func GetZoneInfo(c *gin.Context) {
	type result struct {
		Name  string
		Zone  string
		State string
	}
	//获取法院信息
	var courtInfo []tools.FyCourtInfor
	//获取法院区域
	var zoneCourt []tools.FyZoneInfor
	tools.DB.Find(&zoneCourt)
	tools.DB.Find(&courtInfo)
	var courtData []result
	for _, v := range zoneCourt {
		for _, y := range courtInfo {
			if v.ZoneID == y.CourtZone {
				courtData = append(courtData, result{Name: y.CourtName, Zone: v.ZoneName, State: y.CourtState})
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"data": courtData,
		"code": 200,
	})
}
