package algo

import (
	"fmt"
	"testing"
)

/**
value  5 10 3 6 3
weight 2  5 1 4 3
*/

var (
	value  = []int{5, 10, 3, 6, 3}
	weight = []int{2, 5, 1, 4, 3}
)

func calc(capacity int) {
	dp := make([][]int, len(value))
	for i := 0; i < len(value); i++ {
		dp[i] = make([]int, capacity+1)
		for j := 0; j <= capacity; j++ {
			if i == 0 {
				if j >= weight[i] {
					dp[i][j] = value[i]
				}
				continue
			}
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
				continue
			}
			dp[i][j] = max(value[i]+dp[i-1][j-weight[i]], dp[i-1][j])
		}
	}
	fmt.Println(dp)
}

func optimization(capacity int) {
	line := make([]int, capacity+1)
	for i := 0; i <= capacity; i++ {
		if weight[0] <= i {
			line[i] = value[0]
		}
	}
	for i := 1; i < len(value); i++ {

	}
	fmt.Println(line)
}

func TestBag(t *testing.T) {
	calc(6)
	optimization(6)
}
