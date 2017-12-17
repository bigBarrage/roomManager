package banned

var bannedUserIDList userIDNode

type userIDNode map[byte]userIDNode

func SetUserIDs(list []string) {
	if len(list) == 0 {
		return
	}

	bannedUserIDList = make(userIDNode)
	for _, userID := range list {
		currNode := bannedUserIDList
		for _, b := range []byte(userID) {
			if _, ok := currNode[b]; !ok {
				currNode[b] = make(userIDNode)
			}
			currNode = currNode[b]
		}
	}
}

func IsBannedUserID(userID string) bool {
	currNode := bannedUserIDList
	var ok bool
	for _, i := range []byte(userID) {
		if currNode, ok = currNode[i]; !ok {
			return false
		}
	}
	return true
}
