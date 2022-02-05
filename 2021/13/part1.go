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

type Coords struct {
	RowIdx int
	ColIdx int
}

type FoldMethods struct {
	Dir string
	Num int
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func parseFoldMethods(line string) FoldMethods {
	fold := strings.Split(line, " ")
	dir := fold[len(fold)-1]
	vals := strings.Split(dir, "=")
	var f FoldMethods
	f.Dir = vals[0]
	f.Num, _ = strconv.Atoi(vals[1])
	return f
}

func parseCoords(currLine string) Coords {
	var coord Coords
	line := strings.Split(currLine, ",")
	coord.ColIdx, _ = strconv.Atoi(line[0])
	coord.RowIdx, _ = strconv.Atoi(line[1])
	return coord
}

func findMaxCoords(gridMap map[Coords]bool) (int, int) {
	var maxRowIdx, maxColIdx int
	for key, _ := range gridMap {
		maxRowIdx = max(key.RowIdx, maxRowIdx)
		maxColIdx = max(key.ColIdx, maxColIdx)
	}
	return maxRowIdx, maxColIdx
}

func doFold(gridMap map[Coords]bool, foldSize int, folds []FoldMethods) map[Coords]bool {
	var tempGrid map[Coords]bool
	for i := 0; i < foldSize; i++ {
		tempGrid = make(map[Coords]bool)
		currFold := folds[i]
		if currFold.Dir == "y" { // Fold alongside the row
			for coord, _ := range gridMap {
				var newCoord Coords
				newCoord.ColIdx = coord.ColIdx
				if coord.RowIdx < currFold.Num {
					newCoord.RowIdx = coord.RowIdx
				} else {
					newCoord.RowIdx = 2*currFold.Num - coord.RowIdx
				}
				tempGrid[newCoord] = true
			}
		} else if currFold.Dir == "x" { // .. Fold alongside the column
			for coord, _ := range gridMap {
				var newCoord Coords
				newCoord.RowIdx = coord.RowIdx
				if coord.ColIdx < currFold.Num {
					newCoord.ColIdx = coord.ColIdx
				} else {
					newCoord.ColIdx = 2*currFold.Num - coord.ColIdx
				}
				tempGrid[newCoord] = true
			}
		}
		gridMap = tempGrid
	}
	return gridMap
}

func solvePart1(gridMap map[Coords]bool, foldMethods []FoldMethods) {
	gridMap = doFold(gridMap, 1, foldMethods)
	fmt.Printf("Part 1: \n%v\n", len(gridMap))
}

func solvePart2(gridMap map[Coords]bool, foldMethods []FoldMethods) {
	gridMap = doFold(gridMap, len(foldMethods), foldMethods)
	maxRowIdx, maxColIdx := findMaxCoords(gridMap)
	var out string
	for r := 0; r <= maxRowIdx; r++ {
		for c := 0; c <= maxColIdx; c++ {
			findCoords := Coords{
				RowIdx: r,
				ColIdx: c,
			}
			if gridMap[findCoords] {
				out += "#"
			} else {
				out += "."
			}
		}
		out += "\n"
	}
	fmt.Printf("Part 2: \n%v\n", out)
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
	var foldStarted bool
	var folds []FoldMethods
	gridMap := make(map[Coords]bool)
	for scanner.Scan() {
		curr := scanner.Text()
		if len(strings.TrimSpace(curr)) == 0 {
			foldStarted = true
			continue
		}
		if !foldStarted {
			coord := parseCoords(curr)
			gridMap[coord] = true
		} else {
			fold := parseFoldMethods(curr)
			folds = append(folds, fold)
		}
	}
	solvePart1(gridMap, folds)
	solvePart2(gridMap, folds)
}
