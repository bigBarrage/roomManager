package banned

import "github.com/logan-go/ACautomaton"

var bannedWordList wordNode

type wordNode map[rune]wordNode

//设定敏感词名单
func SetWordList(list []string) {
	ACautomaton.SetList(list)
}

//检查内容中是否包含敏感词
func IsBannedWords(msg string) bool {
	return ACautomaton.Check(msg)
}
