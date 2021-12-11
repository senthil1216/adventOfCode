package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	dirName := path.Dir(filename)
	file, _ := os.Open(dirName + "/input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var count int
	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), "|")
		values := strings.Split(lines[1], " ")
		for _, val := range values {
			if len(val) == 2 || len(val) == 4 || len(val) == 3 || len(val) == 7 {
				count++
			}
		}
	}
	fmt.Println(count)
}
