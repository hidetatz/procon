package main

import "fmt"

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}

	tasksMap := map[byte]int{}

	for _, task := range tasks {
		if cnt, ok := tasksMap[task]; ok {
			tasksMap[task] = cnt + 1
		} else {
			tasksMap[task] = 1
		}
	}

	fmt.Println(tasksMap)

	maxCnt := 0
	tasksCnt := 0
	for _, freq := range tasksMap {
		if maxCnt < freq {
			maxCnt = freq
			tasksCnt = 1
		} else if freq == maxCnt {
			tasksCnt++
		}
	}

	return max(len(tasks), (maxCnt-1)*(n+1)+tasksCnt)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
