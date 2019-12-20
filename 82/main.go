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
	var ret, tail *ListNode
	now := head
	for now != nil {
		if now.Next == nil || now.Val != now.Next.Val {
			if ret == nil {
				ret = now
				tail = now
			} else {
				tail.Next = now
				tail = tail.Next
			}
			now = now.Next
			tail.Next = nil
		} else {
			for now.Next != nil && now.Val == now.Next.Val {
				now.Next = now.Next.Next
			}
			now = now.Next
		}
	}
	return ret
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
