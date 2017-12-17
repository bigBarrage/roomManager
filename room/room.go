package room

import (
	"sync"
	"time"
)

type RoomInfo struct {
	RoomID         string        //房间ID
	Rows           []*RowList    //房间多行Slice
	Lock           *sync.RWMutex //整个房间的锁
	Length         int64         //房间客户端数量
	LastChangeTime time.Time     //最后一次更新时间
}

type RowList struct {
	Nodes   []*ClientNode //节点列表
	RowLock *sync.Mutex   //读写锁
}
