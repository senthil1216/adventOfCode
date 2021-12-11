package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//file, _ := os.Open("./input_back")
	file, _ := os.Open("./input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var nums []int
	for scanner.Scan() {
		line := scanner.Text()
		lines := strings.Split(line, ",")
		for _, l := range lines {
			num, _ := strconv.Atoi(l)
			nums = append(nums, num)
		}
	}
	var i int
	for i < 80 {
		var count int
		for idx, num := range nums {
			if num == 0 {
				nums[idx] = 6
				count++
			} else {
				nums[idx]--
			}
		}
		for c := 0; c < count; c++ {
			nums = append(nums, 8)
		}
		i++
	}
	fmt.Println(len(nums))
}
