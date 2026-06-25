package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	_, newHead := helper(head)
	return newHead
}

func helper(head *ListNode) (*ListNode, *ListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}

	nextCopy := head.Next
	head.Next = nil

	elem, newHead := helper(nextCopy)
	elem.Next = head

	return head, newHead
}
