package main

import (
	"math/bits"
	"sort"
)

type CombinationIterator struct {
	arr []string
	cur int
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	arr := []string{}
	l := len(characters)

	for i := 0; i < 1<<l; i++ {
		if !(bits.OnesCount(uint(i)) == combinationLength) {
			continue
		}
		s := ""
		for j := 0; j < l; j++ {
			if (i & (1 << j)) != 0 {
				s += string(characters[j])
			}
		}
		arr = append(arr, s)
	}

	sort.Strings(arr)
	return CombinationIterator{arr: arr, cur: 0}
}

func (this *CombinationIterator) Next() string {
	ret := this.arr[this.cur]
	this.cur++
	return ret
}

func (this *CombinationIterator) HasNext() bool {
	return this.cur != len(this.arr)
}

/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
