package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	hori := 0
	vert := 0
	for scanner.Scan() {
		line := scanner.Text()
		vals := strings.Split(line, " ")
		num, _ := strconv.Atoi(vals[1])
		dir := vals[0]
		if dir == "forward" {
			hori += num
		} else if dir == "down" {
			vert += num
		} else {
			vert -= num
		}
	}
	fmt.Println(hori * vert)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
