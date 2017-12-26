package room

import "github.com/bigBarrage/roomManager/system"

func sendMessageToRoom(roomInfo *RoomInfo, message interface{}) {
	roomInfo.Lock.RLock()
	defer roomInfo.Lock.RUnlock()
	for _, rows := range roomInfo.Rows {
		go func(r *RowList) {
			rows.RowLock.Lock()
			defer rows.RowLock.Unlock()
			for _, node := range rows.Nodes {
				if node.RoomID == roomInfo.RoomID && node.IsAlive {
					node.SendMessage(message)
				}
			}
		}(rows)
	}
}

func sendMessageToBroadcastingStation(roomInfo *RoomInfo, message interface{}) {

}

func sendMessage(roomInfo *RoomInfo, nm *system.NodeMessage) {
	if nm.MessageTarget == system.MESSAGE_TARGET_ROOM {
		sendMessageToRoom(roomInfo, nm.Body)
	} else if nm.MessageTarget == system.MESSAGE_TARGET_BROADCASTINGSTATION {
		sendMessageToBroadcastingStation(roomInfo, nm.Body)
	}
}
