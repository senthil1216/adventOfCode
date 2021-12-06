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
	var ints []int
	for i := 0; i < len(s); i++ {
		currNum, _ := strconv.Atoi(s[i])
		ints = append(ints, currNum)
	}
	for i := 0; i < len(ints); i++ {
		//fmt.Println(ints[i])
		num1, num2, num3 := GetTwoPair(i, 2020-ints[i], ints)
		if num1 != 0 && num2 != 0 && num3 != 0 {
			fmt.Println(num1 * num2 * num3)
		}
	}
}

func GetTwoPair(idx int, targetSum int, ints []int) (int, int, int) {
	m := make(map[int]int)
	for i := 0; i < len(ints); i++ {
		if i == idx {
			continue
		}
		currNum := ints[i]
		val, prs := m[currNum]
		if prs {
			return 2020 - targetSum, currNum, val
		}
		if targetSum > currNum {
			m[targetSum-currNum] = currNum
		}
	}
	return 0, 0, 0
}
