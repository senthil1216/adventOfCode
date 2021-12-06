package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	dim [5][5]int
}

func (currBoard *Board) mark(bingoNum int) {
	for i := 0; i < len(currBoard.dim); i++ {
		for j := 0; j < len(currBoard.dim[0]); j++ {
			if currBoard.dim[i][j] == bingoNum {
				currBoard.dim[i][j] = -1
			}
		}
	}
}

func (currBoard *Board) checkWin() bool {

	// .. check all cols
	for i := 0; i < len(currBoard.dim[0]); i++ {
		isWon := true
		for j := 0; j < len(currBoard.dim); j++ {
			if currBoard.dim[j][i] < 0 {
				isWon = isWon && true
			} else {
				isWon = isWon && false
			}
		}
		if isWon {
			return true
		}
	}

	// .. check all rows
	for i := 0; i < len(currBoard.dim); i++ {
		isWon := true
		for j := 0; j < len(currBoard.dim[0]); j++ {
			if currBoard.dim[i][j] < 0 {
				isWon = isWon && true
			} else {
				isWon = isWon && false
			}
		}
		if isWon {
			return true
		}
	}
	return false
}

func (currBoard *Board) calcScore(num int) int {
	sum := 0
	for i := 0; i < len(currBoard.dim); i++ {
		for j := 0; j < len(currBoard.dim[0]); j++ {
			if currBoard.dim[i][j] < 0 {
				continue
			}
			sum += currBoard.dim[i][j]
		}
	}
	return sum * num
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

func main() {
	lines, _ := readLines("./input")
	currBoard := new(Board)
	var boards []*Board
	rowIdx := 0
	for i := 2; i < len(lines); i++ {
		currLine := lines[i]
		if len(strings.TrimSpace(currLine)) == 0 {
			boards = append(boards, currBoard)
			currBoard = new(Board)
			rowIdx = 0
		} else {
			nums := strings.Split(strings.TrimSpace(currLine), " ")
			colIdx := 0
			for _, val := range nums {
				if len(strings.TrimSpace(val)) == 0 {
					continue
				}
				numVal, _ := strconv.Atoi(val)
				currBoard.dim[rowIdx][colIdx] = numVal
				colIdx++
			}
			rowIdx++
		}
	}
	var score int
	boards = append(boards, currBoard)
	var found bool
	for _, val := range strings.Split(lines[0], ",") {
		if found {
			break
		}
		currNm, _ := strconv.Atoi(val)
		for i := 0; i < len(boards); i++ {
			currBoard := boards[i]
			currBoard.mark(currNm)
			if currBoard.checkWin() {
				score = currBoard.calcScore(currNm)
				found = true
			}
		}
	}
	fmt.Println(score)
}
