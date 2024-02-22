package main

import "fmt"

func LIS(nums []int) []int {
	result := [][]int{{nums[0]}}
	for i := 0; i < len(nums); i++ {
		for j := len(result) - 1; j >= 0; j-- {
			line := result[j]
			tail := line[len(line)-1]
			if tail < nums[i] {
				result = append(result, append(line, nums[i]))
				break
			} else if tail > nums[i] && j == 0 {
				result[j] = append([]int{}, nums[i])
			}
		}
	}
	return result[len(result)-1]
}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				temp := nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = temp
			}
		}
	}
	fmt.Println(nums)
}

func main() {
	result := LIS([]int{33, 3, 22, 5, 7, 6, 2, 8, 99, 15, 12, 16, 1, 9})
	fmt.Println(result)
	bubbleSort([]int{4, 2, 5, 6, 72, 8})
}
