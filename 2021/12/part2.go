package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"unicode"
)

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func makeNeighbour(graph *map[string][]string, dest1 string, dest2 string) {
	paths := make([]string, 0)
	paths = (*graph)[dest1]
	paths = append(paths, dest2)
	(*graph)[dest1] = paths
	paths = make([]string, 0)
	var paths2 []string
	paths2 = (*graph)[dest2]
	paths2 = append(paths2, dest1)
	(*graph)[dest2] = paths2
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
	graph := make(map[string][]string)
	var count int
	for scanner.Scan() {
		lines := scanner.Text()
		conns := strings.Split(lines, "-")
		dest1 := conns[0]
		dest2 := conns[1]
		makeNeighbour(&graph, dest1, dest2)
	}
	var smallCaves []string
	count = findExit(graph, smallCaves, "start")
	fmt.Println(count)
}
func checkSmallCave(cave string, smallCaves []string, visitOneSmallCaveTwice bool) bool {
	numberOfSmallCavesVisits := make(map[string]int)
	for i := range smallCaves {
		numberOfSmallCavesVisits[smallCaves[i]]++
	}
	if numberOfSmallCavesVisits[cave] >= 1 {
		if !visitOneSmallCaveTwice {
			return false
		}
		for _, i := range numberOfSmallCavesVisits {
			if i >= 2 {
				return false
			}
		}
	}
	return true
}
func findExit(caves map[string][]string, smallCaves []string, nodeName string) int {
	if nodeName == "end" {
		return 1
	}
	if nodeName == "start" && len(smallCaves) != 0 {
		return 0
	}
	if strings.ToLower(nodeName) == nodeName {
		if smallCaves == nil {
			smallCaves = make([]string, 0)
		}
		if !checkSmallCave(nodeName, smallCaves, true) {
			return 0
		}
		smallCaves = append(smallCaves, nodeName)
	}
	var count int
	for _, p := range caves[nodeName] {
		count += findExit(caves, smallCaves, p)
	}
	return count
}
