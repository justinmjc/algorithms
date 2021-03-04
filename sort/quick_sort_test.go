package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{1, 3, 4, 3, 5, 0, 0, -3}
	QuickSort(arr)
	fmt.Println(arr)
}
