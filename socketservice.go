package main

import (
	"github.com/gin-gonic/gin"
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
	"server/methods"
	"server/tools"
)

// 广播方法
//server.BroadcastToRoom("/", "chat", "reply1", []string{"xjg", "xjg1"})
//socketio.Broadcast.SendAll()

// 用于SOCKET io方法

func SocketIo() {
	router := gin.New()
	server := socketio.NewServer(nil)
	// 用于连接提示
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID(), s.RemoteAddr())
		return nil
	})
	// 用于用户信息接收以及验证
	server.OnEvent("/", "userauth", methods.Userauth)
	// 用于接收远程监听的指令
	server.OnEvent("/", "getycjtinit", methods.Getycjt)
	// 用于获取菜单项
	server.OnEvent("/", "menulist", methods.MenuList)
	// 获取设备管理内的URL
	server.OnEvent("/", "geturl", methods.Geturl)
	// 用于系统日志记录接口
	server.OnEvent("/", "Operatelogs", methods.Operatelogs)
	// 用于故障日志记录接口
	server.OnEvent("/", "getfaillogs", methods.Getfaillogs)
	// 用于系统日志记录接口
	server.OnEvent("/", "getoperatelogs", methods.Getoperatelogs)
	// 用于监测页面初始数据
	server.OnEvent("/", "monitorinit", methods.Monitorinit)
	// 用于监测页面的故障日志添加接口
	server.OnEvent("/", "monitorerrinput", methods.Monitorerrinput)
	// 用于获取法院状态
	server.OnEvent("/", "getcourtstate", methods.Getcourtstate)
	// 获取故障设备
	server.OnEvent("/", "getfaildevice", methods.Getfaildevice)
	// 用于打开远程监听
	server.OnEvent("/", "listenopen", methods.Listenopen)
	// 故障处理页面信息方法
	server.OnEvent("/", "failrepair", methods.Failrepair)
	// 故障处理方法信息提交
	server.OnEvent("/", "failinforsumit", methods.Failinforsumit)
	// 关闭远程监听订阅方法
	server.OnEvent("/", "closeycjtbtn", methods.Closeycjtbtn)
	// 主页获取统计数据
	server.OnEvent("/", "getindexcourt", methods.Getindexcourt)
	// 用于获取下载文件列表
	server.OnEvent("/", "downfileinit", methods.Downfileinit)
	// 创建至QSC的TCP
	server.OnEvent("/", "tcpconn", methods.Tcpconn)
	// 用于设备定位
	server.OnEvent("/", "devicelocation", methods.Devicelocation)
	// 关闭至QSC的TCP
	server.OnEvent("/", "closeconn", methods.Closeconn)

	// 用于错误
	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})
	//用于断开连接提示
	server.OnDisconnect("/", func(s socketio.Conn, msg string) {
		log.Println("closed", msg)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()
	router.Use(tools.Cros())
	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))
	router.StaticFS("/public", http.Dir("../asset"))
	if err := router.Run(":8000"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
