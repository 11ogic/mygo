package algo

import (
	"fmt"
	"testing"
)

func LIS(nums []int) []int {
	result := [][]int{{nums[0]}}
	for i := 0; i < len(nums); i++ {
		for j := len(result) - 1; j >= 0; j-- {
			line := result[j]
			tail := line[len(line)-1]
			if nums[i] > tail {
				result = append(result, append(line, nums[i]))
				break
			} else if nums[i] < tail && j == 0 {
				result[j] = []int{nums[i]}
			}
		}
	}
	return result[len(result)-1]
}

func TestLIS(t *testing.T) {
	result := LIS([]int{3, 2, 6, 7, 8, 1, 2, 8, 9, 5, 7, 8, 9, 11})
	fmt.Println(result)
}
