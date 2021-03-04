package leetcode

//给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
//在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。
//找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/container-with-most-water

func MaxArea(height []int) int {
	res := 0
	for left, right := 0, len(height)-1; left < right; {

		min := height[left]
		if height[right] < min {
			min = height[right]
		}
		area := min * (right - left)
		if area > res {
			res = area
		}

		//这里是本题的关键，不是比较height[i] 和height[i+1]，这样没意义
		//要比较左右的大小去移动，很容易上来就比较相邻的数。但是对面积没意义
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return res
}
