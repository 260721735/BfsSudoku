package main

func bfs(shudoku [9][9]int) [9][9]int {
	var uncertainty, mapk = initUncertainty(shudoku)
	simplecheck(&shudoku, &uncertainty, &mapk)
	for i := range mapk {
		if len(uncertainty[mapk[i]]) == 1 {
			shudoku[mapk[i]/9][mapk[i]%9] = uncertainty[mapk[i]][0]
		}
	}
	return shudoku
}
func initUncertainty(shudoku [9][9]int) (map[int][]int, []int) {
	var uncertainty = make(map[int][]int)
	var mapk = []int{}
	//记录有哪些未定位置
	for i := range shudoku {
		for j := range shudoku[i] {
			if shudoku[i][j] == 0 {
				uncertainty[9*i+j] = simpleCheckMay(shudoku, i, j)
				mapk = append(mapk, 9*i+j)
			}
		}
	}
	//计算可能性
	return uncertainty, mapk
}
func simpleCheckMay(shudoku [9][9]int, nowi int, nowj int) []int {
	var may = [9]int{}
	var starti = (nowi / 3) * 3
	var startj = (nowj / 3) * 3
	//判断九宫格里存在的可能性
	for i := starti; i < starti+3; i++ {
		for j := startj; j < startj+3; j++ {
			if shudoku[i][j] != 0 {
				may[shudoku[i][j]-1] = 1
			}
		}
	}
	//判断当前行存在的可能性
	for i := 0; i < 9; i++ {
		if i == nowi {
			continue
		}
		if shudoku[i][nowj] != 0 {
			may[shudoku[i][nowj]-1] = 1
		}
	}
	//判断当前列存在的可能性
	for j := 0; j < 9; j++ {
		if j == nowj {
			continue
		}
		if shudoku[nowi][j] != 0 {
			may[shudoku[nowi][j]-1] = 1
		}
	}
	var result = []int{}
	for i := range may {
		if may[i] == 0 {
			result = append(result, i+1)
		}
	}
	return result
}
func simplecheck(shudoku *[9][9]int, publicuncertainty *map[int][]int, publicmapk *[]int) {
	var uncertainty = *publicuncertainty
	var mapk = *publicmapk
	//fmt.Println(uncertainty[0])
	for index := 0; index < len(uncertainty); index++ {
		if len(uncertainty[mapk[index]]) == 1 {
			i := mapk[index] / 9
			j := mapk[index] % 9
			shudoku[i][j] = uncertainty[mapk[index]][0]
			delete(uncertainty, mapk[index])
			mapk = append(mapk[:index], mapk[index+1:]...)
			fixMay(*shudoku, i, j, publicuncertainty)
			index = 0
		}
	}
	//log.Println("length",len(uncertainty))
	//	for i:=range uncertainty{
	//		log.Println(i,uncertainty[i])
	//	}
}
func fixMay(shudoku [9][9]int, fixi int, fixj int, publicuncertainty *map[int][]int) {
	var uncertainty = *publicuncertainty
	//修复当前九宫格内的uncertainty map列表
	var starti = fixi / 3
	var startj = fixj / 3
	//判断九宫格里存在的可能性
	for i := starti; i < starti+3; i++ {
		for j := startj; j < startj+3; j++ {
			if shudoku[i][j] == 0 {
				for index := range uncertainty[9*i+j] {
					if uncertainty[9*i+j][index] == shudoku[fixi][fixj] {
						uncertainty[9*i+j] = append(uncertainty[9*i+j][:index], uncertainty[9*i+j][index+1:]...)
						break
					}
				}
			}
		}
	}
	//修复当前行的uncertainty map列表
	for i := 0; i < 9; i++ {
		if i == fixi {
			continue
		}
		if shudoku[i][fixj] == 0 {
			for index := range uncertainty[9*i+fixj] {
				if uncertainty[9*i+fixj][index] == shudoku[fixi][fixj] {
					uncertainty[9*i+fixj] = append(uncertainty[9*i+fixj][:index], uncertainty[9*i+fixj][index+1:]...)
					break
				}
			}
		}
	}
	//修复当前列的uncertainty map列表
	for j := 0; j < 9; j++ {
		if j == fixj {
			continue
		}
		if shudoku[fixi][j] == 0 {
			for index := range uncertainty[9*fixi+j] {
				if uncertainty[9*fixi+j][index] == shudoku[fixi][fixj] {
					uncertainty[9*fixi+j] = append(uncertainty[9*fixi+j][:index], uncertainty[9*fixi+j][index+1:]...)
					break
				}
			}
		}
	}
}
