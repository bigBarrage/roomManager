package room

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
