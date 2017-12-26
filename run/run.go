package run

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/bigBarrage/roomManager/banned"
	"github.com/bigBarrage/roomManager/config"
	"github.com/bigBarrage/roomManager/register"
	"github.com/bigBarrage/roomManager/room"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Run() error {
	if check() != nil {
		os.Exit(1)
	}
	bsopt := config.GetBroadcastingStation()
	http.HandleFunc("/"+bsopt.Path, handler)
	return http.ListenAndServe(fmt.Sprint(bsopt.Port), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	node := &room.ClientNode{}

	addrSlice := strings.Split(r.RemoteAddr, ":")
	node.IP = addrSlice[0]

	node.IsAlive = true
	node.Conn = conn
	node.Add()
	processMessage(node)
}

//做启动前检查，保证各项设置基本运行状况
func check() error {
	if register.ProcessMessageFunc == nil {
		fmt.Fprintln(os.Stderr, "消息处理方法未注册...")
		return errors.New("check failed")
	}
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
			continue
		}

		if banned.IsBannedUserID(node.UserID) {
			node.DisableRead = true
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

			register.ProcessMessageFunc(msg, node)
		}
	}
}
