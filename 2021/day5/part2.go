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

// readLines reads a whole file into memory and returns a slice of its lines.
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

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func main() {
	//lines, _ := readLines("./input_back")
	lines, _ := readLines("./input")
	var grid [1000][1000]int
	for i := 0; i < len(lines); i++ {
		coords := strings.Split(lines[i], "->")
		vals := strings.Split(strings.TrimSpace(coords[0]), ",")
		x1, _ := strconv.Atoi(vals[0])
		y1, _ := strconv.Atoi(vals[1])

		vals = strings.Split(strings.TrimSpace(coords[1]), ",")
		x2, _ := strconv.Atoi(vals[0])
		y2, _ := strconv.Atoi(vals[1])

		// .. horizontal lines
		if y1 == y2 {
			for x := min(x1, x2); x <= max(x1, x2); x++ {
				grid[y1][x] = grid[y1][x] + 1
			}
		}

		// .. vertical lines
		if x1 == x2 {
			for y := min(y1, y2); y <= max(y1, y2); y++ {
				grid[y][x1] = grid[y][x1] + 1
			}
		}

		// .. diagonals
		if !(x1 == x2 || y1 == y2) {
			deltaX := 1
			if x1 > x2 {
				deltaX = -1
			}
			deltaY := 1
			if y1 > y2 {
				deltaY = -1
			}
			grid[y1][x1]++
			for x1 != x2 {
				x1 += deltaX
				y1 += deltaY
				grid[y1][x1] = grid[y1][x1] + 1
			}
		}
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] >= 2 {
				count++
			}
		}
	}
	fmt.Println(count)
}
