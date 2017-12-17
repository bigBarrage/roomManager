package banned

import "github.com/logan-go/ACautomaton"

var bannedWordList wordNode

type wordNode map[rune]wordNode

func SetWordList(list []string) {
	ACautomaton.SetList(list)
}

func IsBannedWords(msg string) bool {
	return ACautomaton.Check(msg)
}