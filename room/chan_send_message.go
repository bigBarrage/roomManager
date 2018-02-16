package room

import (
	"encoding/json"
	"fmt"

	"github.com/bigBarrage/roomManager/system"
	"github.com/gorilla/websocket"
)

func sendMessageToRoom(roomInfo *RoomInfo, message interface{}) {
	roomInfo.Lock.RLock()
	defer roomInfo.Lock.RUnlock()
	fmt.Println("房间内开始发送")
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
	fmt.Println("房间内发送结束")
}

func sendMessageToBroadcastingStation(roomInfo *RoomInfo, message interface{}) {
	w, err := BroadcastingConnection.NextWriter(websocket.TextMessage)
	if err != nil {
		return
	}
	defer w.Close()
	msg, err := json.Marshal(message)
	if err != nil {
		return
	}
	w.Write(msg)
}

func sendMessage(roomInfo *RoomInfo, nm *system.NodeMessage) {
	if nm.MessageTarget == system.MESSAGE_TARGET_ROOM {
		fmt.Println("发送到房间")
		sendMessageToRoom(roomInfo, nm.Body)
	} else if nm.MessageTarget == system.MESSAGE_TARGET_BROADCASTINGSTATION {
		sendMessageToBroadcastingStation(roomInfo, nm.Body)
	}
}
