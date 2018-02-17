package banned

import "errors"

var (
	WORDS_IS_BANNED = errors.New("word in message is banned") //发送内容包含敏感词
	IP_IS_BANNED    = errors.New("ip is banned")              //IP地址被设置在黑名单内
	UID_IS_BANNED   = errors.New("user id is banned")         //用户ID被设置在黑名单内
)
