package banned

import (
	"net"
)

var bannedIpList ipNode

type ipNode map[uint8]ipNode

//设置IP黑名单
func SetIpList(list []string) {
	if len(list) == 0 {
		return
	}
	bannedIpList = ipNode{}
	for _, ipStr := range list {
		ip := net.ParseIP(ipStr).To4()

		currNode := bannedIpList
		for _, v := range ip {
			if _, ok := currNode[v]; !ok {
				currNode[v] = make(ipNode)
			}
			currNode = currNode[v]

		}
	}
}

//检查IP是否在黑名单内
func IsBannedIP(ip string) bool {
	parsedIp := net.ParseIP(ip).To4()
	currNode := bannedIpList
	var ok bool
	for _, ip := range parsedIp {
		if currNode, ok = currNode[ip]; !ok {
			return false
		}
	}
	return true
}
