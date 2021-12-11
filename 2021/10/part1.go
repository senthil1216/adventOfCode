package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
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
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	bracketMap := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
		"<": ">",
	}
	totalCost := 0
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
					currCost, _ := costMap[bracket]
					totalCost += currCost
					loopBreak = true
				}
			}
		}
	}
	fmt.Println(totalCost)
}
