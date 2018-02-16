package room

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"

	"github.com/bigBarrage/roomManager/config"
	"github.com/bigBarrage/roomManager/register"
	"github.com/gorilla/websocket"
)

var (
	BroadcastingConnection *websocket.Conn
	err                    error
)

func ConnBroadcastingStation() {
	if config.UseBoradcasting {
		tryToConnBroadcastingStation()
		//处理堵到消息之后
		go func() {
			for {
				mType, reader, err := BroadcastingConnection.NextReader()
				if mType == websocket.CloseMessage || mType == -1 {
					tryToConnBroadcastingStation()
				}
				if err != nil {
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
				//未完：这里需要加入读到消息之后的处理器
				fmt.Println("从广播站收到消息：", string(msg))
				register.ProcessMessageFromBroadcastingFunc(msg)
			}
		}()
	}
}

func tryToConnBroadcastingStation() {
	//尝试创建广播站的连接
	bo := config.GetBroadcastingStation()
	for connCount := 0; connCount < 10; connCount++ {
		conn, err := net.Dial("tcp", bo.Host+":"+fmt.Sprint(bo.Port))
		if err != nil {
			fmt.Println("广播站连接失败：", err)
			continue
		}
		u := url.URL{}
		u.Host = bo.Host + ":" + fmt.Sprint(bo.Port)
		u.Path = bo.Path
		u.Scheme = "ws"

		h := http.Header{}
		BroadcastingConnection, _, err = websocket.NewClient(conn, &u, h, 1024, 1024)
		if err != nil {
			fmt.Println("websocket连接创建失败：", err)
			conn.Close()
			continue
		}
		fmt.Println("广播站连接创建成功！")
		return
	}
	os.Exit(0)
}
