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

func reverseKGroup(head *ListNode, k int) *ListNode {
	var retHead *ListNode
	var lastTail *ListNode
	now := head
	for now != nil {
		segStart := now
		segNow := now
		ok := true
		for i := 1; i < k; i = i + 1 {
			if segNow.Next == nil {
				ok = false
				break
			}
			segNow = segNow.Next
		}

		if ok == false {
			if retHead == nil {
				retHead = segStart
			} else {
				lastTail.Next = segStart
			}
			break
		}

		now = segNow.Next

		segTail := segNow

		var segPre *ListNode
		segPre = nil
		segNow = segStart

		for i := 0; i < k; i = i + 1 {
			segNext := segNow.Next
			segNow.Next = segPre
			segPre = segNow
			segNow = segNext
		}

		if retHead == nil {
			retHead = segTail
			lastTail = segStart
		} else {
			lastTail.Next = segTail
			lastTail = segStart
		}

		// Print(segTail)
	}
	return retHead
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

	k := 5

	Print(reverseKGroup(input, k))
}
