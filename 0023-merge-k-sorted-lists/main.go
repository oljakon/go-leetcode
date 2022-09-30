package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		mergedLists := []*ListNode{}

		for i := 0; i < len(lists); i = i + 2 {
			l1 := lists[i]

			var l2 *ListNode
			if i+1 < len(lists) {
				l2 = lists[i+1]
			} else {
				l2 = nil
			}

			mergedLists = append(mergedLists, mergeTwoLists(l1, l2))
		}

		lists = mergedLists
	}

	return lists[0]
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	tail := res

	if list1 == nil && list2 == nil {
		return res.Next
	}

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	} else if list2 != nil {
		tail.Next = list2
	}

	return res.Next
}
