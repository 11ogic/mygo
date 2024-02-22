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

func binarySearch(arr []int, target int, low int, high int) int {
	mid := (low + high) / 2
	if low > high {
		return -1
	}

	if arr[mid] > target {
		return binarySearch(arr, target, low, mid-1)
	} else if arr[mid] < target {
		return binarySearch(arr, target, mid+1, high)
	} else {
		return mid
	}
}

func main() {
	arr := fbn(20)
	fmt.Printf("%v\n", arr)
	bubbleSort([]int{30, 1, 2, 81, 49, 20, 99})
	nums := []int{1, 3, 5, 6, 7, 8, 9, 11, 22, 33, 45, 66, 77, 99}
	index := binarySearch(nums, 33, 0, len(nums))
	fmt.Println(index)
}
