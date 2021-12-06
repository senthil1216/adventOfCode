package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
func swap(x int, y int) (int, int) {
	temp := x
	x = y
	return temp, y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func main() {
	//lines, _ := readLines("./input_back")
	lines, _ := readLines("./input")
	var lineObjs []*Line
	var diagonals []*Line
	for i := 0; i < len(lines); i++ {
		coords := strings.Split(lines[i], "->")
		line := new(Line)

		vals := strings.Split(strings.TrimSpace(coords[0]), ",")
		line.X1, _ = strconv.Atoi(vals[0])
		line.Y1, _ = strconv.Atoi(vals[1])

		vals = strings.Split(strings.TrimSpace(coords[1]), ",")
		line.X2, _ = strconv.Atoi(vals[0])
		line.Y2, _ = strconv.Atoi(vals[1])
		dy := line.Y2 - line.Y1
		dx := line.X2 - line.X1
		if abs(dy) == abs(dx) {
			diagonals = append(diagonals, line)
		} else if line.X1 == line.X2 || line.Y1 == line.Y2 {
			lineObjs = append(lineObjs, line)
		}
	}
	var grid [1000][1000]int
	count := 0
	for i := 0; i < len(lineObjs); i++ {
		currLine := lineObjs[i]
		for y := min(currLine.Y1, currLine.Y2); y <= max(currLine.Y1, currLine.Y2); y++ {
			for x := min(currLine.X1, currLine.X2); x <= max(currLine.X1, currLine.X2); x++ {
				grid[y][x]++
			}
		}
	}

	for i := 0; i < len(diagonals); i++ {
		currLine := diagonals[i]
		deltaX := 1
		if currLine.X1 > currLine.X2 {
			deltaX = -1
		}
		deltaY := 1
		if currLine.Y1 > currLine.Y2 {
			deltaY = -1
		}
		X1 := currLine.X1
		Y1 := currLine.Y1
		for X1 != currLine.X2 && Y1 != currLine.Y2 {
			grid[X1][Y1] += 1
			X1 += deltaX
			Y1 += deltaY

		}
		grid[X1][Y1]++
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] > 1 {
				count++
			}
		}
	}
	fmt.Println(count)
}
