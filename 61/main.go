package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) Print() {
	if node == nil {
		fmt.Print("NULL\n")
		return
	}
	fmt.Print(node.Val, "->")
	node.Next.Print()
}

func rotateRight(head *ListNode, k int) *ListNode {
	now := head
	len := 0
	for now != nil {
		len = len + 1
		now = now.Next
	}

	if len == 0 {
		return head
	}

	k = k % len

	now = head
	for ; k != 0; k = k - 1 {
		now = now.Next
	}
	back := head
	for now.Next != nil {
		now = now.Next
		back = back.Next
	}
	now.Next = head
	ret := back.Next
	back.Next = nil
	return ret
}

func main() {
	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	k := 1

	rotateRight(input, k).Print()
	// input.Print()
}
