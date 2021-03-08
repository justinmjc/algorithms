package leetcode

func MoveZeroes(nums []int) {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			k++
			nums = nums
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	for j := 0; j < k; j++ {
		nums = append(nums, 0)
	}
}
