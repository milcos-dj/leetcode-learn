package main

import (
	"fmt"
	"testing"
)

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
