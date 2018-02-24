package room

import (
	"sync"
	"time"
)

//房间信息结构体
type RoomInfo struct {
	RoomID         string        //房间ID
	Rows           []*RowList    //房间多行Slice
	Lock           *sync.RWMutex //整个房间的锁
	Length         int64         //房间客户端数量
	HasBrokenNode  bool          //是否出现了坏死节点（需要整理）
	LastChangeTime time.Time     //最后一次更新时间
}

//行结构体
type RowList struct {
	Nodes   []*ClientNode //节点列表
	RowLock *sync.Mutex   //读写锁
}
