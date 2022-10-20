package tools

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

// FyCourtInfor 用于法院信息数据构建
type FyCourtInfor struct {
	ID         uint   `gorm:"primarykey"`
	CourtID    string //法院ID与设备表中的法院ID相关联
	CourtZone  string //所属区域与FYZoneInfor的ZoneID关联
	CourtName  string //法院中文名称
	CourtState string //法院状态，0代表红离线，1代表蓝色故障，2代表绿在线
	CourtSort  int    //区域排序
}

// 设备表
type FyDeviceInfor struct {
	ID          uint   `gorm:"primarykey"`
	CourtID     string //法院设备ID
	Deviceid    string //设备ID
	DeviceName  string //设备名称
	DeviceIp    string //设备ip
	ControlUrl  string //控制URL
	DeviceState string //设备状态 0代表故障，1代表正常
}

// 设备通道表
type FyDeviceChannel struct {
	ID          uint   `gorm:"primarykey"`
	Deviceid    string //设备ID
	ChannelName string //通道名称
	ChannelCode string //处理器中通道名称
}

// 法院文档
type FyDownFile struct {
	ID           uint   `gorm:"primarykey"`
	Courtid      string //法院ID
	Filename     string //文件真实名称
	Filesavename string //文件保存名
	Filetext     string //文件描述
	Filetype     string //文件类型
}

// 用于区域法院
type FyZoneInfor struct {
	ID       uint   `gorm:"primarykey"`
	ZoneID   string //与fyCourtInfor里的CourtZone关联
	ZoneName string //区域中文名称
	ZoneSort int    //区域排序
}

// 用于导航菜单及其权限，0代表只管理员访问，1代表区域可访问，2代表所有用户可访问
type FyNavigateMenu struct {
	ID             uint   `gorm:"primarykey"`
	MenuName       string //菜单名
	MenuLink       string //菜单链接
	MenuPrivileges int    //菜单权限
	MenuSort       int    //菜单排序
}

// 用户表，权限0代表管理员1代表区域2代表用户
type FyUsers struct {
	ID         uint   `gorm:"primarykey"`
	UserId     int    //用户ID
	UserName   string //用户名
	UserPasswd string //密码
	Privileges int    //权限
	CourtID    string //所属法院与法院信息表相关联
	ZoneID     string //存储区域ID与区域名关联
}

// 用于系统日志记录
type FySystemlogs struct {
	ID       uint   `gorm:"primarykey"`
	Time     string //时间
	Address  string //访问地址
	UserName string //用户名
	Content  string //日志内容
}

// 用于故障日志
type FyFaillogs struct {
	ID             uint   `gorm:"primarykey"`
	Failid         string //故障ID
	Reporttime     string //上报时间
	Processtime    string //处理时间
	Courtid        string //故障节点，用法院ID联连上表
	Reportuser     string //上报用户
	Processuser    string //处理用户
	Failcontent    string //故障内容
	Processcontent string //处理方法
	Deviceid       string //设备ID
	Repair         string //修复状态 0代表代修复，1代表完成
}

// 实现数据库的连接
func DatabaseConn() *gorm.DB {
	dsn := "root:iloveyou@tcp(127.0.0.1:3306)/audiosystem?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(50)  //空闲连接数
	sqlDB.SetMaxOpenConns(300) //最大连接数
	sqlDB.SetConnMaxLifetime(time.Minute)
	if err != nil {
		log.Println("数据库连接失败")
		panic("数据库连接失败")
	}
	return db
}

// 实现数据库的初始建表工作
func DatabaseInit() *gorm.DB {
	dsn := "root:iloveyou@tcp(127.0.0.1:3306)/audiosystem?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("数据库连接失败")
		panic("数据库连接失败")
	}
	db.AutoMigrate(&FyCourtInfor{}, &FyZoneInfor{}, &FyNavigateMenu{}, &FyUsers{}, &FyDeviceInfor{}, &FySystemlogs{}, &FyFaillogs{}, &FyDeviceChannel{}, &FyDownFile{})
	return db
}
