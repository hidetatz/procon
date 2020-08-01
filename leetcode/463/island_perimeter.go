package main

import "fmt"

func islandPerimeter(grid [][]int) int {
	perimeter := 0
	for i, line := range grid {
		for j, cell := range line {
			if cell == 0 {
				continue
			}

			perimeter += 4

			// check right and left cells
			if j-1 >= 0 {
				if grid[i][j-1] == 1 {
					perimeter--
				}
			}

			if j+1 < len(line) {
				if grid[i][j+1] == 1 {
					perimeter--
				}
			}

			// check above and below cells
			if i-1 >= 0 {
				if grid[i-1][j] == 1 {
					perimeter--
				}
			}

			if i+1 < len(grid) {
				if grid[i+1][j] == 1 {
					perimeter--
				}
			}
		}
	}

	return perimeter
}

func main() {
	fmt.Println(islandPerimeter([][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}))
}
