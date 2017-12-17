package room

import "errors"

var (
	ROOM_NOT_EXISTS         = errors.New("room is not exists")
	CLIENT_IS_DISABLED_READ = errors.New("client is disabled read")
)
