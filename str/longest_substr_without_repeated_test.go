package str

import (
	"fmt"
	"testing"
)

/**
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
示例 4:

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

输入: s = ""
输出: 0


0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成
*/

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	max := 0
	preMax := 1
	m := make(map[uint8]int)
	for i := 0; i < len(s); i++{
		l := preMax-1
		out := false
		for j := i + preMax-1; j < len(s); j++ {
			v := s[j]
			if _, ok := m[v]; ok {
				out = true
				delete(m, s[i])
			} else {
				m[v] = i
				l++
			}
			if l > max {
				max = l
			}
			preMax = l
			if out {
				break
			}
		}
	}

	return max
}

func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}
	max := 0
	length := 0
	nextStart := 0
	m := make(map[int32]int)
	for i, ch := range s {
		if index, ok := m[ch]; ok {
			length = i - index

			for _,v := range s[nextStart:index+1]{
				delete(m, v)
			}
			nextStart = index + 1
		} else {
			length++
		}
		m[ch] = i
		if length > max {
			max = length
		}
	}

	return max
}

func lengthOfLongestSubstring3(s string) int {
	var flag [128] int
	res := 0
	cnt := 0
	for i, c := range(s) {
		if (flag[c] == 0) {
			flag[c] = 1
			cnt++
			if (cnt > res) {
				res = cnt
			}
			continue
		}
		for _, cc := range(s[i-cnt:i]) {
			if (cc == c) {
				break
			}
			flag[cc] = 0
			cnt--
		}
	}
	return res
}


func Test(t *testing.T) {
	var str = ""
	var res = 0
	//str = "1234567adweasdawe"
	//res = lengthOfLongestSubstring(str)
	//fmt.Println(res)
	//res = lengthOfLongestSubstring2(str)
	//fmt.Println(res)
	//
	//str = "aaaaa"
	//res = lengthOfLongestSubstring(str)
	//fmt.Println(res)
	//res = lengthOfLongestSubstring2(str)
	//fmt.Println(res)
	//
	//str = ""
	//res = lengthOfLongestSubstring(str)
	//fmt.Println(res)
	//res = lengthOfLongestSubstring2(str)
	//fmt.Println(res)
	//
	//str = "123456781a1aqazwsx"
	//res = lengthOfLongestSubstring(str)
	//fmt.Println(res)
	//res = lengthOfLongestSubstring2(str)
	//fmt.Println(res)
	//
	//
	//str = "aabaab!bb"
	//res = lengthOfLongestSubstring(str)
	//fmt.Println(res)
	//res = lengthOfLongestSubstring2(str)
	//fmt.Println(res)


	str = "1234567770123456789"
	res = lengthOfLongestSubstring(str)
	fmt.Println(res)
	res = lengthOfLongestSubstring2(str)
	fmt.Println(res)
	res = lengthOfLongestSubstring3(str)
	fmt.Println(res)
}
