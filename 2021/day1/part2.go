package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	s := strings.Split(string(dat), "\n")
	ints := make([]int, len(s))
	for i, s := range s {
		ints[i], _ = strconv.Atoi(s)
	}
	count := 0
	low := 0
	prev := 0
	for low+2 < len(ints) {
		curr := ints[low] + ints[low+1] + ints[low+2]
		if prev > 0 && curr > prev {
			count++
		}
		prev = curr
		low++
	}
	fmt.Println(count)
}
