package tools

func Systeminit() {
	databaseDataInit()
}
func databaseDataInit() {
	db := DatabaseInit()
	//用于创建法院资料文件内容
	db.Create(&FyDownFile{
		Courtid:      "1",
		Filesavename: "asdfasdfa232f2f23f2ff",
		Filetext:     "大法庭审系统图纸",
		Filetype:     "图纸",
		Filename:     "大法庭审系统图.jpg",
	})
	//用于创建设备通道信息表
	db.Create(&FyDeviceChannel{
		Deviceid:    "1",
		ChannelCode: "meter 1",
		ChannelName: "话筒1",
	})
	db.Create(&FyDeviceChannel{
		Deviceid:    "1",
		ChannelCode: "meter 2",
		ChannelName: "话筒2",
	})
	db.Create(&FyDeviceChannel{
		Deviceid:    "1",
		ChannelCode: "meter 3",
		ChannelName: "话筒3",
	})
	db.Create(&FyDeviceChannel{
		Deviceid:    "1",
		ChannelCode: "meter 4",
		ChannelName: "话筒4",
	})
	db.Create(&FyDeviceChannel{
		Deviceid:    "1",
		ChannelCode: "meter 5",
		ChannelName: "话筒5",
	})
	db.Create(&FyDeviceChannel{
		Deviceid:    "2",
		ChannelCode: "meter 1",
		ChannelName: "话筒1",
	})
	db.Create(&FyDeviceChannel{
		Deviceid:    "2",
		ChannelCode: "meter 2",
		ChannelName: "话筒2",
	})
	db.Create(&FyDeviceChannel{
		Deviceid:    "2",
		ChannelCode: "meter 3",
		ChannelName: "话筒3",
	})
	db.Create(&FyDeviceChannel{
		Deviceid:    "2",
		ChannelCode: "meter 4",
		ChannelName: "话筒4",
	})
	db.Create(&FyDeviceChannel{
		Deviceid:    "2",
		ChannelCode: "meter 5",
		ChannelName: "话筒5",
	})
	//用于设备表数据
	//https://10.0.0.20/uci-viewer/?uci=%E5%A4%A7%E6%B3%95%E5%BA%AD%E9%9F%B3%E9%A2%91%E6%8E%A7%E5%88%B6&file=2.UCI.xml&directory=/designs/current_design/UCIs/
	db.Create(&FyDeviceInfor{
		CourtID:    "1",
		DeviceName: "大法庭",
		DeviceIp:   "10.0.0.20",
		ControlUrl: "http://hao123.com",
		Deviceid:   "1",
	})
	db.Create(&FyDeviceInfor{
		CourtID:    "1",
		DeviceName: "大法庭1",
		DeviceIp:   "10.0.0.21",
		ControlUrl: "http://163.com",
		Deviceid:   "2",
	})
	//用于创建菜单连接功能
	db.Create(&FyNavigateMenu{
		MenuName:       "主页",
		MenuLink:       "/",
		MenuPrivileges: 3,
		MenuSort:       1,
	})
	db.Create(&FyNavigateMenu{
		MenuName:       "监测管理",
		MenuLink:       "/monitor",
		MenuPrivileges: 2,
		MenuSort:       2,
	})
	db.Create(&FyNavigateMenu{
		MenuName:       "设备管理",
		MenuLink:       "/devicecontroluser/0",
		MenuPrivileges: 3,
		MenuSort:       3,
	})
	db.Create(&FyNavigateMenu{
		MenuName:       "日志中心",
		MenuLink:       "/log",
		MenuPrivileges: 3,
		MenuSort:       4,
	})

	//用于登录用户及权限，权限分0，1，2，3在获取监测管管理数据时，0-1代表可全部获取，2代表获取自己区域，3代表不获取
	db.Create(&FyUsers{
		UserId:     1,
		UserName:   "admin",
		UserPasswd: "admin",
		Privileges: 0,
		CourtID:    "1",
		ZoneID:     "0",
	})
	db.Create(&FyUsers{
		UserId:     1,
		UserName:   "hz",
		UserPasswd: "hz",
		Privileges: 2,
		CourtID:    "1",
		ZoneID:     "1",
	})
	db.Create(&FyUsers{
		UserId:     1,
		UserName:   "user",
		UserPasswd: "user",
		Privileges: 3,
		CourtID:    "1",
		ZoneID:     "2",
	})
	//用于创建法院区域表
	db.Create(&FyZoneInfor{
		ZoneID:   "1",
		ZoneName: "杭州地区",
		ZoneSort: 1,
	})
	db.Create(&FyZoneInfor{
		ZoneID:   "2",
		ZoneName: "宁波地区",
		ZoneSort: 2,
	})
	db.Create(&FyZoneInfor{
		ZoneID:   "3",
		ZoneName: "温州地区",
		ZoneSort: 3,
	})
	db.Create(&FyZoneInfor{
		ZoneID:   "4",
		ZoneName: "嘉兴地区",
		ZoneSort: 4,
	})
	db.Create(&FyZoneInfor{
		ZoneID:   "5",
		ZoneName: "湖州地区",
		ZoneSort: 5,
	})
	db.Create(&FyZoneInfor{
		ZoneID:   "6",
		ZoneName: "绍兴地区",
		ZoneSort: 6,
	})
	db.Create(&FyZoneInfor{
		ZoneID:   "7",
		ZoneName: "金华地区",
		ZoneSort: 7,
	})
	db.Create(&FyZoneInfor{
		ZoneID:   "8",
		ZoneName: "衢州地区",
		ZoneSort: 8,
	})
	db.Create(&FyZoneInfor{
		ZoneID:   "9",
		ZoneName: "丽水地区",
		ZoneSort: 9,
	})
	db.Create(&FyZoneInfor{
		ZoneID:   "10",
		ZoneName: "台州地区",
		ZoneSort: 10,
	})
	db.Create(&FyZoneInfor{
		ZoneID:   "11",
		ZoneName: "舟山地区",
		ZoneSort: 11,
	})
	//用于创建法院信息表
	//杭州表
	db.Create(&FyCourtInfor{
		CourtID:    "1",
		CourtName:  "杭州中院",
		CourtSort:  1,
		CourtZone:  "1",
		CourtState: "1",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "2",
		CourtName:  "江干法院",
		CourtSort:  2,
		CourtZone:  "1",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "3",
		CourtName:  "拱墅法院",
		CourtSort:  3,
		CourtZone:  "1",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "4",
		CourtName:  "萧山法院",
		CourtSort:  4,
		CourtZone:  "1",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "5",
		CourtName:  "余杭法院",
		CourtSort:  5,
		CourtZone:  "1",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "6",
		CourtName:  "下城法院",
		CourtSort:  6,
		CourtZone:  "1",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "7",
		CourtName:  "富阳法院",
		CourtSort:  7,
		CourtZone:  "1",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "8",
		CourtName:  "淳安法院",
		CourtSort:  8,
		CourtZone:  "1",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "9",
		CourtName:  "建德法院",
		CourtSort:  9,
		CourtZone:  "1",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "10",
		CourtName:  "下沙经济开发区",
		CourtSort:  10,
		CourtZone:  "1",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "11",
		CourtName:  "西湖法院",
		CourtSort:  11,
		CourtZone:  "1",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "12",
		CourtName:  "滨江法院",
		CourtSort:  12,
		CourtZone:  "1",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "13",
		CourtName:  "上城法院",
		CourtSort:  13,
		CourtZone:  "1",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "14",
		CourtName:  "桐庐法院",
		CourtSort:  14,
		CourtZone:  "1",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "15",
		CourtName:  "临安法院",
		CourtSort:  15,
		CourtZone:  "1",
		CourtState: "0",
	})
	//宁波表
	db.Create(&FyCourtInfor{
		CourtID:    "16",
		CourtName:  "宁波中院",
		CourtSort:  1,
		CourtZone:  "2",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "17",
		CourtName:  "江东法院",
		CourtSort:  2,
		CourtZone:  "2",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "18",
		CourtName:  "海事法院",
		CourtSort:  3,
		CourtZone:  "2",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "19",
		CourtName:  "慈溪法院",
		CourtSort:  4,
		CourtZone:  "2",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "20",
		CourtName:  "余姚法院",
		CourtSort:  5,
		CourtZone:  "2",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "21",
		CourtName:  "奉化法院",
		CourtSort:  6,
		CourtZone:  "2",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "22",
		CourtName:  "海曙法院",
		CourtSort:  7,
		CourtZone:  "2",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "23",
		CourtName:  "鄞州法院",
		CourtSort:  8,
		CourtZone:  "2",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "24",
		CourtName:  "宁海法院",
		CourtSort:  9,
		CourtZone:  "2",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "25",
		CourtName:  "镇海法院",
		CourtSort:  10,
		CourtZone:  "2",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "26",
		CourtName:  "江北法院",
		CourtSort:  11,
		CourtZone:  "2",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "27",
		CourtName:  "北仑法院",
		CourtSort:  12,
		CourtZone:  "2",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "28",
		CourtName:  "象山法院",
		CourtSort:  13,
		CourtZone:  "2",
		CourtState: "0",
	})
	//温州地区
	db.Create(&FyCourtInfor{
		CourtID:    "29",
		CourtName:  "温州中院",
		CourtSort:  30,
		CourtZone:  "3",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "31",
		CourtName:  "乐清法院",
		CourtSort:  2,
		CourtZone:  "3",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "32",
		CourtName:  "永嘉法院",
		CourtSort:  3,
		CourtZone:  "3",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "33",
		CourtName:  "鹿城法院",
		CourtSort:  4,
		CourtZone:  "3",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "34",
		CourtName:  "洞头法院",
		CourtSort:  5,
		CourtZone:  "3",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "35",
		CourtName:  "瓯海法院",
		CourtSort:  6,
		CourtZone:  "3",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "36",
		CourtName:  "瑞安法院",
		CourtSort:  7,
		CourtZone:  "3",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "37",
		CourtName:  "泰顺法院",
		CourtSort:  8,
		CourtZone:  "3",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "38",
		CourtName:  "苍南法院",
		CourtSort:  9,
		CourtZone:  "3",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "39",
		CourtName:  "文城法院",
		CourtSort:  10,
		CourtZone:  "3",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "40",
		CourtName:  "龙湾法院",
		CourtSort:  11,
		CourtZone:  "3",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "41",
		CourtName:  "平阳法院",
		CourtSort:  12,
		CourtZone:  "3",
		CourtState: "0",
	})
	//嘉兴地区
	db.Create(&FyCourtInfor{
		CourtID:    "42",
		CourtName:  "嘉兴中院",
		CourtSort:  1,
		CourtZone:  "4",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "43",
		CourtName:  "秀洲法院",
		CourtSort:  2,
		CourtZone:  "4",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "44",
		CourtName:  "南湖法院",
		CourtSort:  3,
		CourtZone:  "4",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "45",
		CourtName:  "嘉善法院",
		CourtSort:  4,
		CourtZone:  "4",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "46",
		CourtName:  "海宁法院",
		CourtSort:  5,
		CourtZone:  "4",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "47",
		CourtName:  "桐乡法院",
		CourtSort:  6,
		CourtZone:  "4",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "48",
		CourtName:  "海盐法院",
		CourtSort:  7,
		CourtZone:  "4",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "49",
		CourtName:  "平湖法院",
		CourtSort:  8,
		CourtZone:  "4",
		CourtState: "0",
	})
	//湖州地区
	db.Create(&FyCourtInfor{
		CourtID:    "50",
		CourtName:  "湖州中院",
		CourtSort:  1,
		CourtZone:  "5",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "51",
		CourtName:  "吴兴法院",
		CourtSort:  2,
		CourtZone:  "5",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "52",
		CourtName:  "南浔法院",
		CourtSort:  3,
		CourtZone:  "5",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "53",
		CourtName:  "长兴法院",
		CourtSort:  4,
		CourtZone:  "5",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "54",
		CourtName:  "安吉法院",
		CourtSort:  5,
		CourtZone:  "5",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "55",
		CourtName:  "德清法院",
		CourtSort:  6,
		CourtZone:  "5",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "56",
		CourtName:  "南太湖法院",
		CourtSort:  7,
		CourtZone:  "5",
		CourtState: "0",
	})
	//绍兴地区
	db.Create(&FyCourtInfor{
		CourtID:    "57",
		CourtName:  "绍兴中院",
		CourtSort:  1,
		CourtZone:  "6",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "58",
		CourtName:  "上虞法院",
		CourtSort:  2,
		CourtZone:  "6",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "59",
		CourtName:  "越城法院",
		CourtSort:  3,
		CourtZone:  "6",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "60",
		CourtName:  "诸暨法院",
		CourtSort:  4,
		CourtZone:  "6",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "61",
		CourtName:  "柯桥法院",
		CourtSort:  5,
		CourtZone:  "6",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "62",
		CourtName:  "嵊州法院",
		CourtSort:  6,
		CourtZone:  "6",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "63",
		CourtName:  "新昌法院",
		CourtSort:  7,
		CourtZone:  "6",
		CourtState: "0",
	})
	//金华地区
	db.Create(&FyCourtInfor{
		CourtID:    "64",
		CourtName:  "金华中院",
		CourtSort:  1,
		CourtZone:  "7",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "65",
		CourtName:  "兰溪法院",
		CourtSort:  2,
		CourtZone:  "7",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "66",
		CourtName:  "浦江法院",
		CourtSort:  3,
		CourtZone:  "7",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "67",
		CourtName:  "金东法院",
		CourtSort:  4,
		CourtZone:  "7",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "68",
		CourtName:  "东阳法院",
		CourtSort:  5,
		CourtZone:  "7",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "69",
		CourtName:  "义乌法院",
		CourtSort:  6,
		CourtZone:  "7",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "70",
		CourtName:  "婺城法院",
		CourtSort:  7,
		CourtZone:  "7",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "71",
		CourtName:  "武义法院",
		CourtSort:  8,
		CourtZone:  "7",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "72",
		CourtName:  "磐安法院",
		CourtSort:  9,
		CourtZone:  "7",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "73",
		CourtName:  "永康法院",
		CourtSort:  10,
		CourtZone:  "7",
		CourtState: "0",
	})
	//衢州地区
	db.Create(&FyCourtInfor{
		CourtID:    "74",
		CourtName:  "衢州中院",
		CourtSort:  1,
		CourtZone:  "8",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "75",
		CourtName:  "衢江法院",
		CourtSort:  2,
		CourtZone:  "8",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "76",
		CourtName:  "柯城法院",
		CourtSort:  3,
		CourtZone:  "8",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "77",
		CourtName:  "江山法院",
		CourtSort:  4,
		CourtZone:  "8",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "78",
		CourtName:  "龙游法院",
		CourtSort:  5,
		CourtZone:  "8",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "79",
		CourtName:  "常山法院",
		CourtSort:  6,
		CourtZone:  "8",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "80",
		CourtName:  "开化法院",
		CourtSort:  7,
		CourtZone:  "8",
		CourtState: "0",
	})
	//丽水地区
	db.Create(&FyCourtInfor{
		CourtID:    "81",
		CourtName:  "丽水中院",
		CourtSort:  1,
		CourtZone:  "9",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "82",
		CourtName:  "莲都法院",
		CourtSort:  2,
		CourtZone:  "9",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "83",
		CourtName:  "青田法院",
		CourtSort:  3,
		CourtZone:  "9",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "84",
		CourtName:  "龙泉法院",
		CourtSort:  4,
		CourtZone:  "9",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "85",
		CourtName:  "松阳法院",
		CourtSort:  5,
		CourtZone:  "9",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "86",
		CourtName:  "遂昌法院",
		CourtSort:  6,
		CourtZone:  "9",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "87",
		CourtName:  "庆元法院",
		CourtSort:  7,
		CourtZone:  "9",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "88",
		CourtName:  "云和法院",
		CourtSort:  8,
		CourtZone:  "9",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "89",
		CourtName:  "景宁法院",
		CourtSort:  9,
		CourtZone:  "9",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "90",
		CourtName:  "缙云法院",
		CourtSort:  10,
		CourtZone:  "9",
		CourtState: "0",
	})
	//台州地区
	db.Create(&FyCourtInfor{
		CourtID:    "91",
		CourtName:  "台州中院",
		CourtSort:  1,
		CourtZone:  "10",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "92",
		CourtName:  "椒江法院",
		CourtSort:  2,
		CourtZone:  "10",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "93",
		CourtName:  "黄岩法院",
		CourtSort:  3,
		CourtZone:  "10",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "94",
		CourtName:  "临海法院",
		CourtSort:  4,
		CourtZone:  "10",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "95",
		CourtName:  "玉环法院",
		CourtSort:  5,
		CourtZone:  "10",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "96",
		CourtName:  "仙居法院",
		CourtSort:  6,
		CourtZone:  "10",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "97",
		CourtName:  "路桥法院",
		CourtSort:  7,
		CourtZone:  "10",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "98",
		CourtName:  "三门法院",
		CourtSort:  8,
		CourtZone:  "10",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "99",
		CourtName:  "天台法院",
		CourtSort:  9,
		CourtZone:  "10",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "100",
		CourtName:  "温岭法院",
		CourtSort:  10,
		CourtZone:  "10",
		CourtState: "0",
	})
	//舟山地区
	db.Create(&FyCourtInfor{
		CourtID:    "101",
		CourtName:  "舟山中院",
		CourtSort:  1,
		CourtZone:  "11",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "102",
		CourtName:  "定海法院",
		CourtSort:  2,
		CourtZone:  "11",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "103",
		CourtName:  "普陀法院",
		CourtSort:  3,
		CourtZone:  "11",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "104",
		CourtName:  "岱山法院",
		CourtSort:  4,
		CourtZone:  "11",
		CourtState: "0",
	})
	db.Create(&FyCourtInfor{
		CourtID:    "105",
		CourtName:  "嵊泗法院",
		CourtSort:  5,
		CourtZone:  "11",
		CourtState: "0",
	})
}
