package run

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bigBarrage/roomManager/banned"
	"github.com/bigBarrage/roomManager/config"
	"github.com/bigBarrage/roomManager/logs"
	"github.com/bigBarrage/roomManager/register"
	"github.com/bigBarrage/roomManager/room"
	"github.com/bigBarrage/roomManager/system"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

//启动服务入口
func Run() error {
	if check() != nil {
		os.Exit(1)
	}
	room.ConnBroadcastingStation()
	http.HandleFunc("/"+config.ListenPath, handler)
	return http.ListenAndServe(":"+fmt.Sprint(config.ListenPort), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logs.DisplayLog(logs.DISPLAY_ERROR_LEVEL_NOTICE, "连接失败"+err.Error())
		return
	}

	node := &room.ClientNode{}

	fmt.Println("建立新连接")
	addrSlice := strings.Split(r.RemoteAddr, ":")
	node.IP = addrSlice[0]

	node.IsAlive = true
	node.Conn = conn
	node.Add()
	processMessage(node)
}

//做启动前检查，保证各项设置基本运行状况
func check() error {
	fmt.Fprintln(os.Stdout, "开始进行启动前检查：")
	if register.ProcessMessageFunc == nil {
		logs.DisplayLog(logs.DISPLAY_ERROR_LEVEL_ERROR, "消息处理方法未注册...")
		return errors.New("check failed")
	}

	if config.UseBoradcasting {
		if register.ProcessMessageFromBroadcastingFunc == nil {
			logs.DisplayLog(logs.DISPLAY_ERROR_LEVEL_ERROR, "检查广播站处理程序失败...")
			return errors.New("check failed")
		}
	}
	logs.DisplayLog(logs.DISPLAY_ERROR_LEVEL_NOTICE, "检查完毕")
	return nil
}

//处理消息
func processMessage(node *room.ClientNode) {
	for {
		mType, reader, err := node.Conn.NextReader()

		if mType == websocket.CloseMessage || mType == -1 {
			node.Close()
			return
		}

		if err != nil {
			return
		}

		if node.DisableRead {
			node.SendMessage(system.ErrorReadDisabledMessage)
			continue
		}

		if banned.IsBannedUserID(node.UserID) {
			node.DisableRead = true
			node.SendMessage(system.ErrorReadDisabledMessage)
			continue
		}

		msg := make([]byte, 0, config.MessageReadBufferLength)
		for {
			tmp := make([]byte, config.MessageReadBufferLength)
			l, err := reader.Read(tmp)
			if err == io.EOF || l < config.MessageReadBufferLength {
				msg = append(msg, tmp[:l]...)
				break
			}
			msg = append(msg, tmp...)
		}
		//如果发送消息时间间隔小于规定时间，不会被发送
		if time.Now().Truncate(config.MessageTimeInterval).Before(node.LastSendTime) {
			node.SendMessage(system.ErrorTalkTooFastMessage)
			continue
		}

		register.ProcessMessageFunc(msg, node)
	}
}
