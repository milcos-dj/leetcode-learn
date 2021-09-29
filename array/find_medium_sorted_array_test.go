package array

import (
	"fmt"
	"testing"
)
/**
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。


输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2


输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000

输入：nums1 = [], nums2 = [1]
输出：1.00000

输入：nums1 = [2], nums2 = []
输出：2.00000


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/



/**
合并两个数组再去中间数
 */
func findMedianSortedArrays_merge(nums1 []int, nums2 []int) float64 {
	res := mergeSortedArray(nums1, nums2)
	fmt.Println(res)
	l := len(res)
	median := 0.0
	if l % 2 ==0 {
		median = float64(res[l/2 -1] + res[l/2]) / 2.0
	} else {
		median = float64(res[l/2])
	}
	return median
}

func mergeSortedArray(nums1 []int, nums2 []int) []int{
	l := len(nums1) + len(nums2)
	var res = make([]int, l)
	start1 := 0
	start2 := 0
	index := 0
	for start1 < len(nums1) && start2 < len(nums2) {
		if nums1[start1] <= nums2[start2] {
			res[index] = nums1[start1]
			start1++
		} else {
			res[index] = nums2[start2]
			start2++
		}
		index++
	}

	if start1 < len(nums1) {
		for _,v := range nums1[start1:] {
			res[index] = v
			index++
		}
	}

	if start2 < len(nums2) {
		for _,v := range nums2[start2:] {
			res[index] = v
			index++
		}
	}
	return res
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	total := l1 + l2
	even := total % 2 == 0
	median := total / 2
	start := 0
	start1 := 0
	start2 := 0
	count := 0
	flg := true	// true表示当前循环去nums1的值 否则取nums2的值
	for start <= median  {
		if start1 >= l1 {
			flg = false
			start2++
		} else if start2 >= l2 {
			flg = true
			start1++
		} else if nums1[start1] <= nums2[start2] {
			flg = true
			start1++
		} else {
			flg = false
			start2++
		}

		// 数组长度为偶数
		if even && (start == median-1 || start == median) {
			if flg {
				count += nums1[start1-1]
			} else {
				count += nums2[start2-1]
			}
		}

		if !even && start == median {
			if flg {
				count += nums1[start1-1]
			} else {
				count += nums2[start2-1]
			}
		}


		start++
	}
	res := 0.0
	if even {
		 res = float64(count) / 2.0
	} else {
		 res = float64(count)
	}
	return res
}

func findMedianSortedArrays_2(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	total := l1 + l2
	even := total % 2 == 0
	median := total / 2
	start := 0
	start1 := 0
	start2 := 0
	preNum := 0
	curNum := 0
	for start <= median  {
		preNum = curNum
		if start1 >= l1 {
			curNum = nums2[start2]
			start2++
		} else if start2 >= l2 {
			curNum = nums1[start1]
			start1++
		} else if nums1[start1] <= nums2[start2] {
			curNum = nums1[start1]
			start1++
		} else {
			curNum = nums2[start2]
			start2++
		}
		start++
	}

	res := 0.0
	if even {
		res = float64(preNum + curNum) / 2.0
	} else {
		res = float64(curNum)
	}
	return res
}



func TestMedianSortedArrays(t *testing.T) {
	var nums1 = []int{1,3,4,6}
	var nums2  = []int{2,5,7,8}
	res := 0.0
	res = findMedianSortedArrays_merge(nums1, nums2)
	fmt.Println(res)
	res = findMedianSortedArrays(nums1, nums2)
	fmt.Println(res)
	res = findMedianSortedArrays_2(nums1, nums2)
	fmt.Println(res)
}