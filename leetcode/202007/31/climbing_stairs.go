package main

func climbStairs(n int) int {
	return climbStairs_recursive(n)
	// return climbStairs_for(n)
}

func climbStairs_recursive(n int) int {
	if n <= 3 {
		return n
	}
	f, s := 2, 3
	for i := 4; i < n; i++ {
		f, s = s, f+s
	}
	return f + s
}

func climbStairs_for(n int) int {
	return helper(n, map[int]int{})
}

func helper(n int, memo map[int]int) int {
	if n <= 2 {
		return n
	}
	if v, ok := memo[n]; ok {
		return v
	}
	memo[n] = helper(n-2, memo) + helper(n-1, memo)
	return memo[n]
}
