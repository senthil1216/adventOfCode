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

func createFoldMethods(line string) FoldMethods {
	fold := strings.Split(line, " ")
	dir := fold[len(fold)-1]
	vals := strings.Split(dir, "=")
	var f FoldMethods
	f.Dir = vals[0]
	f.Num, _ = strconv.Atoi(vals[1])
	return f
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
	var foldStarted bool
	var folds []FoldMethods
	var coords []Coords
	var maxColIdx, maxRowIdx int
	for scanner.Scan() {
		curr := scanner.Text()
		if len(strings.TrimSpace(curr)) == 0 {
			foldStarted = true
			continue
		}
		if !foldStarted {
			var coord Coords
			line := strings.Split(curr, ",")
			coord.ColIdx, _ = strconv.Atoi(line[0])
			coord.RowIdx, _ = strconv.Atoi(line[1])
			maxColIdx = max(maxColIdx, coord.ColIdx)
			maxRowIdx = max(maxRowIdx, coord.RowIdx)
			coords = append(coords, coord)
		} else {
			fold := createFoldMethods(curr)
			folds = append(folds, fold)
		}
	}
	var newCoords []Coords
	for i := 0; i < len(folds); i++ {
		currFold := folds[i]
		if currFold.Dir == "y" {
			diff := maxRowIdx - currFold.Num
			for _, coord := range coords {
				var newCoord Coords
				newCoord.ColIdx = coord.ColIdx
				if coord.RowIdx < currFold.Num {
					newCoord.RowIdx = coord.RowIdx
				} else {
					newCoord.RowIdx = abs(coord.RowIdx - (diff * 2))
				}
				newCoords = append(newCoords, newCoord)
			}
		} else {
			diff := maxColIdx - currFold.Num
			for _, coord := range coords {
				var newCoord Coords
				newCoord.RowIdx = coord.RowIdx
				if coord.ColIdx < currFold.Num {
					newCoord.ColIdx = coord.ColIdx
				} else {
					newCoord.ColIdx = abs(coord.ColIdx - (diff * 2))
				}
				newCoords = append(newCoords, newCoord)
			}
		}
	}
	sort.SliceStable(newCoords, func(i, j int) bool {
		if newCoords[i].RowIdx < newCoords[j].RowIdx {
			return true
		} else if newCoords[i].RowIdx == newCoords[j].RowIdx {
			return newCoords[i].ColIdx <= newCoords[j].ColIdx
		}
		return false
	})
	var cnt int = 1
	prev := newCoords[0]
	fmt.Println(maxColIdx, maxRowIdx)
	maxColIdx = 0
	maxRowIdx = 0
	for i := 1; i < len(newCoords); i++ {
		curr := newCoords[i]
		if curr.ColIdx != prev.ColIdx || curr.RowIdx != prev.RowIdx {
			cnt++
		}
		maxColIdx = max(maxColIdx, curr.ColIdx)
		maxRowIdx = max(maxRowIdx, curr.RowIdx)
		prev = curr
	}
	fmt.Println(maxColIdx, maxRowIdx)
	fmt.Println(len(newCoords))
	var line string
	// for i := 0; i < maxRowIdx; i++ {
	// 	for j := 0; j < maxColIdx; j++ {
	// 		line += "."
	// 	}
	// 	line += "\n"
	// }
	fmt.Println(line)
}
