package room

import (
	"sync"
	"time"

	json "github.com/json-iterator/go"

	"github.com/bigBarrage/roomManager/banned"
	"github.com/bigBarrage/roomManager/config"
	"github.com/bigBarrage/roomManager/logs"
	"github.com/bigBarrage/roomManager/system"
)

var messageChannel map[string]chan system.NodeMessage
var messageChannelLock = &sync.RWMutex{}

func init() {
	messageChannel = make(map[string]chan system.NodeMessage)
}

//给channel发送消息
func sendMessageToChannel(cn *ClientNode, nm system.NodeMessage) error {
	if cn.DisableRead {
		return CLIENT_IS_DISABLED_READ
	}
	if nm.MessageType == system.NODE_MESSAGE_TYPE_SEND_MESSAGE {
		n, err := json.Marshal(nm.Body)
		if err != nil {
			return err
		}
		if banned.IsBannedWords(string(n)) {
			return banned.WORDS_IS_BANNED
		}
	}
	if banned.IsBannedIP(cn.IP) {
		cn.DisableRead = true
		logs.DisplayLog(logs.DISPLAY_ERROR_LEVEL_NOTICE, "被限定的IP，关闭消息通道"+cn.IP)
		return banned.IP_IS_BANNED
	}
	if banned.IsBannedUserID(cn.UserID) {
		cn.DisableRead = true
		logs.DisplayLog(logs.DISPLAY_ERROR_LEVEL_NOTICE, "被限定的用户ID，关闭消息通道")
		return banned.UID_IS_BANNED
	}
	*cn.SendChannel <- nm
	return nil
}

//创建房间
func CreateRoom(roomID string) error {
	messageChannelLock.Lock()
	defer messageChannelLock.Unlock()
	c := make(chan system.NodeMessage, config.MaxMessageChannelLength)
	messageChannel[roomID] = c

	roomInfo := &RoomInfo{}
	roomInfo.Lock = &sync.RWMutex{}
	roomInfo.RoomID = roomID
	roomInfo.LastChangeTime = time.Now()
	//启动一系列worker
	go daemonMainReciver(c, roomInfo)
	return nil
}

//发送关闭房间的请求，准备关闭房间
func prepareToCloseRoom(roomID string) error {
	messageChannelLock.RLock()
	defer messageChannelLock.RUnlock()
	if c, ok := messageChannel[roomID]; ok {
		nm := system.NodeMessage{
			MessageType:   system.NODE_MESSAGE_TYPE_CLOSE_ROOM,
			MessageTarget: system.MESSAGE_TARGET_ROOM,
			Body:          nil,
		}
		c <- nm
	}
	return nil
}

//正式关闭房间
func CloseRoom(roomID string) error {
	prepareToCloseRoom(roomID)
	messageChannelLock.Lock()
	defer messageChannelLock.Unlock()
	if c, ok := messageChannel[roomID]; ok {
		delete(messageChannel, roomID)
		close(c)
	}
	return nil
}

//外部对房间发送消息
func SendMessageFromOuter(roomID string, nm system.NodeMessage) {
	if roomID == "" {
		return
	}
	if c, ok := messageChannel[roomID]; ok {
		c <- nm
	}
	return
}
