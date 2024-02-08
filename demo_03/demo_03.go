package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	str1 := "今天天气真不错"
	newStr1 := strings.Replace(str1, "今天", "昨天", 1)
	fmt.Println(newStr1)

	str2 := "! hello golang! !"
	newStr2 := strings.Trim(str2, "! ")
	fmt.Println(newStr2)

	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Format("2006.01.02 15:04:05"))
}
