package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func main() {
	//file, _ := os.Open("/Users/senthil.sivasubraman/repo/adventOfCode/2021/day7/input_back")
	file, _ := os.Open("/Users/senthil.sivasubraman/repo/adventOfCode/2021/day7/input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	lines := strings.Split(scanner.Text(), ",")
	crabPosition := make(map[int]int)
	for i := 0; i < len(lines); i++ {
		pos, _ := strconv.Atoi(lines[i])
		val, _ := crabPosition[pos]
		crabPosition[pos] = val + 1
	}
	minCost := math.MaxFloat64
	for pointer, _ := range crabPosition {
		var cost int
		for k, v := range crabPosition {
			cost += abs(k-pointer) * v
		}
		minCost = math.Min(minCost, float64(cost))
	}
	fmt.Println(minCost)
}
