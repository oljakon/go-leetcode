package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resultSum := &ListNode{}
	curr := resultSum

	var carry, sum int

	for l1 != nil || l2 != nil || carry != 0 {
		sum = carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		} else {
			l1 = nil
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		} else {
			l2 = nil
		}

		carry = sum / 10
		sum := sum % 10

		curr.Next = &ListNode{sum, nil}
		curr = curr.Next
	}

	return resultSum.Next
}
