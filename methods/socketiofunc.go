package methods

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"server/datastore"
	"server/qsc"
	"server/tools"
	"strings"
	"time"
)

// 接收处理器回复的确认信息
type Returnmsg struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  bool   `json:"result"`
	Id      int    `json:"id"`
}

// 用于登录页面用户信息接收以及验证
func Userauth(s socketio.Conn, msg datastore.UsrTables) {
	var userinfor tools.FyUsers
	tools.DB.Where("user_name=?", msg.User).Find(&userinfor)
	if userinfor.UserName == msg.User {
		if userinfor.UserPasswd == msg.Passwd {
			token, err := GenerateToken(msg.User, msg.Passwd)
			if err != nil {
				log.Println("token生成错误")
			}
			s.Emit("recvToken", token)
		} else {
			s.Emit("recvError", 0)
		}
	} else {
		s.Emit("recvError", 1)
	}
}

// 用于远程监听页面数据接收
func Getycjt(s socketio.Conn, deviceid string) {
	var line []tools.FyDeviceChannel
	tools.DB.Where("deviceid = ?", deviceid).Find(&line)
	s.Emit("recvYcjtinit", line)
}

// 用于前台菜单的获取
func MenuList(s socketio.Conn, token string) {
	//用于解密token
	user, err := ParseToken(token)
	if err != nil {
		log.Println("获取前台菜单时，token解密出错")
	}
	//获取用户表
	var privileges tools.FyUsers
	tools.DB.Where("user_name = ?", user.Username).Find(&privileges)
	//获取菜单从数据库内
	var navigate []tools.FyNavigateMenu
	tools.DB.Where("menu_privileges >= ?", privileges.Privileges).Find(&navigate)
	s.Emit("revMenuList", navigate)
}

// 用于获取设备URL
func Geturl(s socketio.Conn, token string, id string) {
	//用于解密token
	user, err := ParseToken(token)
	if err != nil {
		log.Println("获取设备URL时，token解密出错")
	}
	if id == "0" {
		//获取用户表
		var courtid tools.FyUsers
		tools.DB.Where("user_name = ?", user.Username).Find(&courtid)
		//获取URL从数据库内
		var deviceurl []tools.FyDeviceInfor
		tools.DB.Where("court_id = ?", courtid.CourtID).Find(&deviceurl)
		s.Emit("recvdeviceadd", deviceurl)
	} else {
		//获取URL从数据库内
		var deviceurl []tools.FyDeviceInfor
		tools.DB.Where("court_id = ?", id).Find(&deviceurl)
		s.Emit("recvdeviceadd", deviceurl)
	}
}

// 用于系统日志的记录
func Operatelogs(s socketio.Conn, msg tools.FySystemlogs) {
	user, err := ParseToken(msg.UserName)
	if err != nil {
		log.Println("获取设备URL时，token解密出错")
	}
	ipx := s.RemoteAddr().String()
	ip, _, _ := strings.Cut(ipx, ":")
	tools.DB.Create(&tools.FySystemlogs{
		Time:     time.Now().Format("2006-01-02 15:04:05"),
		UserName: user.Username,
		Content:  msg.Content,
		Address:  ip,
	})
}

// 用于获取系统日志数据
func Getoperatelogs(s socketio.Conn, token string) {
	user, err := ParseToken(token)
	if err != nil {
		log.Println("获取设备URL时，token解密出错")
	}
	//获取用户表
	var systemlogs []tools.FySystemlogs
	if user.Username == "admin" {
		tools.DB.Order("time desc").Find(&systemlogs)
		s.Emit("recvoperateLogs", systemlogs)
	} else {
		tools.DB.Where("user_name = ?", user.Username).Order("time desc").Find(&systemlogs)
		s.Emit("recvoperateLogs", systemlogs)
	}

}

// 用于故障日志的获取
func Getfaillogs(s socketio.Conn) {
	var faillogs []tools.FyFaillogs
	var courtinfor []tools.FyCourtInfor
	var deviceinfor []tools.FyDeviceInfor
	tools.DB.Order("reporttime desc").Find(&faillogs)
	tools.DB.Find(&courtinfor)
	tools.DB.Find(&deviceinfor)
	for i := 0; i < len(faillogs); i++ {
		for y := 0; y < len(deviceinfor); y++ {
			if faillogs[i].Deviceid == deviceinfor[y].Deviceid {
				faillogs[i].Deviceid = deviceinfor[y].DeviceName
			}
		}
		for v := 0; v < len(courtinfor); v++ {
			if faillogs[i].Courtid == courtinfor[v].CourtID {
				faillogs[i].Courtid = courtinfor[v].CourtName
			}
		}
	}
	s.Emit("recvfailLogs", faillogs)
}

