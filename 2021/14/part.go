package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path"
	"runtime"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
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

func parseInputAsPairsCount(input string) map[string]int {
	pairsCount := make(map[string]int)
	prev := strings.TrimSpace(string(input[0]))
	for i := 1; i < len(input); i++ {
		curr := strings.TrimSpace(string(input[i]))
		pairsCount[strings.TrimSpace(prev+curr)] += 1
		prev = curr
	}
	return pairsCount
}

func parseInputTemplateMap(input string) (string, string) {
	split := strings.Split(input, "->")
	return strings.TrimSpace(split[0]), strings.TrimSpace(split[1])
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
	var insertionPairs bool
	templateMap := make(map[string]string)
	var pairsCount map[string]int
	for scanner.Scan() {
		curr := scanner.Text()
		if len(strings.TrimSpace(curr)) == 0 {
			insertionPairs = true
			continue
		}
		if insertionPairs {
			key, val := parseInputTemplateMap(curr)
			templateMap[key] = val
		} else {
			pairsCount = parseInputAsPairsCount(curr)
		}
	}
	solvePart1(10, templateMap, pairsCount)
	solvePart1(40, templateMap, pairsCount)
}

func getCountOrDefault(pairs map[string]int, key string, defaultVal int) int {
	val, pres := pairs[key]
	if !pres {
		return defaultVal
	} else {
		return val
	}
}

func findCharCount(pairsCount map[string]int) map[string]int {
	charCount := make(map[string]int)
	for pairs, count := range pairsCount {
		charCount[string(pairs[0])] += count
		charCount[string(pairs[1])] += count
	}
	return charCount
}

func maxDiff(charCount map[string]int) int {
	var maxVal int = math.MinInt
	var minVal int = math.MaxInt
	for _, val := range charCount {
		maxVal = max(val, maxVal)
		minVal = min(val, minVal)
	}
	return (maxVal - minVal) / 2 // .. There is a double counting thats happening in the findCharCount map, thats why we need to divide by 2
}

func getPairsCount(pairsCount map[string]int, template map[string]string) map[string]int {
	tempPairs := make(map[string]int)
	for pair, count := range pairsCount {
		insertChar := template[pair]
		firstPair := string(pair[0]) + insertChar
		secondPair := insertChar + string(pair[1])
		tempPairs[firstPair] = getCountOrDefault(tempPairs, firstPair, 0) + count
		tempPairs[secondPair] = getCountOrDefault(tempPairs, secondPair, 0) + count
	}
	return tempPairs
}

func solvePart1(iterCount int, templateMap map[string]string, pairsCount map[string]int) {
	for i := 0; i < iterCount; i++ {
		pairsCount = getPairsCount(pairsCount, templateMap)
	}
	charCount := findCharCount(pairsCount)
	fmt.Println(maxDiff(charCount))
}
