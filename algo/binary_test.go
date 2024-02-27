package algo

import (
	"testing"
)

/*
O(log n)
*/
func binarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func recursionBinarySearch(nums []int, target int, low int, high int) int {
	mid := (low + high) / 2
	if low > high {
		return -1
	}
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return recursionBinarySearch(nums, target, low, mid-1)
	} else {
		return recursionBinarySearch(nums, target, mid+1, high)
	}
}

func TestBinary(t *testing.T) {
	arr := make([]int, 100000000)
	for i, _ := range arr {
		arr[i] = i * 3
	}
	target := 927
	low := 0
	high := len(arr) - 1
	TimeSpent(recursionBinarySearch, arr, target, low, high)
	TimeSpent(binarySearch, arr, target)
}
