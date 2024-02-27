package algo

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

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

func TimeSpent(fn interface{}, params ...interface{}) interface{} {
	function := reflect.ValueOf(fn)
	if function.Kind() != reflect.Func {
		panic("expected a function")
	}

	inner := make([]reflect.Value, len(params))
	for i, val := range params {
		inner[i] = reflect.ValueOf(val)
	}

	start := time.Now()
	result := function.Call(inner)

	elapsed := time.Since(start).Seconds()
	fmt.Printf("elapsed: %.10f  result: %v \n", elapsed, result[0].Int())

	return result
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
