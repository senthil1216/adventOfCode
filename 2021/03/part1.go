package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//dat, err := os.ReadFile("./input_back")
	dat, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	s := strings.Split(string(dat), "\n")
	var array [12][2]int
	//var array [5][2]int
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			curr := string(s[i][j])
			if curr == "1" {
				array[j][0] += 1
			} else {
				array[j][1] += 1
			}
		}
	}
	var gamma string
	var epsilon string
	for i := 0; i < len(array); i++ {
		if array[i][0] > array[i][1] {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	output_gamma, err := strconv.ParseInt(gamma, 2, 64)
	epsilon_gamma, err := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println(output_gamma * epsilon_gamma)
}
