package array

import (
	"fmt"
	"testing"
)

/**
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.

输入：l1 = [0], l2 = [0]
输出：[0]

输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	take := false
	next := res
	for {
		var val1 = 0
		var val2 = 0
		if l1 != nil {
			val1 = l1.Val
		}
		if l2 != nil {
			val2 = l2.Val
		}
		sum := val1 + val2

		divideRes := sum / 10
		if divideRes == 0 { // 两数相加小于10
			next.Val = sum
		} else { // 两数相加大于10
			val := sum % 10
			next.Val = val
			take = true
		}

		if l1 != nil && l1.Next != nil {
			l1 = l1.Next
			if take {
				l1.Val += 1
				take = false
			}

		}  else {
			l1 = nil
		}
		if l2 != nil && l2.Next != nil {
			l2 = l2.Next
			if take {
				l2.Val += 1
				take = false
			}

		} else {
			l2 = nil
		}

		if l1 == nil && l2 == nil {
			if take {
				next.Next = &ListNode{Val: 1}
			}
			break
		}
		next.Next = &ListNode{}
		next = next.Next
	}
	return res
}

func Test_main(t *testing.T) {
	l1 := &ListNode{Val: 9}
	next := l1
	next.Next = &ListNode{Val: 9}
	next = next.Next
	next.Next = &ListNode{Val: 9}
	next = next.Next
	next.Next = &ListNode{Val: 9}
	next = next.Next
	next.Next = &ListNode{Val: 9}
	next = next.Next
	next.Next = &ListNode{Val: 9}
	next = next.Next
	next.Next = &ListNode{Val: 9}
	next = next.Next


	l2 := &ListNode{Val: 9}
	next2 := l2
	next2.Next = &ListNode{Val: 9}
	next2 = next2.Next
	next2.Next = &ListNode{Val: 9}
	next2 = next2.Next
	next2.Next = &ListNode{Val: 9}

	res := addTwoNumbers(l1, l2)
	for res.Next != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
	fmt.Println(res.Val)
}
