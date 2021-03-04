package leetcode

import "fmt"

/*

给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/move-zeroes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
思路：
1 loop  cout zeros,
2 开新数组，0放后，非零放前 (不符合题解)
3 循环整个数据遇到0删除，然后数组末尾加个0  时间复杂度0（n^2）
*/
func moveZeroes(nums []int) {
	p := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[p-1] = nums[i]
			p = p - 1
		}
	}
}

/*
思路：
    循环整个数据遇到0删除，然后数组末尾加个0  时间复杂度0（n^2）
*/

func MoveZeroes1(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
		}
	}
	fmt.Println(nums)
}

func MoveZeroes2(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
		}
	}

	fmt.Println(nums)
}
