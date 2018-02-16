package register

import "github.com/bigBarrage/roomManager/system"

var ProcessMessageFunc func([]byte, system.MessageNode)

//注册处理从客户端收到的消息的方法
func RegisterProcessMessageFunc(pFunc func([]byte, system.MessageNode)) {
	ProcessMessageFunc = pFunc
}
