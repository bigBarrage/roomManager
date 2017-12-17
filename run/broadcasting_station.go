package run

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"

	"github.com/bigBarrage/roomManager/config"
	"github.com/gorilla/websocket"
)

var (
	BroadcastingConnection *websocket.Conn
	err                    error
)

func init() {

	if config.UseBoradcasting {
		tryToConnBroadcastingStation()
		//处理堵到消息之后
		go func() {
			mType, reader, err := BroadcastingConnection.NextReader()
			if mType == websocket.CloseError || mType == -1 {
				tryToConnBroadcastingStation()
			}
			if err != nil {
				continue
			}
			msg := make([]byte, 1024)
			l, err := reader.Read(msg)
			if err != nil {
				continue
			}
			//未完：这里需要加入读到消息之后的处理器
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
		fmt.Println("创建成功！")
		return
	}
	os.Exit(0)
}
