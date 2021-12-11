package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

const rowLimit int = 100
const colLimit int = 100

//const rowLimit int = 5
//const colLimit int = 10

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

type Point struct {
	Row int
	Col int
}

func findLowPoints(grid [rowLimit][colLimit]int) []*Point {
	var lowPoints []*Point
	for row := 0; row < rowLimit; row++ {
		for col := 0; col < colLimit; col++ {
			isLow := true
			for _, vals := range getDirs(row, col) {
				if len(vals) == 0 {
					continue
				}
				if grid[vals[0]][vals[1]] <= grid[row][col] {
					isLow = false
				}
			}
			if isLow {
				pointer := new(Point)
				pointer.Row = row
				pointer.Col = col
				lowPoints = append(lowPoints, pointer)
			}
		}
	}
	return lowPoints
}

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	dirName := path.Dir(filename)
	file, _ := os.Open(dirName + "/input")
	//file, _ := os.Open(dirName + "/input_back")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	rowCounter := 0
	var grid [rowLimit][colLimit]int
	for scanner.Scan() {
		line := scanner.Text()
		for i, l := range line {
			gridVal, _ := strconv.Atoi(strings.TrimSpace(string(l)))
			grid[rowCounter][i] = gridVal
		}
		rowCounter++
	}
	lowPoints := findLowPoints(grid)
	var visited [rowLimit][colLimit]bool
	var basins []int
	for _, lowPoint := range lowPoints {
		if visited[lowPoint.Row][lowPoint.Col] {
			continue
		}
		basinCount := findBasin(grid, lowPoint.Row, lowPoint.Col, &visited)
		basins = append(basins, basinCount)
	}
	sort.Ints(basins)
	c := len(basins)
	fmt.Println(basins[c-1] * basins[c-2] * basins[c-3])
}

func findBasin(grid [rowLimit][colLimit]int, row int, col int, visited *[rowLimit][colLimit]bool) int {
	if (*visited)[row][col] {
		return 0
	}

	(*visited)[row][col] = true
	count := 1
	dirs := getDirs(row, col)
	for _, dirs := range dirs {
		if len(dirs) == 0 {
			continue
		}
		adjRowIdx := dirs[0]
		adjColIdx := dirs[1]
		if grid[adjRowIdx][adjColIdx] == 9 {
			continue
		}
		if grid[adjRowIdx][adjColIdx] > grid[row][col] {
			count += findBasin(grid, adjRowIdx, adjColIdx, visited)
		}
	}
	return count
}
