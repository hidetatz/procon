package main

type MyHashSet struct {
	table [1000000]bool
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{table: [1000000]bool{}}
}

func (this *MyHashSet) Add(key int) {
	this.table[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.table[key] = false
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.table[key] == true
}
