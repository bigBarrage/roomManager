package register

import "github.com/bigBarrage/roomManager/room"

var ProcessMessageFunc func([]byte, *room.ClientNode)

//注册处理从客户端收到的消息的方法
func RegisterProcessMessageFunc(pFunc func([]byte, *room.ClientNode)) {
	ProcessMessageFunc = pFunc
}
