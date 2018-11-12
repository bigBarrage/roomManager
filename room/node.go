package room

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bigBarrage/roomManager/config"
	"github.com/bigBarrage/roomManager/system"
	"github.com/gorilla/websocket"
)

//客户端节点
type ClientNode struct {
	RoomID       string                   //房间ID
	IP           string                   //当前IP地址
	UserID       string                   //用户标识
	DisableRead  bool                     //是否停止接收该链接内容
	SendChannel  *chan system.NodeMessage //发送消息使用的消息通道
	Conn         *websocket.Conn          //websocket链接
	UpdateTime   time.Time                //最后一次整理时间
	LastSendTime time.Time                //最后一次发送消息时间
	IsAlive      bool                     //是否存活
}

//默认加入大厅
func (this *ClientNode) Add() {
	this.IsAlive = true
	this.RoomID = ""

	messageChannelLock.RLock()
	if c, ok := messageChannel[this.RoomID]; ok {
		this.SendChannel = &c

		nm := system.NodeMessage{
			MessageType:   system.NODE_MESSAGE_TYPE_IN_HALL,
			MessageTarget: system.MESSAGE_TARGET_ROOM,
			Body:          this,
		}
		sendMessageToChannel(this, nm)
	}
	messageChannelLock.RUnlock()
}

//修改房间
func (this *ClientNode) ChangeRoom(roomID string) {
	if this.IsAlive == false || this.RoomID == roomID {
		return
	}
	this.RoomID = roomID

	messageChannelLock.RLock()
	if c, ok := messageChannel[this.RoomID]; ok {
		this.SendChannel = &c

		nm := system.NodeMessage{
			MessageType:   system.NODE_MESSAGE_TYPE_IN_HALL,
			MessageTarget: system.MESSAGE_TARGET_ROOM,
			Body:          this,
		}
		sendMessageToChannel(this, nm)
		fmt.Println("更换房间成功")
	} else {
		fmt.Println("更换房间失败")
		this.SendMessage(system.NodeMessage{MessageType: system.NODE_MESSAGE_TYPE_ROOM_NOT_EXISTS})
	}
	messageChannelLock.RUnlock()
}

//发送消息
func (this *ClientNode) SendMessageToRoom(message interface{}) {
	if this.IsAlive == false || this.DisableRead || this.RoomID == "" || this.UserID == "" {
		this.SendMessage(system.ErrorUnknownMessage)
		return
	}

	nm := system.NodeMessage{
		MessageType:   system.NODE_MESSAGE_TYPE_SEND_MESSAGE,
		MessageTarget: system.MESSAGE_TARGET_BROADCASTINGSTATION,
		Body:          message,
	}
	if config.UseBoradcasting == true {
		nm.MessageTarget = system.MESSAGE_TARGET_BROADCASTINGSTATION
	} else {
		nm.MessageTarget = system.MESSAGE_TARGET_ROOM
	}

	rs := sendMessageToChannel(this, nm)
	fmt.Println("发送结果：", rs)
}

//对本节点发送消息
func (this *ClientNode) SendMessage(message interface{}) {
	w, err := this.Conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return
	}
	defer w.Close()
	msg, err := json.Marshal(message)
	if err != nil {
		return
	}
	w.Write(msg)
	this.LastSendTime = time.Now()
}

//更换用户ID
func (this *ClientNode) ChangeUserID(userID string) {
	this.UserID = userID
}

//关闭连接
func (this *ClientNode) Close() {
	this.IsAlive = false
}
