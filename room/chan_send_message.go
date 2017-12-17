package room

func sendMessageToRoom(roomInfo *RoomInfo, message interface{}) {
	roomInfo.Lock.RLock()
	defer roomInfo.Lock.RUnlock()
	for _, rows := range roomInfo.Rows {
		go func(rows) {
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
