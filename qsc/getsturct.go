package qsc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net"
	"net/http"
	"server/tools"
	"strings"
	"time"
)

var QscData = make(map[string]interface{})
var Courtdate = make(map[string]interface{})

type Header struct {
	Jsonrpc string      `json:"jsonrpc"`
	Id      int         `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

// 获取全局设备信息结构
type getGlobalDeviceStatusParams struct {
	Id        string
	Component GlobalDeviceStatusData
}

type GlobalDeviceStatusData struct {
	Name     string
	Controls []map[string]string
}

// 自动获取全局数据结构
type sumitrequest struct {
	Id   string
	Rate int
}

// Returnmsg 接收处理器回复的确认信息
type Returnmsg struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  bool   `json:"result"`
	Id      int    `json:"id"`
}

// 接收订阅内容（服务器时间，法院状态等）
type recvglobaldata struct {
	Jsonrpc string               `json:"jsonrpc"`
	Method  string               `json:"method"`
	Params  recvglobaldataparams `json:"params"`
}
type recvglobaldataparams struct {
	Id      string
	Changes []map[string]interface{}
}

// 心跳数据
type heartbeat struct {
	Jsonrpc string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  map[string]string `json:"params"`
}

func GetStatus(ctx *gin.Context) {
	v := tools.NetTcp()
	a := Header{Jsonrpc: "2.0", Id: 1234, Method: "StatusGet",
		Params: 0}
	jsona, _ := json.Marshal(a)
	v.Conn.Write([]byte(string(jsona) + "\x00"))
	var x [1024]byte
	var msg Returnmsg
	readSize, _ := v.Conn.Read(x[0:])
	json.Unmarshal(x[0:readSize], &msg)
	fmt.Println(msg)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

// 获取全局数据及心跳
func GetGlobalData() {
	v := tools.NetTcp()
	heartbeatdata := heartbeat{Jsonrpc: "2.0", Method: "NoOp", Params: map[string]string{}}
	go func() {
		for {
			jsonheart, _ := json.Marshal(heartbeatdata)
			v.Conn.Write([]byte(string(jsonheart) + "\x00"))
			time.Sleep(50 * time.Second)
		}
	}()
	var msg Returnmsg
	var globaldata recvglobaldata
	command := []map[string]string{}
	command = append(command, map[string]string{"Name": "serverTime"})
	command = append(command, map[string]string{"Name": "tcp.status 1"})
	command = append(command, map[string]string{"Name": "core.Status.out 1"})
	a := Header{Jsonrpc: "2.0", Id: 1234, Method: "ChangeGroup.AddComponentControl",
		Params: getGlobalDeviceStatusParams{Id: "my_change_group_Status", Component: GlobalDeviceStatusData{Name: "web_Global_Device_Status",
			Controls: command}}}
	jsona, _ := json.Marshal(a)
	fmt.Println("111111111111111111111111111111111111", string(jsona))
	b := Header{Jsonrpc: "2.0", Id: 1234, Method: "ChangeGroup.AutoPoll", Params: sumitrequest{Id: "my_change_group_Status", Rate: 1}}
	jsonb, _ := json.Marshal(b)
	time.Sleep(200 * time.Millisecond)
	var x [1024]byte
	//发送需要定阅的内容
	v.Conn.Write([]byte(string(jsona) + "\x00"))
	readSize, _ := v.Conn.Read(x[0:])
	json.Unmarshal(x[0:readSize], &msg)
	if msg.Result {
		fmt.Println("订阅内容成功")
	}
	//发送自动订阅指令
	v.Conn.Write([]byte(string(jsonb) + "\x00"))
	readSize, _ = v.Conn.Read(x[0:])
	json.Unmarshal(x[0:readSize], &msg)
	fmt.Println("111111111111111111111111111111111111", string(jsonb))
	if msg.Result {
		var courtInfor tools.FyCourtInfor
		fmt.Println("自动订阅成功")
		for {
			readSize, _ = v.Conn.Read(x[0:])
			json.Unmarshal(x[0:readSize], &globaldata)
			go func() {
				for _, v := range globaldata.Params.Changes {
					if v["Name"] == "serverTime" {
						QscData["serverTime"] = v["String"]
					} else {
						str := fmt.Sprintf("%v", v["Name"])
						slic := strings.Split(str, " ")
						if slic[0] == "tcp.status" {
							str = fmt.Sprintf("%v", v["String"])
							if str == "true" {
								log.Println("设备ID：" + slic[0] + "上线")
							} else {
								log.Println("设备ID：" + slic[0] + "离线")
							}
						} else if slic[0] == "core.Status.out" {
							str = fmt.Sprintf("%v", v["String"])
							tools.DB.Where("court_id = ?", slic[1]).Find(&courtInfor)
							if str == "online" {
								Courtdate[courtInfor.CourtID] = "2"
							} else if str == "fault" {
								Courtdate[courtInfor.CourtID] = "1"
							} else {
								Courtdate[courtInfor.CourtID] = "0"
								if courtInfor.CourtState != "1" {
								}
							}
						}
					}
				}
			}()
			time.Sleep(1 * time.Second)
		}
	}

	defer v.Conn.Close()
}

// 用于设备定位
func DevicePosition(courtid string) Returnmsg {
	fmt.Println("收到" + courtid + "定位请求")
	v := tools.NetTcp()
	var msg Returnmsg
	type devicePositionsub struct {
		Name     string
		Controls []map[string]interface{}
	}
	ds := []map[string]interface{}{}
	ds = append(ds, map[string]interface{}{"Name": "coreIdentify", "Value": 1})
	if len(courtid) == 1 {
		courtid = "00" + courtid
	} else if len(courtid) == 2 {
		courtid = "0" + courtid
	}
	d := Header{Jsonrpc: "2.0", Id: 1234, Method: "Component.Set", Params: devicePositionsub{
		Name: "web_remoteMonitor_" + courtid, Controls: ds}}
	fmt.Println(d)
	jsona, _ := json.Marshal(d)
	var x [1024]byte
	fmt.Println(string(jsona))
	readSize, _ := v.Conn.Read(x[0:])
	json.Unmarshal(x[0:readSize], &msg)
	if msg.Result {
		fmt.Println("打开定位成功")
	} else {
		fmt.Println("打开定位失败")
	}
	return msg
}

// 实现远程监听功能,要传通道，设备ID，还有开启关闭
func RemoteListen(line string, deviceid string, comm string) {
	if len(deviceid) == 1 {
		deviceid = "00" + deviceid
	} else if len(deviceid) == 2 {
		deviceid = "0" + deviceid
	}
	v := tools.NetTcp()
	var msg Returnmsg
	var x [1024]byte
	command := []map[string]interface{}{}
	type devicePositionsub struct {
		Name     string
		Controls []map[string]interface{}
	}
	if comm == "open" {
		h := line
		command = append(command, map[string]interface{}{"Name": "channelSelect", "Value": h})
		fmt.Println(11111, command)
		c := Header{Jsonrpc: "2.0", Id: 1234, Method: "Component.Set", Params: devicePositionsub{
			Name: "web_remoteMonitor_" + deviceid, Controls: command}}
		jsona, _ := json.Marshal(c)
		v.Conn.Write([]byte(string(jsona) + "\x00"))
		readSize, _ := v.Conn.Read(x[0:])
		json.Unmarshal(x[0:readSize], &msg)
		if msg.Result {
			fmt.Println("通道切换成功")
			// 用于通道静音开启
			command = append(command[:0], command[1:]...)
			command = append(command, map[string]interface{}{"Name": "mute", "Value": false})
			d := Header{Jsonrpc: "2.0", Id: 1234, Method: "Component.Set", Params: devicePositionsub{
				Name: "web_remoteMonitor_" + deviceid, Controls: command}}
			jsona, _ = json.Marshal(d)
			v.Conn.Write([]byte(string(jsona) + "\x00"))
			readSize, _ = v.Conn.Read(x[0:])
			json.Unmarshal(x[0:readSize], &msg)
			if msg.Result {
				fmt.Println("打开静音成功")
			} else {
				fmt.Println("打开静音失败")
			}
		} else {
			fmt.Println("通道切换失败，无法实现监听")
		}
	} else {
		// 用于通道静音开启
		command = append(command, map[string]interface{}{"Name": "mute", "Value": true})
		d := Header{Jsonrpc: "2.0", Id: 1234, Method: "Component.Set", Params: devicePositionsub{
			Name: "web_remoteMonitor_" + deviceid, Controls: command}}
		jsona, _ := json.Marshal(d)
		v.Conn.Write([]byte(string(jsona) + "\x00"))
		readSize, _ := v.Conn.Read(x[0:])
		json.Unmarshal(x[0:readSize], &msg)
		if msg.Result {
			fmt.Println("关闭静音成功")
		} else {
			fmt.Println("关闭静音失败")
		}
	}

}

// 用于检测法院状态，故障侦测
func MonitorPageState() {
	for {
		//未修复法院ID
		courtid := make(map[string]int)
		//得到未修复日志
		var faillogs []tools.FyFaillogs
		tools.DB.Where("repair = ?", "0").Pluck("courtid", &faillogs)
		//得到未修复法院ID
		for _, v := range faillogs {
			courtid[v.Courtid] = 0
		}

		//通过未修复日志更新法院状态，，未修复直接赋值故障也就是1
		var temp tools.FyCourtInfor
		//查找故障法院
		var fail []tools.FyCourtInfor
		tools.DB.Where("court_state = ?", "1").Pluck("court_id", &fail)
		fmt.Println(23232323, courtid)
		for i, v := range Courtdate {
			if len(courtid) != 0 {
				for h, _ := range courtid {
					if i == h {
						fmt.Println(i, h)
						tools.DB.Model(&temp).Where("court_id = ?", i).Update("court_state", "1")
					}
				}
			} else {
				tools.DB.Model(&temp).Where("court_id = ?", i).Update("court_state", v)
			}
		}
		/*for _, v := range fail {
			fmt.Println(1111, v)
			for i, _ := range courtid {
				fmt.Println(9999, i, v)
				tools.DB.Model(&temp).Where("court_id = ?", i).Update("court_state", "1")
				if _, ok := courtid[v.CourtID]; !ok {
					if _, ok := qsc.QscData[v.CourtID]; ok {
						fmt.Println(232323, qsc.QscData)
						tools.DB.Model(&temp).Where("court_id = ?", v.CourtID).Update("court_state", qsc.QscData[v.CourtID])
					}
				}
			}
		}*/
		time.Sleep(500 * time.Millisecond)
	}
}

//获取电频,传设备号、通道号
/*func ChannelLevel(deviceid string, line []tools.FyDeviceChannel, s socketio.Conn) {
	type l struct {
		Id   string
		Rate float32
	}
	var v net.Conn
	for i := 0; i < len(tools.Pools); i++ {
		if s.ID() == tools.Pools[i].Id {
			v = tools.Pools[i].Conn
		}
	}
	heartbeatdata := heartbeat{Jsonrpc: "2.0", Method: "NoOp", Params: map[string]string{}}
	go func() {
		for {
			jsonheart, _ := json.Marshal(heartbeatdata)
			v.Write([]byte(string(jsonheart) + "\x00"))
			time.Sleep(50 * time.Second)
		}
	}()
	var msg returnmsg
	var x [1024]byte
	var globaldata recvglobaldata
	var qscLevelData = make(map[string]string)
	command := []map[string]interface{}{}
	type z struct {
		Id        string
		Component interface{}
	}
	type u struct {
		Name     string
		Controls []map[string]interface{}
	}
	for _, i := range line {
		command = append(command, map[string]interface{}{"Name": i.ChannelCode})
	}
	fmt.Println(222, deviceid)
	if len(deviceid) == 1 {
		deviceid = "00" + deviceid
	} else if len(deviceid) == 2 {
		deviceid = "0" + deviceid
	}
	c := Header{Jsonrpc: "2.0", Id: 1234, Method: "ChangeGroup.AddComponentControl", Params: z{
		Id: "changegroup_3", Component: u{Name: "web_remoteMonitor_" + deviceid, Controls: command}}}
	jsona, _ := json.Marshal(c)
	v.Write([]byte(string(jsona) + "\x00"))
	readSize, _ := v.Read(x[0:])
	json.Unmarshal(x[0:readSize], &msg)
	if msg.Result {
		fmt.Println("通道数据设定成功")
		b := Header{Jsonrpc: "2.0", Id: 1234, Method: "ChangeGroup.AutoPoll", Params: l{Id: "changegroup_3", Rate: 0.3}}
		jsona, _ = json.Marshal(b)
		v.Write([]byte(string(jsona) + "\x00"))
		readSize, _ := v.Read(x[0:])
		json.Unmarshal(x[0:readSize], &msg)
		fmt.Println(msg)
		if msg.Result {
			fmt.Println("订阅成功")
			m := 1
			for m != 91 {
				readSize, _ = v.Read(x[0:])
				json.Unmarshal(x[0:readSize], &globaldata)
				fmt.Println(111111, readSize)
				m = readSize
				if len(globaldata.Params.Changes) != 0 {
					for _, i := range globaldata.Params.Changes {
						name := fmt.Sprintf("%v", i["Name"])
						value := fmt.Sprintf("%v", i["Value"])
						qscLevelData[name] = value
						s.Emit("recvYcjtlevel", qscLevelData)
					}
				}
				time.Sleep(300 * time.Millisecond)
			}
		} else {
			fmt.Println("订阅失败")
		}

	} else {
		fmt.Println("通道数据设定不成功")
	}
	defer v.Close()
}*/

// 关闭订阅,传相应params": {"Id": "changegroup_3
func CloseFeeds(context string, s socketio.Conn) {
	var v net.Conn
	for i := 0; i < len(tools.Pools); i++ {
		if s.ID() == tools.Pools[i].Id {
			v = tools.Pools[i].Conn
		}
	}
	var msg Returnmsg
	var x [1024]byte
	c := Header{Jsonrpc: "2.0", Id: 1234, Method: "ChangeGroup.Destroy", Params: map[string]string{"Id": context}}
	jsona, _ := json.Marshal(c)
	v.Write([]byte(string(jsona) + "\x00"))
	readSize, _ := v.Read(x[0:])
	json.Unmarshal(x[0:readSize], &msg)
	if msg.Result {
		fmt.Println("取消订阅成功")
	} else {
		fmt.Println("取消订阅失败")
	}
	defer v.Close()
}

// 获取设备信息
func DeviceInfor() {
	type s struct {
		Id   string
		Rate float32
	}
	v := tools.NetTcp()
	var msg Returnmsg
	var x [1024]byte
	var globaldata recvglobaldata
	command := []map[string]interface{}{}
	type z struct {
		Id        string
		Component interface{}
	}
	type u struct {
		Name     string
		Controls []map[string]interface{}
	}
	command = append(command, map[string]interface{}{"Name": "chassisTemp"})
	command = append(command, map[string]interface{}{"Name": "processorTemp"})
	command = append(command, map[string]interface{}{"Name": "runningTime"})
	command = append(command, map[string]interface{}{"Name": "coreIdentifyReport"})

	c := Header{Jsonrpc: "2.0", Id: 1234, Method: "ChangeGroup.AddComponentControl", Params: z{
		Id: "changegroup_3", Component: u{Name: "web_remoteMonitor_001", Controls: command}}}
	jsona, _ := json.Marshal(c)
	v.Conn.Write([]byte(string(jsona) + "\x00"))
	readSize, _ := v.Conn.Read(x[0:])
	json.Unmarshal(x[0:readSize], &msg)
	if msg.Result {
		fmt.Println("通道数据设定成功")
		b := Header{Jsonrpc: "2.0", Id: 1234, Method: "ChangeGroup.AutoPoll", Params: s{Id: "changegroup_3", Rate: 1}}
		jsona, _ = json.Marshal(b)
		v.Conn.Write([]byte(string(jsona) + "\x00"))
		readSize, _ := v.Conn.Read(x[0:])
		json.Unmarshal(x[0:readSize], &msg)
		if msg.Result {
			fmt.Println("订阅成功")
			for {
				readSize, _ = v.Conn.Read(x[0:])
				json.Unmarshal(x[0:readSize], &globaldata)
				if len(globaldata.Params.Changes) != 0 {
					for _, i := range globaldata.Params.Changes {
						name := fmt.Sprintf("%v", i["Name"])
						value := fmt.Sprintf("%v", i["Value"])
						QscData[name] = value
					}
				}
				time.Sleep(300 * time.Millisecond)
			}
		} else {
			fmt.Println("订阅失败")
		}

	} else {
		fmt.Println("通道数据设定不成功")
	}
	defer v.Conn.Close()
}

func Run() {
	go GetGlobalData()
	go MonitorPageState()
	//DevicePosition()
	//RemoteListen(true)
	//ChannelLevel()
	//CloseFeeds()
	//DeviceInfor()
}
