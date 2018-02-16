package room

import (
	"fmt"

	"github.com/bigBarrage/roomManager/banned/utils"
	"github.com/bigBarrage/roomManager/system"
)

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
		case system.NODE_MESSAGE_TYPE_CLEAN_ROOM:
			cleanRoom(roomInfo)
		case system.NODE_MESSAGE_TYPE_RELOAD_BANNED_IP:
			utils.LoadIpList()
		case system.NODE_MESSAGE_TYPE_RELOAD_BANNED_USER:
			utils.LoadUserList()
		case system.NODE_MESSAGE_TYPE_RELOAD_BANNED_WORDS:
			utils.LoadWordList()
		default:
			fmt.Println("main revicer获得消息")
			sendMessage(roomInfo, &s)
		}
	}
}
