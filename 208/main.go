package main

import "fmt"

type Trie struct {
	Val rune

	Next []*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	ret := Trie{}
	ret.Next = make([]*Trie, 27)
	return ret
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		this.Next[26] = new(Trie)
		return
	}
	next := int(word[0]) - int('a')
	if this.Next[next] == nil {
		this.Next[next] = new(Trie)
		this.Next[next].Val = rune(word[0])
		this.Next[next].Next = make([]*Trie, 27)
	}
	this.Next[next].Insert(word[1:])
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		if this.Next[26] == nil {
			return false
		} else {
			return true
		}
	}
	next := int(word[0]) - int('a')
	if this.Next[next] == nil {
		return false
	}
	return this.Next[next].Search(word[1:])
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	next := int(prefix[0]) - int('a')
	if this.Next[next] == nil {
		return false
	}
	return this.Next[next].StartsWith(prefix[1:])
}

func main() {
	trie := Constructor()

	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // returns true
	fmt.Println(trie.Search("app"))     // returns false
	fmt.Println(trie.StartsWith("app")) // returns true
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // returns true
}
