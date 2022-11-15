package tools

import (
	"net"
	"time"
)

type Qnet struct {
	Id   string
	Conn net.Conn
}

var Pools []Qnet

func NetTcp() Qnet {
	//192.168.100.200
	conn, _ := net.Dial("tcp", "10.0.0.202:1701")
	var x [1024]byte
	_, _ = conn.Read(x[0:])
	return Qnet{Conn: conn}
}

func (c Qnet) Send(id string) {
	c.Conn.Write([]byte(id))
	time.Sleep(200 * time.Millisecond)
}
