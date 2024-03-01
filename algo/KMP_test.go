package algo

import "testing"

func BF(str string, target string) int {
	strRune := []rune(str)
	targetRune := []rune(target)
	for i := 0; i <= len(strRune)-len(targetRune); i++ {
		match := true
		for j, v := range targetRune {
			if strRune[i+j] != v {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}

func KMP(str string, target string) int {
	return -1
}

func TestKMP(t *testing.T) {
	TimeSpent(BF, "今天天气真不错", "真不错")
}
