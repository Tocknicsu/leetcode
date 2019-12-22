package main

import "fmt"

type WordDictionary struct {
	Next []*WordDictionary
}

func Trans(x rune) int {
	if x >= rune('a') && x <= rune('z') {
		return int(x - rune('a'))
	}
	return int(x - rune('A'))
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	ret := WordDictionary{}
	ret.Next = make([]*WordDictionary, 53)
	return ret
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	if len(word) == 0 {
		this.Next[52] = new(WordDictionary)
		return
	}
	next := Trans(rune(word[0]))
	if this.Next[next] == nil {
		this.Next[next] = new(WordDictionary)
		this.Next[next].Next = make([]*WordDictionary, 53)
	}
	this.Next[next].AddWord(word[1:])
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	if this == nil {
		return false
	}
	if len(word) == 0 {
		return this.Next[52] != nil
	}
	if word[0] == '.' {
		for i := 0; i < 52; i = i + 1 {
			if this.Next[i].Search(word[1:]) {
				return true
			}
		}
		return false
	} else {
		next := Trans(rune(word[0]))
		return this.Next[next].Search(word[1:])
	}
}

func main() {
	obj := Constructor()
	obj.AddWord("bad")
	obj.AddWord("dad")
	obj.AddWord("mad")
	fmt.Println(obj.Search("pad"))
	fmt.Println(obj.Search("bad"))
	fmt.Println(obj.Search(".ad"))
	fmt.Println(obj.Search("b.."))
}
