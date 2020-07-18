package main

var (
	visited map[int]string
	graph   map[int][]int
	ret     []int
	cycle   bool
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	visited = make(map[int]string)
	graph = make(map[int][]int)
	ret = make([]int, 0)
	cycle = false

	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
	}

	for i := 0; i < numCourses; i++ {
		visited[i] = "unvisited"
	}

	for n := 0; n < numCourses; n++ {
		if visited[n] == "unvisited" {
			dfs(n)
		}
	}

	if cycle {
		return []int{}
	}

	return ret
}

func dfs(n int) {
	if visited[n] == "wip" {
		cycle = true
		return
	} else if visited[n] == "unvisited" {
		visited[n] = "wip"
		for _, v := range graph[n] {
			dfs(v)
		}
		visited[n] = "finished"
		ret = append([]int{n}, ret...)
	}
}
