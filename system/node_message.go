package system

//roomManager内传递消息用的结构
type NodeMessage struct {
	MessageType   int         `json:"type"`    //消息类型，类型为NODE_MESSAGE_TYPE组
	MessageTarget int         `json:"target"`  //发送目标
	RoomId        string      `json:"room_id"` //房间ID
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
	NODE_MESSAGE_TYPE_PING                //PING
	NODE_MESSAGE_TYPE_PONG                //PONG
)

const (
	//一些错误信息
	_                                 = iota + 400
	MESSAGE_ERROR_TYPE_READDISABLES   //被禁言
	MESSAGE_ERROR_TYPE_USERID_IS_NULL //用户信息不完善
	MESSAGE_ERROR_TYPE_UNKNOWN        //未知错误
	MESSAGE_ERROR_TYPE_TALK_TOO_FAST  //发言速度太快
)

const (
	_                                  = iota
	MESSAGE_TARGET_BROADCASTINGSTATION //发送目标：广播站
	MESSAGE_TARGET_ROOM                //发送目标：房间
)

//用户节点接口
type MessageNode interface {
	ChangeRoom(RoomID string)
	SendMessageToRoom(message interface{})
	SendMessage(message interface{})
	ChangeUserID(userID string)
}

var (
	ErrorReadDisabledMessage *NodeMessage
	ErrorUserIdIsNullMessage *NodeMessage
	ErrorUnknownMessage      *NodeMessage
	ErrorTalkTooFastMessage  *NodeMessage
)

func init() {
	ErrorReadDisabledMessage = &NodeMessage{MessageType: MESSAGE_ERROR_TYPE_READDISABLES}
	ErrorUserIdIsNullMessage = &NodeMessage{MessageType: MESSAGE_ERROR_TYPE_USERID_IS_NULL}
	ErrorUnknownMessage = &NodeMessage{MessageType: MESSAGE_ERROR_TYPE_UNKNOWN}
	ErrorTalkTooFastMessage = &NodeMessage{MessageType: MESSAGE_ERROR_TYPE_TALK_TOO_FAST}
}
