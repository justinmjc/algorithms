package leetcode

import (
	"fmt"
	"testing"
	"time"
)

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	MoveZeroes(nums)
	time.Sleep(1 * time.Second)
	fmt.Println(nums)
}
