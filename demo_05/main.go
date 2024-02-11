package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func isBigMonth(month int) bool {
	bigMonths := []int{1, 3, 5, 7, 8, 10, 12}
	for _, m := range bigMonths {
		if month == m {
			return true
		}
	}
	return false
}

func displayDays(years, month int) (isLeapYear bool, days int) {
	if month > 12 || month < 1 {
		panic("Please enter the correct month!")
	}
	isLeapYear = years%400 == 0 || (years%4 == 0 && years%100 != 0)
	if month == 2 {
		if isLeapYear {
			days = 29
		} else {
			days = 28
		}
	} else {
		if isBigMonth(month) {
			days = 31
		} else {
			days = 30
		}
	}
	fmt.Printf("%v年%v月有%v天", years, month, days)
	return
}

func ranDomGame() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	random := 1 + rand.Intn(99)
	fmt.Println("DEBUG: ", random)
	fmt.Println("输入猜的数字: ")
	answer := 0
	for i := 1; i <= 10; i++ {
		fmt.Scanln(&answer)
		if random == answer {
			if i == 1 {
				fmt.Println("你真是个天才")
			} else if i == 2 || i == 3 {
				fmt.Println("你很聪明，赶上我了")
			} else if i > 3 && i < 10 {
				fmt.Println("一般般")
			} else {
				fmt.Println("可算猜对啦")
			}
			return
		} else {
			if i == 10 {
				fmt.Println("说你点啥好呢")
				return
			}
			fmt.Printf("输入错误，还有%v机会\n", 10-i)
		}
	}
}

func getPrimeNumber(num int) {
	var (
		message string
		count   = 1
		sum     int
	)
	for i := 1; i <= num; i++ {
		flag := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			sum += i
			message += strconv.Itoa(i) + " "
			if count%5 == 0 {
				message += "sum: " + strconv.Itoa(sum)
				fmt.Println(message)
				sum = 0
				message = ""
			}
			count++
		}
	}
}

func judgeFishingOrDryingNets(day time.Time) {
	date := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	dateStr := date.Format("2006.01.02 15:04:05")
	diff := int((day.Unix() - date.Unix()) / 60 / 60 / 24)
	fmt.Println(diff)
	if diff%5 < 3 {
		fmt.Println(dateStr, " 这天打鱼呢")
	} else {
		fmt.Println(dateStr, " 这天晒网呢")
	}
}

func main() {
	//displayDays(2024, 2)
	//ranDomGame()
	//getPrimeNumber(200)
	//judgeFishingOrDryingNets(time.Date(1990, 1, 11, 0, 0, 0, 0, time.UTC))
}
