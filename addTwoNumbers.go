/*
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
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
