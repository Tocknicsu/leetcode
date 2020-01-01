package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	if l == nil {
		fmt.Print("Null\n")
	} else {
		fmt.Print(l.Val, "->")
		l.Next.Print()
	}
}

func Reverse(head *ListNode) *ListNode {
	ret := head
	now := head
	for now != nil {
		ret = now
		nextNow := now.Next
		if now.Next != nil {
			now.Next.Next = now
		}
		now = nextNow
	}
	head.Next = nil
	return ret
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	Reverse(headA).Print()
	return nil
}

func main() {
	section := &ListNode{
		8,
		&ListNode{
			4,
			&ListNode{
				5,
				nil,
			},
		},
	}
	a := &ListNode{
		4,
		&ListNode{
			1,
			section,
		},
	}
	b := &ListNode{
		5,
		&ListNode{
			0,
			&ListNode{
				1,
				section,
			},
		},
	}
	getIntersectionNode(a, b).Print()
}
