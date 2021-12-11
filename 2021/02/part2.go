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
	hori := 0
	aim := 0
	depth := 0
	for i := 0; i < len(s); i++ {
		val := strings.Split(s[i], " ")
		dir := val[0]
		num, _ := strconv.Atoi(val[1])
		if dir == "forward" {
			hori += num
			depth += (aim * num)
		} else if dir == "down" {
			aim += num
		} else {
			aim -= num
		}
	}
	fmt.Println(hori * depth)
}
