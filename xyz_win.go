package main


var board [9][9]int
var deep int
var withXYZWIN bool
var tryTime int

func init() {
	deep = 0
	tryTime = 0
	withXYZWIN = true
}
func xyz_win(shudoku [9][9]int)[9][9]int{
	board=shudoku
	solveSudoku(0)
	return board

}
// func main() {
// 	inputBoard := [9][9]int{
// 		{0, 0, 0, 0, 0, 9, 0, 0, 8},
// 		{0, 0, 0, 0, 0, 0, 0, 1, 0},
// 		{0, 0, 3, 0, 0, 0, 9, 6, 5},
// 		{0, 0, 0, 2, 7, 0, 0, 0, 0},
// 		{8, 0, 0, 0, 9, 0, 0, 0, 0},
// 		{0, 6, 0, 0, 0, 0, 0, 0, 4},
// 		{9, 0, 0, 5, 1, 0, 0, 8, 0},
// 		{6, 3, 0, 0, 0, 0, 0, 0, 0},
// 		{0, 8, 7, 0, 0, 6, 0, 9, 3},
// 	}
// 	board = inputBoard
// 	start := time.Now()
// 	log.Println(solveSudoku(0))
// 	end := time.Now()
// 	log.Println("deep层", deep)
// 	log.Println("耗时", end.Sub(start))
// 	log.Println("尝试次数", tryTime) //296805
// 	for i := 0; i < 9; i++ {
// 		log.Println(board[i])
// 	}
// }

// 回溯
func solveSudoku(innderDeep int) bool {
	if deep < innderDeep {
		deep = innderDeep
	}
	// 查找一个未赋值的单元格
	row, col := findEmptyCell()
	if row == -1 && col == -1 {
		// 如果不存在未赋值的单元格，说明数独已经被解决
		return true
	}

	// 尝试填入1-9中的一个数字
	for num := 1; num <= 9; num++ {
		if isValid(row, col, num) {
			board[row][col] = num
			tryTime++
			if solveSudoku(innderDeep + 1) {
				return true
			}
			board[row][col] = 0
		}
	}

	return false
}

func findEmptyCell() (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func isValid(row, col, num int) bool {
	// 检查同一行中是否有重复数字
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
	}

	// 检查同一列中是否有重复数字
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}

	// 检查同一个九宫格中是否有重复数字
	gridRow := row / 3 * 3
	gridCol := col / 3 * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[gridRow+i][gridCol+j] == num {
				return false
			}
		}
	}
	if !withXYZWIN {
		return true
	}
	// 剪枝
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if i == row && j == col {
				continue
			}
			if board[i][j] == 0 && isXYZWing(row, col, i, j, num) {
				return false
			}
		}
	}

	// 如果通过以上所有检查，则说明当前数字是有效的
	return true
}

func isXYZWing(row1, col1, row2, col2, num int) bool {
	// 如果行或列不同，则无法使用xyz-wing技巧进行剪枝
	if row1 != row2 && col1 != col2 {
		return false
	}

	// 如果两个单元格之间的距离超过2，则无法使用xyz-wing技巧进行剪枝
	if abs(row1-row2) > 2 || abs(col1-col2) > 2 {
		return false
	}

	// 如果两个单元格不在同一个九宫格内，则无法使用xyz
	gridRow1 := row1 / 3 * 3
	gridCol1 := col1 / 3 * 3
	gridRow2 := row2 / 3 * 3
	gridCol2 := col2 / 3 * 3
	if gridRow1 != gridRow2 || gridCol1 != gridCol2 {
		return false
	}

	// 确定第三个单元格的位置
	var row3, col3 int
	if row1 == row2 {
		// 如果两个单元格在同一行上，则第三个单元格的位置为同一列上除了这两个单元格之外的另一个单元格
		for i := 0; i < 9; i++ {
			if i == col1 || i == col2 {
				continue
			}
			if board[row1][i] == num && isXWing(col1, col2, i, num) {
				row3 = row1
				col3 = i
				break
			}
		}
	} else {
		// 如果两个单元格在同一列上，则第三个单元格的位置为同一行上除了这两个单元格之外的另一个单元格
		for i := 0; i < 9; i++ {
			if i == row1 || i == row2 {
				continue
			}
			if board[i][col1] == num && isXWing(row1, row2, i, num) {
				row3 = i
				col3 = col1
				break
			}
		}
	}

	// 如果没有找到第三个单元格，则无法使用xyz-wing技巧进行剪枝
	if row3 == 0 && col3 == 0 {
		return false
	}

	// 检查第三个单元格与第一个和第二个单元格的关系
	if row1 == row2 {
		// 如果两个单元格在同一行上，则第三个单元格必须在另外一行上
		if row3 == row1 || row3 == row2 {
			return false
		}
	} else {
		// 如果两个单元格在同一列上，则第三个单元格必须在另外一列上
		if col3 == col1 || col3 == col2 {
			return false
		}
	}

	// 检查第三个单元格是否与第一个单元格和第二个单元格之间存在相同的数字
	if board[row3][col3] == num {
		return false
	}

	// 检查第三个单元格和第一个单元格、第二个单元格之间是否存在X-Wing关系
	if row1 == row2 {
		return isXWing(col1, col2, col3, num)
	} else {
		return isXWing(row1, row2, row3, num)
	}
}

func isXWing(col1, col2, col3, num int) bool {
	// 检查第一个和第二个单元格是否存在相同的数字
	var rows []int
	for i := 0; i < 9; i++ {
		if board[i][col1] == num && board[i][col2] == num {
			rows = append(rows, i)
		}
	}

	// 如果第一个和第二个单元格没有共同的数字，则无法使用xyz-wing技巧进行剪枝
	if len(rows) != 2 {
		return false
	}

	// 检查第三个单元格与第一个和第二个单元格的关系
	for _, row := range rows {
		if board[row][col3] == num {
			return false
		}
	}

	// 检查第三个单元格所在的行中是否有另一个单元格与第一个和第二个单元格都存在相同的数字
	for i := 0; i < 9; i++ {
		if i == col1 || i == col2 || i == col3 {
			continue
		}
		if board[rows[0]][i] == num && board[rows[1]][i] == num {
			return true
		}
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
