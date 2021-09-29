package str

import (
	"fmt"
	"testing"
)

/**
给你一个字符串 s，找到 s 中最长的回文子串。

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

输入：s = "cbbd"
输出："bb"

输入：s = "a"
输出："a"

输入：s = "ac"
输出："a"
*/

func longestPalindrome(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}
	if l == 1 {
		return s
	}
	max, palindrome := 0, ""
	for i := 0; i < l; i++ {
		for j := i + 1; j <= l; j++ {
			str := s[i:j]
			if isPalindrome(str) && len(str) > max {
				max = len(str)
				palindrome = str
			}
		}

	}
	return palindrome
}

func isPalindrome(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

/**
最长公共子串计算方式
eg: caba -> abac
						i
			0	1	2	3
			c	a	b	a
	0	a	0	1	0	1
	1	b	0	0	2	0
j	2	a	0	1	0	3
	3	c	1	0	0	0

首先 i，j 始终指向子串的末尾字符。所以 j 指向的红色的 a 倒置前的下标是 beforeRev = length-1-j=4-1-2=1，
对应的是字符串首位的下标，我们还需要加上字符串的长度才是末尾字符的下标，也就是 beforeRev+arr\[i][j]-1=1+3-1=3，
因为 arr\[i][j] 保存的就是当前子串的长度，也就是图中的数字 3。此时再和它与 i 比较，如果相等，则说明它是我们要找的回文串。

作者：windliang
链接：https://leetcode-cn.com/problems/longest-palindromic-substring/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-bao-gu/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/
func longestPalindromeMaxPublicSubStr(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}
	if l == 1 {
		return s
	}
	var arr = make([][]int, l)
	for i, _ := range arr {
		arr[i] = make([]int, l)
	}
	maxLen := 0
	maxEnd := 0
	for i := 0; i < l; i++ {
		for j := l - 1; j >= 0; j-- {
			if s[i] == s[j] {
				if i == 0 || j == l-1 {
					arr[i][l-j-1] = 1
				} else {
					arr[i][l-j-1] = arr[i-1][l-j-1-1] + 1
				}

			}
			if arr[i][l-j-1] > maxLen && j+arr[i][l-j-1]-1 == i {

				maxLen = arr[i][l-j-1]
				maxEnd = i
			}

		}
	}
	return s[maxEnd-maxLen+1 : maxEnd+1]
}

func longestPalindromeDiffusion(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}
	if l == 1 {
		return s
	}


	left, right, maxLen, curLen, startPos := 0, 0, 0, 1, 0

	for i := 0; i < l; i++ {
		left = i - 1
		right = i + 1
		for left >= 0 && s[i] == s[left] {
			curLen++
			left--
		}

		for right < l && s[i] == s[right] {
			curLen++
			right++
		}

		for left >= 0 && right < l && s[left] == s[right] {
			left--
			right++
			curLen = curLen + 2
		}
		if curLen > maxLen {
			maxLen = curLen
			startPos = left
		}
		curLen = 1
	}

	return s[startPos+1:startPos+maxLen+1]
}

func Test_palindrome(t *testing.T) {
	var s = "321012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210012321001232100123210123210012321001232100123210123"
	var res = ""
	res = longestPalindrome(s)
	fmt.Println(res)

	res = longestPalindromeMaxPublicSubStr(s)
	fmt.Println(res)
	res = longestPalindromeDiffusion(s)
	fmt.Println(res)

}
