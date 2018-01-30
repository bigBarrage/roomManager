package system

type NodeMessage struct {
	MessageType   int         `json:"type"`    //消息类型，类型为NODE_MESSAGE_TYPE组
	MessageTarget int         `json:"target"`  //发送目标
	RommId        string      `json:"room_id"` //房间ID
	Body          interface{} `json:"body"`    //消息体
}

const (
	_                                     = iota
	NODE_MESSAGE_TYPE_IN_HALL             //添加节点入大厅
	NODE_MESSAGE_TYPE_CLOSE_ROOM          //关闭房间服务
	NODE_MESSAGE_TYPE_CHANGE_ROOM         //节点房间变更
	NODE_MESSAGE_TYPE_SEND_MESSAGE        //节点发送消息
	NODE_MESSAGE_TYPE_CLEAN_ROOM          //清理房间垃圾节点
	NODE_MESSAGE_TYPE_FILL_USERINFO       //完善用户信息
	NODE_MESSAGE_TYPE_RELOAD_BANNED_IP    //重载IP地址黑名单
	NODE_MESSAGE_TYPE_RELOAD_BANNED_USER  //重载用户ID黑名单
	NODE_MESSAGE_TYPE_RELOAD_BANNED_WORDS //重载敏感词列表
	NODE_MESSAGE_TYPE_ROOM_NOT_EXISTS     //房间不存在
)

const (
	_                                  = iota
	MESSAGE_TARGET_BROADCASTINGSTATION //发送目标：广播站
	MESSAGE_TARGET_ROOM                //发送目标：房间
)
