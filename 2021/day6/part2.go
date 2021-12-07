package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// file, _ := os.Open("./input_back")
	file, _ := os.Open("./input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	lines := strings.Split(scanner.Text(), ",")
	var lifetimes [9]int
	for _, l := range lines {
		num, _ := strconv.Atoi(l)
		lifetimes[num]++
	}
	var i int
	for i < 256 {
		newFishes := lifetimes[0]
		for idx := 0; idx < len(lifetimes)-1; idx++ {
			lifetimes[idx] = lifetimes[idx+1]
		}
		lifetimes[6] += newFishes
		lifetimes[8] = newFishes
		i++
	}
	count := 0
	for j := 0; j < len(lifetimes); j++ {
		count += lifetimes[j]
	}
	fmt.Println(count)
	// var i int
	// for i < 80 {
	// 	i++
	// }
	// 	var count int
	// 	for idx, num := range nums {
	// 		if num == 0 {
	// 			nums[idx] = 6
	// 			count++
	// 		} else {
	// 			nums[idx]--
	// 		}
	// 	}
	// 	for c := 0; c < count; c++ {
	// 		nums = append(nums, 8)
	// 	}
	// 	i++
	// }
}
