package banned

import "strings"

var bannedUserIDList []string

//设置用户黑名单
func SetUserIDs(list []string) {
	if len(list) == 0 {
		return
	}
	bannedUserIDList = make([]string, 0, 10)
	for _, id := range list {
		if strings.TrimSpace(id) != "" {
			bannedUserIDList = append(bannedUserIDList, strings.TrimSpace(id))
		}
	}
}

//检查用户是否在黑名单内
func IsBannedUserID(userID string) bool {
	if bannedUserIDList == nil || len(bannedUserIDList) == 0 {
		return false
	}
	for _, i := range bannedUserIDList {
		if i == userID {
			return true
		}
	}
	return false
}
