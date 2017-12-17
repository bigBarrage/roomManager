package banned

import (
	"testing"
)

func TestSetUserIDs(t *testing.T) {
	list := []string{
		"DFS78FDS",
		"FDS45JKL",
		"JKN567SD",
	}
	SetUserIDs(list)
	t.Log("SetUerIDs end")
}

func TestIsBannedUserID(t *testing.T) {
	list := []string{
		"DFS78FD5",
		"FDS45JKL",
		"JKN5675D",
	}
	for _, id := range list {
		t.Log(id, "的鉴定结果是：", IsBannedUserID(id))
	}
}
