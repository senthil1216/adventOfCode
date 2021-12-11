package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

const rowLimit int = 100
const colLimit int = 100

func getDirs(row, col int) [][]int {
	dirs := make([][]int, 4)
	if row-1 >= 0 {
		dirs[0] = make([]int, 2)
		dirs[0][0] = row - 1
		dirs[0][1] = col
	}

	if row+1 <= rowLimit-1 {
		dirs[1] = make([]int, 2)
		dirs[1][0] = row + 1
		dirs[1][1] = col
	}

	if col-1 >= 0 {
		dirs[2] = make([]int, 2)
		dirs[2][0] = row
		dirs[2][1] = col - 1
	}

	if col+1 <= colLimit-1 {
		dirs[3] = make([]int, 2)
		dirs[3][0] = row
		dirs[3][1] = col + 1
	}
	return dirs
}

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	dirName := path.Dir(filename)
	file, _ := os.Open(dirName + "/input")
	// file, _ := os.Open(dirName + "/input_back")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	rowCounter := 0
	var grid [rowLimit][colLimit]string
	for scanner.Scan() {
		line := scanner.Text()
		for i, l := range line {
			grid[rowCounter][i] = strings.TrimSpace(string(l))
		}
		rowCounter++
	}
	var lowPoint int
	for row := 0; row < rowLimit; row++ {
		for col := 0; col < colLimit; col++ {
			isLow := true
			gridVal, _ := strconv.Atoi(grid[row][col])
			for _, vals := range getDirs(row, col) {
				if len(vals) == 0 {
					continue
				}
				neighVal, _ := strconv.Atoi(grid[vals[0]][vals[1]])
				if neighVal <= gridVal {
					isLow = false
				}
			}
			if isLow {
				lowPoint += gridVal + 1
			}
		}
	}
	fmt.Println(lowPoint)
}
