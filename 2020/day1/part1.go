package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	m := make(map[int]int)
	dat, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	s := strings.Split(string(dat), "\n")
	for _, s := range s {
		currNum, _ := strconv.Atoi(s)
		val, prs := m[currNum]
		if prs {
			fmt.Println(currNum * val)
		}
		m[2020-currNum] = currNum
	}
}
