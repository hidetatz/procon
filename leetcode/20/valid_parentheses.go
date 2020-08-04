package main

var open = map[string]struct{}{
	"(": struct{}{},
	"[": struct{}{},
	"{": struct{}{},
}

var close = map[string]struct{}{
	")": struct{}{},
	"]": struct{}{},
	"}": struct{}{},
}

type Queue struct {
	arr []string
}

func (q *Queue) Push(s string) {
	q.arr = append(q.arr, s)
}

func (q *Queue) Pop() string {
	if len(q.arr) == 0 {
		return ""
	}

	s := q.arr[len(q.arr)-1]
	q.arr = q.arr[:len(q.arr)-1]
	return s
}

func (q *Queue) Empty() bool {
	return len(q.arr) == 0
}

func compatible(l, r string) bool {
	if _, ok := open[l]; !ok {
		return false
	}

	if _, ok := close[r]; !ok {
		return false
	}

	return l == "(" && r == ")" || l == "[" && r == "]" || l == "{" && r == "}"
}

func isValid(s string) bool {
	if s == "" {
		return true
	}

	queue := &Queue{}

	for _, c := range s {
		if _, ok := open[string(c)]; ok {
			queue.Push(string(c))
			continue
		}

		p := queue.Pop()
		if !compatible(p, string(c)) {
			return false
		}
	}

	return queue.Empty()
}
