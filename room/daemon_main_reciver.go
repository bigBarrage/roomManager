package room

import "github.com/bigBarrage/roomManager/system"

//守护协程，用于一直从通道里面获取消息
func daemonMainReciver(c chan system.NodeMessage, roomInfo *RoomInfo) {
	for s := range c {
		switch s.MessageType {
		case system.NODE_MESSAGE_TYPE_IN_HALL:
			changeRoom(roomInfo, s.Body.(*ClientNode))
		case system.NODE_MESSAGE_TYPE_CLOSE_ROOM:
			CloseRoom(roomInfo.RoomID)
			break
		case system.NODE_MESSAGE_TYPE_CHANGE_ROOM:
			changeRoom(roomInfo, s.Body.(*ClientNode))
		case system.NODE_MESSAGE_TYPE_SEND_MESSAGE:
			sendMessage(roomInfo, s.Body)
		case system.NODE_MESSAGE_TYPE_CLEAN_ROOM:
			cleanRoom(roomInfo)
		}
	}
}
