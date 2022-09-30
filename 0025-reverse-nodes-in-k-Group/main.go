package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	res := &ListNode{0, head}
	prev, curr := res, head

	if head == nil || head.Next == nil {
		return head
	}

	for curr != nil && curr.Next != nil {
		nextPtr := curr.Next.Next
		second := curr.Next

		curr.Next = nextPtr
		second.Next = curr
		prev.Next = second

		prev = curr
		curr = nextPtr
	}

	return res.Next
}
