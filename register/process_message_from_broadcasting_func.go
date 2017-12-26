package register

var ProcessMessageFromBroadcastingFunc func([]byte)

//注册从广播站接受到的消息处理方法
func RegisterProcessMessageFromBroadcastingFunc(pFunc func([]byte)) {
	ProcessMessageFromBroadcastingFunc = pFunc
}
