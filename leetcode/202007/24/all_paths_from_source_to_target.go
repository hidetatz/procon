package main

var (
	target int
	graph  [][]int
	memo   map[int][][]int
)

func allPathsSourceTarget(g [][]int) [][]int {
	target = len(g) - 1
	graph = g
	memo = map[int][][]int{}

	return dp(0)
}

func dp(cur int) [][]int {
	if nod, ok := memo[cur]; ok {
		return nod
	}

	results := [][]int{}

	if cur == target {
		path := []int{target}
		results = append(results, path)
		return results
	}

	for _, next := range graph[cur] {
		for _, path := range dp(next) {
			newPath := []int{cur}
			newPath = append(newPath, path...)
			results = append(results, newPath)
		}
	}

	memo[cur] = results

	return results
}
