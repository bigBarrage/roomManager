package room

/*
新的想法是：由外界对房间进行发送整理和重载黑名单的触发，不再进行内部维护

import (
	"time"

	"github.com/bigBarrage/roomManager/config"
	"github.com/bigBarrage/roomManager/system"
)

func init() {
	go CleanRoomCircle()
}

func CleanRoomCircle() {
	for {
		messageChannelLock.RLock()
		for _, c := range messageChannel {
			nm := system.NodeMessage{
				MessageType:   system.NODE_MESSAGE_TYPE_CLEAN_ROOM,
				MessageTarget: system.MESSAGE_TARGET_ROOM,
			}
			c <- nm
		}
		messageChannelLock.RUnlock()
		time.Sleep(config.RoomCleanTimer)
	}
}
*/
