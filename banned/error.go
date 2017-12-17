package banned

import "errors"

var (
	WORDS_IS_BANNED = errors.New("word in message is banned")
	IP_IS_BANNED    = errors.New("ip is banned")
	UID_IS_BANNED   = errors.New("user id is banned")
)
