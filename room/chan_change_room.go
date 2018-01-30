package room

import (
	"sync"

	"github.com/bigBarrage/roomManager/config"
)

//被Main Recive调用
//节点更换房间时被调用的房间
//因为前面已经修改了RoomID，所以这里只需要把接受者节点的指针重新对接到新房间的节点上，就可以了
//新的房间更新逻辑：
//1.扫描所有列，找到长度小于最大长度的，加入队尾
//2.修改结点中索引信息
func changeRoom(roomInfo *RoomInfo, node *ClientNode) {
	roomInfo.Lock.Lock()
	defer roomInfo.Lock.Unlock()
	//如果房间为空，则添加一个行
	if len(roomInfo.Rows) == 0 {
		row := &RowList{}
		row.RowLock = &sync.Mutex{}
		row.Nodes = make([]*ClientNode, 0, config.MaxLengthOfRows)
		roomInfo.Rows = append(roomInfo.Rows, row)
	}
	//设定是否添加完成标记
	addSuccess := false
	//寻找一个未满的行，把新的节点加进去
	for _, v := range roomInfo.Rows {
		if len(v.Nodes) < config.MaxLengthOfRows {
			v.RowLock.Lock()
			v.Nodes = append(v.Nodes, node)
			addSuccess = true
			v.RowLock.Unlock()
			break
		}
	}
	//如果当前房间所有的行都是满的，则创建一个新行，把节点放进去
	if !addSuccess {
		row := &RowList{}
		row.RowLock = &sync.Mutex{}
		row.Nodes = make([]*ClientNode, 0, config.MaxLengthOfRows)
		row.Nodes = append(row.Nodes, node)
		roomInfo.Rows = append(roomInfo.Rows, row)
	}
}
