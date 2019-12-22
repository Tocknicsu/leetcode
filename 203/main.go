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

func removeElements(head *ListNode, k int) *ListNode {
	var ret, tail *ListNode
	now := head
	for now != nil {
		if now.Val != k {
			if ret == nil {
				ret = now
				tail = now
			} else {
				tail.Next = now
				tail = now
			}
		}
		now = now.Next
	}
	if tail != nil {
		tail.Next = nil
	}
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

	removeElements(input, k).Print()
	// input.Print()
}
