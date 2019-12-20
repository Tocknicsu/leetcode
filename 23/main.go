package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int           { return len(h) }
func (h ListNodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h ListNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
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

func mergeKLists(list []*ListNode) *ListNode {
	h := &ListNodeHeap{}
	heap.Init(h)

	for _, node := range list {
		if node != nil {
			heap.Push(h, node)
		}
	}

	if h.Len() == 0 {
		return nil
	}

	head := heap.Pop(h).(*ListNode)
	now := head
	if now.Next != nil {
		heap.Push(h, now.Next)
	}

	for h.Len() != 0 {
		now.Next = heap.Pop(h).(*ListNode)
		now = now.Next
		if now.Next != nil {
			heap.Push(h, now.Next)
		}
	}

	return head
}

func main() {
	input := []*ListNode{
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  6,
				Next: nil,
			},
		},
	}
	for _, x := range input {
		Print(x)
	}
	Print(mergeKLists(input))
}
