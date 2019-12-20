package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	if l == nil {
		fmt.Print("NULL\n")
		return
	}
	fmt.Print(l.Val, "->")
	l.Next.Print()
}

func deleteDuplicates(head *ListNode) *ListNode {
	now := head
	for now != nil {
		if now.Next != nil {
			if now.Val == now.Next.Val {
				now.Next = now.Next.Next
			} else {
				now = now.Next
			}
		} else {
			now = now.Next
		}
	}
	return head
}

func main() {
	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
		},
	}
	deleteDuplicates(input).Print()
}
