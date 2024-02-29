package algo

import (
	"testing"
)

/*
* O(n * log n)
 */
func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	pivot := nums[mid]
	left := make([]int, 0)
	right := make([]int, 0)
	for idx, val := range nums {
		if idx == mid {
			continue
		}
		if val > pivot {
			right = append(right, val)
		} else if val < pivot {
			left = append(left, val)
		}
	}
	return append(quickSort(left), append([]int{pivot}, quickSort(right)...)...)
}

func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func TestSort(t *testing.T) {
	arr := []int{3, 12, 51, 65, 23, 6, 27, 63, 45, 2, 3, 6, 3, 7, 458, 74, 1}
	TimeSpent(bubbleSort, arr)
}
