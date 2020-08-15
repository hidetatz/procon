package main

func findPermutation(s string) []int {
	stack := []int{}
	push := func(i int) {
		stack = append(stack, i)
	}

	pop := func() int {
		if len(stack) == 0 {
			panic("empty")
		}
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1] // remove last
		return last
	}

	ret := []int{}

	for i := 1; i <= len(s); i++ {
		push(i)
		if s[i-1] == 'I' {
			for len(stack) != 0 {
				ret = append(ret, pop())
			}
		}
	}

	push(len(s) + 1)
	for len(stack) != 0 {
		ret = append(ret, pop())
	}
	return ret
}
