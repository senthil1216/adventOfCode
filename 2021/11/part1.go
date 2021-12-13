package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
)

const steps int = 100
const rowLimit int = 10
const colLimit int = 10

func getDirs(r int, c int) [][]int {
	dirs := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},

		{1, -1},
		{1, 0},
		{1, 1},

		{0, 1},
		{0, -1},
	}
	var valid [][]int
	for i := 0; i < len(dirs); i++ {
		newRow := r + dirs[i][0]
		newCol := c + dirs[i][1]
		if newRow < 0 || newCol < 0 || newRow >= rowLimit || newCol >= colLimit {
			continue
		}
		newDir := []int{newRow, newCol}
		valid = append(valid, newDir)
	}
	return valid
}

func findFlash(grid *[rowLimit][colLimit]int, row int, col int, visited *[rowLimit][colLimit]bool) int {
	if (*visited)[row][col] {
		return 0
	}

	(*visited)[row][col] = true
	count := 1
	dirs := getDirs(row, col)
	for _, dirs := range dirs {
		adjRowIdx := dirs[0]
		adjColIdx := dirs[1]
		if (*visited)[adjRowIdx][adjColIdx] {
			continue
		}
		grid[adjRowIdx][adjColIdx]++
		if grid[adjRowIdx][adjColIdx] > 9 {
			count += findFlash(grid, adjRowIdx, adjColIdx, visited)
			grid[adjRowIdx][adjColIdx] = 0
		}
	}
	return count
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
	var grid [rowLimit][colLimit]int
	var r, c int
	for scanner.Scan() {
		rowVal := scanner.Text()
		c = 0
		for i := 0; i < len(rowVal); i++ {
			colVal := string(rowVal[i])
			num, _ := strconv.Atoi(colVal)
			grid[r][c] = num
			c++
		}
		r++
	}
	var s int
	var flashes int
	// printGrid(grid)
	// fmt.Println("-------------")
	for s < steps {
		// .. update the individual vals
		for r := 0; r < rowLimit; r++ {
			for c := 0; c < colLimit; c++ {
				grid[r][c] += 1
			}
		}
		var visited [rowLimit][colLimit]bool
		// .. check for 9s
		for r := 0; r < rowLimit; r++ {
			for c := 0; c < colLimit; c++ {
				if grid[r][c] > 9 {
					flashes += findFlash(&grid, r, c, &visited)
					grid[r][c] = 0
				}
			}
		}
		// printGrid(grid)
		// fmt.Println(s + 1)
		// fmt.Println("---------------")
		s++
	}
	fmt.Println(flashes)
}

func printGrid(grid [rowLimit][colLimit]int) {
	var print string
	for i := 0; i < rowLimit; i++ {
		for j := 0; j < colLimit; j++ {
			print += " " + strconv.Itoa(grid[i][j])
		}
		print += "\n"
	}
	fmt.Println(print)
}
