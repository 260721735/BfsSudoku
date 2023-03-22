package main

import (
	"log"
	"time"
)
func main(){
	inputBoard := [9][9]int{
		{0, 0, 0, 0, 0, 9, 0, 0, 8},
		{0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 3, 0, 0, 0, 9, 6, 5},
		{0, 0, 0, 2, 7, 0, 0, 0, 0},
		{8, 0, 0, 0, 9, 0, 0, 0, 0},
		{0, 6, 0, 0, 0, 0, 0, 0, 4},
		{9, 0, 0, 5, 1, 0, 0, 8, 0},
		{6, 3, 0, 0, 0, 0, 0, 0, 0},
		{0, 8, 7, 0, 0, 6, 0, 9, 3},
	}

	log.Println("xyz-win输入")
	for i:=range inputBoard{
		log.Println(inputBoard[i])
	}
	start1 := time.Now()
	xyzs:=xyz_win(inputBoard)
	end1 := time.Now()
	log.Println("xyz-win结果")
	for i:=range xyzs{
		log.Println(xyzs[i])
	}
	log.Println("xyz_win耗时", end1.Sub(start1))
}