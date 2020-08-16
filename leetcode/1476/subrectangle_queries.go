package main

type SubrectangleQueries struct {
	matrix [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{matrix: rectangle}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for i := range this.matrix {
		for j := range this.matrix[i] {
			if row1 <= i && i <= row2 && col1 <= j && j <= col2 {
				this.matrix[i][j] = newValue
			}
		}
	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.matrix[row][col]
}

/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */
