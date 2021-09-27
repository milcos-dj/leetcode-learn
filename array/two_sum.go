package main

import (
	"fmt"
)
/**
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */





func twoSum(nums []int, target int) []int {

	for i:= 0; i < len(nums); i++ {
		for j:= i+1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i,j}
			}
		}
	}

	return nil
}

func twoSum2(nums []int, target int) []int {
	m := map[int]int{}
	for i:= 0; i < len(nums); i++ {
		if v, ok := m[target - nums[i]]; ok {
			return []int{v, i}
		} else {
			m[nums[i]] = i
		}
	}
	return nil
}

func main()  {
	arr := []int{2,7,11,15, 16, 17,18,19,20,21,22,23,24,25,26}
	fmt.Println(twoSum(arr, 9))
	fmt.Println(twoSum2(arr, 26))
}