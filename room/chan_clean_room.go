package room

import (
	"fmt"
	"time"

	"github.com/bigBarrage/roomManager/config"
)

//被Main Revice调用
//清理当前房间里面不属于自己房间的节点
func cleanRoom(roomInfo *RoomInfo) {
	//打印整理前的房间信息
	fmt.Println("========================整理前", time.Now().Format("2006-01-02 15:04:05"), "========================")

	startTime := time.Now()
	//创建一个空的列组，准备装整理过的节点
	colList := make([]*RowList, 0, 5)

	//创建行
	rowList := &RowList{}
	rowList.Nodes = make([]*ClientNode, 0, config.MaxLengthOfRows)

	//本次调整时间
	currentUpdateTime := time.Now()

	//列加入列组
	colList = append(colList, rowList)

	//列和行当前索引
	colIndex := 0

	//循环列表内的节点
	for _, row := range roomInfo.Rows {
		for _, node := range row.Nodes {
			if node.RoomID != roomInfo.RoomID || node.IsAlive == false || node.UpdateTime == currentUpdateTime {
				continue
			}
			node.UpdateTime = currentUpdateTime
			//添加节点到正常节点
			tmpIndex := len(rowList.Nodes)
			if tmpIndex >= config.MaxLengthOfRows {
				tmpIndex = 0
				//创建行
				rowList = &RowList{}
				rowList.Nodes = make([]*ClientNode, 0, config.MaxLengthOfRows)
				colIndex++

				//列加入列组
				colList = append(colList, rowList)
				colList[colIndex].Nodes = append(colList[colIndex].Nodes, node)
			} else {
				colList[colIndex].Nodes = append(colList[colIndex].Nodes, node)
			}
		}
	}
	//整理完毕之后，把新的整理结果赋予房间
	roomInfo.Lock.Lock()
	defer roomInfo.Lock.Unlock()
	roomInfo.Rows = colList
	roomInfo.LastChangeTime = time.Now()
	endTime := time.Now()

	fmt.Println("行数：", len(roomInfo.Rows))
	for _, row := range roomInfo.Rows {
		fmt.Println("总列：", len(row.Nodes), "|", cap(row.Nodes))
	}
	fmt.Println("耗时：", endTime.Sub(startTime))
	fmt.Println("========================整理完毕", time.Now().Format("2006-01-02 15:04:05"), "========================")
}
