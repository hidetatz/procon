package main

type WordDictionary struct {
	m map[int][]string
}

func Constructor() WordDictionary {
	return WordDictionary{m: map[int][]string{}}
}

func (this *WordDictionary) AddWord(word string) {
	l := len(word)
	if _, ok := this.m[l]; !ok {
		this.m[l] = []string{}
	}

	this.m[l] = append(this.m[l], word)
}

func (this *WordDictionary) Search(word string) bool {
	l := len(word)
	if arr, ok := this.m[l]; ok {
		for _, w := range arr {
			i := 0
			for (i < l) && (w[i] == word[i] || string(word[i]) == ".") {
				i++
			}

			if i == l {
				return true
			}
		}
	}
	return false
}
