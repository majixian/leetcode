/*
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var (
		a1 = &ListNode{1, nil}
		a2 = &ListNode{2, nil}
		a3 = &ListNode{3, nil}
		b1 = &ListNode{4, nil}
		b2 = &ListNode{5, nil}
		b3 = &ListNode{9, nil}
	)

	a1.Next = a2
	a2.Next = a3

	b1.Next = b2
	b2.Next = b3

	res := addTwoNumbers(a1, b1)
	for res != nil {
		fmt.Println("%d", res.Val)
		res = res.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var begin *ListNode
	var end *ListNode
	flag := 0

	for nil != l1 || nil != l2 {
		sum := 0

		if nil != l1 {
			sum += l1.Val
			l1 = l1.Next
		}

		if nil != l2 {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += flag
		flag = 0

		if sum >= 10 {
			flag = 1
			sum = sum % 10
		}

		node := &ListNode{
			sum,
			nil,
		}

		if nil == begin {
			begin = node
			end = node
		} else {
			end.Next = node
			end = node
		}
	}

	if flag > 0 {
		end.Next = &ListNode{
			flag,
			nil,
		}
	}
	return begin
}