// 用于获取监测管理页面初始数据
func Monitorinit(s socketio.Conn, token string) {
	user, err := ParseToken(token)
	if err != nil {
		log.Println("获取监测管理页面初始数据时，token解密出错")
	}
	//获取用户权限
	var privileges tools.FyUsers
	//获取法院信息
	var courtInfor []tools.FyCourtInfor
	//获取法院区域
	var zonecourt []tools.FyZoneInfor
	tools.DB.Find(&zonecourt)
	tools.DB.Where("user_name = ?", user.Username).Find(&privileges)
	//判断用户权限发送数据
	courtdata := make(map[string][]interface{})
	if privileges.Privileges == 0 || privileges.Privileges == 1 {
		tools.DB.Find(&courtInfor)
		for _, v := range zonecourt {
			for _, y := range courtInfor {
				if v.ZoneID == y.CourtZone {
					courtdata[v.ZoneName] = append(courtdata[v.ZoneName], map[string]interface{}{"name": y.CourtName, "id": y.CourtID, "sort": y.CourtSort})
				}
			}
		}
	} else if privileges.Privileges == 2 {
		tools.DB.Where("court_zone = ?", privileges.ZoneID).Find(&courtInfor)
		var zonename tools.FyZoneInfor
		tools.DB.Where("zone_id = ?", privileges.ZoneID).Find(&zonename)
		for _, y := range courtInfor {
			courtdata[zonename.ZoneName] = append(courtdata[zonename.ZoneName], map[string]interface{}{"name": y.CourtName, "id": y.CourtID, "sort": y.CourtSort})
		}
	}
	s.Emit("monitorinit", courtdata)
}

// 用于监测页面的故障日志添加接口
func Monitorerrinput(s socketio.Conn, courtid string, errinput string, token string, deviceid string) {
	user, err := ParseToken(token)
	if err != nil {
		log.Println("获取设备URL时，token解密出错")
	}
	tools.DB.Create(&tools.FyFaillogs{
		Failid:      fmt.Sprintf("%v", time.Now().UnixNano()),
		Failcontent: errinput,
		Reporttime:  time.Now().Format("2006-01-02 15:04:05"),
		Reportuser:  user.Username,
		Courtid:     courtid,
		Deviceid:    deviceid,
		Repair:      "0",
	})

}

// 用于获取故障状态中的设备列表
func Getfaildevice(s socketio.Conn, courtid string) {
	var device []tools.FyDeviceInfor
	tools.DB.Where("court_id = ?", courtid).Find(&device)
	s.Emit("recvdevicelist", device)
}

// 用于打开远程监听
func Listenopen(line string, deviceid string, comm string) {
	qsc.RemoteListen(line, deviceid, comm)
}

// 故障处理页面获取页面信息方法
func Failrepair(s socketio.Conn, courtid string) {
	var failinfor []tools.FyFaillogs
	tools.DB.Where("courtid = ? and repair = ?", courtid, "0").Find(&failinfor)
	s.Emit("recvfailinfor", failinfor)
}

// 用于关闭远程监听方法
func Closeycjtbtn(s socketio.Conn) {
	qsc.CloseFeeds("changegroup_3", s)
}

// 用于下载文件列表
func Downfileinit(s socketio.Conn, courtid string) {
	var filelist []tools.FyDownFile
	tools.DB.Where("courtid =?", courtid).Find(&filelist)
	for i, v := range filelist {
		filelist[i].Filesavename = "http://127.0.0.1/" + v.Filesavename
	}
	s.Emit("downfileinit", filelist)
}

// 用于获取主页的统计数据
func Getindexcourt(s socketio.Conn) {
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
	s.Emit("recvcourt", send)
}

// 用于更新设备故障库内的信息
func Failinforsumit(s socketio.Conn, failid string, failmsg string, token string) {
	user, err := ParseToken(token)
	if err != nil {
		log.Println("获取监测管理页面初始数据时，token解密出错")
	}
	updates := tools.FyFaillogs{
		Failid:         failid,
		Processcontent: failmsg,
		Processtime:    time.Now().Format("2006-01-02 15:04:05"),
		Processuser:    user.Username,
		Repair:         "1",
	}
	var temp tools.FyFaillogs
	tools.DB.Model(&temp).Where("failid = ?", failid).Updates(&updates)

}

// 用于获取法院状态
func Getcourtstate(s socketio.Conn) {
	type FyCourtInfor struct {
		CourtID    string
		CourtState string
	}
	go func() {
		for {
			var statedata []FyCourtInfor
			tools.DB.Find(&statedata)
			v := fmt.Sprintf("%v", qsc.QscData["serverTime"])
			statedata = append(statedata, FyCourtInfor{CourtID: "serverTime", CourtState: v})
			s.Emit("recvcourtstate", statedata)
			statedata = []FyCourtInfor{}
			time.Sleep(1 * time.Second)
		}
	}()
}
func Devicelocation(courtid string) Returnmsg {
	//v := tools.NetTcp()
	//tools.Pools = append(tools.Pools, v)
	return Returnmsg(qsc.DevicePosition(courtid))
}

func Tcpconn(s socketio.Conn) {
	v := tools.NetTcp()
	v.Id = s.ID()
	tools.Pools = append(tools.Pools, v)
	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(400 * time.Millisecond)
		}
	}()
}

func Closeconn(s socketio.Conn) {
	for i := 0; i < len(tools.Pools); i++ {
		if s.ID() == tools.Pools[i].Id {
			tools.Pools[i].Conn.Write([]byte(s.ID()))
			tools.Pools[i].Conn.Write([]byte("close!"))
			tools.Pools[i].Conn.Close()
		}
	}
}
