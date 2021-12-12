package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"sort"
)

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
	costMap := map[string]int{
		"(": 1,
		"[": 2,
		"{": 3,
		"<": 4,
	}
	bracketMap := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
		"<": ">",
	}
	var incomplete []string
	var costs []int
	for scanner.Scan() {
		pattern := scanner.Text()
		var st []string
		loopBreak := false
		for i := 0; i < len(pattern); i++ {
			if loopBreak {
				break
			}
			bracket := string(pattern[i])
			val, prs := bracketMap[bracket]
			if prs {
				st = append(st, val)
			} else {
				if bracket == st[len(st)-1] {
					st = st[:len(st)-1]
				} else {
					loopBreak = true
				}
			}
		}
		if !loopBreak {
			incomplete = append(incomplete, pattern)
		}
	}
	i := 0
	for i < len(incomplete) {
		pattern := incomplete[i]
		var st []string
		j := 0
		for j < len(pattern) {
			bracket := string(pattern[j])
			_, prs := bracketMap[bracket]
			if prs {
				st = append(st, bracket)
			} else {
				st = st[:len(st)-1]
			}
			j++
		}
		k := len(st) - 1
		currCost := 0
		for k >= 0 {
			open := st[k]
			openCost, _ := costMap[open]
			currCost = (currCost * 5) + openCost
			k--
		}
		costs = append(costs, currCost)
		i++
	}
	sort.Ints(costs)
	fmt.Println(costs[len(costs)/2])
}
