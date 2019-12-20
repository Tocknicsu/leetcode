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
func partition(head *ListNode, x int) *ListNode {
	nhead := []*ListNode{nil, nil}
	nnow := []*ListNode{nil, nil}
	now := head
	for now != nil {
		id := 0
		if now.Val >= x {
			id = 1
		}
		if nhead[id] == nil {
			nhead[id] = now
			nnow[id] = now
		} else {
			nnow[id].Next = now
			nnow[id] = nnow[id].Next
		}
		now = now.Next
	}
	if nhead[0] != nil {
		nnow[0].Next = nhead[1]
		if nnow[1] != nil {
			nnow[1].Next = nil
		}
		return nhead[0]
	} else {
		return nhead[1]
	}
}

func main() {

	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  2,
							Next: nil,
						},
					},
				},
			},
		},
	}
	x := 3
	partition(input, x).Print()

}
