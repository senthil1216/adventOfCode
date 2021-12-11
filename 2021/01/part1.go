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
	prev, _ := strconv.Atoi(s[0])
	count := 0
	for j := 1; j < len(s); j++ {
		next, err := strconv.Atoi(s[j])
		if err != nil {
			panic(err)
		}
		if next > prev {
			count++
		}
		prev = next
	}
	fmt.Println(count)
}
