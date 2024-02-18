package main

import "fmt"

func fbn(n uint) []int {
	arr := make([]int, n)
	for i, _ := range arr {
		if i < 2 {
			arr[i] = 1
		} else {
			arr[i] = arr[i-1] + arr[i-2]
		}
	}
	return arr
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] < arr[j] {
				temp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Printf("%v", arr)
}

func LIS(arr []int) {

}

func main() {
	arr := fbn(20)
	fmt.Printf("%v\n", arr)
	bubbleSort([]int{30, 1, 2, 81, 49, 20, 99})
}
