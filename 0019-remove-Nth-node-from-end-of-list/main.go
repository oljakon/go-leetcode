package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var fast, slow *ListNode
	fast = head
	slow = head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	if fast == nil {
		head = head.Next
		return head
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return head
}
