package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Print(list *ListNode) {
	for list != nil {
		fmt.Print(list.Val)
		list = list.Next
		if list != nil {
			fmt.Print("->")
		}
	}
	fmt.Println()
}

func swapPairs(head *ListNode) *ListNode {
	var h, now *ListNode

	for head != nil {
		if head.Next != nil {
			nextHead := head.Next.Next
			if h == nil {
				h = head.Next
				h.Next = head
				now = head
			} else {
				now.Next = head.Next
				now.Next.Next = head
				now = head
			}
			now.Next = nil
			head = nextHead
		} else {
			if h == nil {
				h = head
			} else {
				now.Next = head
			}
			head = head.Next
		}
	}

	return h
}

func main() {
	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	Print(swapPairs(input))
}
