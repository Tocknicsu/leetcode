package main

import (
	"container/list"
	"fmt"
)

type MinStack struct {
	mi *list.List
	st *list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	ret := MinStack{
		list.New(),
		list.New(),
	}

	ret.mi.PushBack(2147483647)

	return ret
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (this *MinStack) Push(x int) {
	this.st.PushBack(x)
	this.mi.PushBack(Min(this.mi.Back().Value.(int), x))
}

func (this *MinStack) Pop() {
	this.st.Remove(this.st.Back())
	this.mi.Remove(this.mi.Back())
}

func (this *MinStack) Top() int {
	return this.st.Back().Value.(int)

}

func (this *MinStack) GetMin() int {
	return this.mi.Back().Value.(int)
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)

	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	obj.Pop()
	fmt.Println(obj.GetMin())

}
