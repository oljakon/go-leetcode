package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	listLen := 1
	tail := head

	for tail.Next != nil {
		tail = tail.Next
		listLen++
	}

	k = k % listLen
	if k == 0 {
		return head
	}

	curr := head
	for i := 0; i < listLen-k-1; i++ {
		curr = curr.Next
	}

	newHead := curr.Next
	curr.Next = nil
	tail.Next = head

	return newHead
}
