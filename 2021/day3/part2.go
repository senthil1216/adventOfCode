package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SelectCommonBit func(countOnes int, countZeroes int) string

func main() {
	//dat, err := os.ReadFile("./input_back")
	dat, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	s := strings.Split(string(dat), "\n")
	oxygen := GetRating(s, SelectMostCommonBit)
	co2 := GetRating(s, SelectLeastCommonBit)
	fmt.Println(oxygen * co2)
}

func SelectMostCommonBit(countOnes int, countZeroes int) string {
	if countOnes >= countZeroes {
		return "1"
	}
	return "0"
}

func SelectLeastCommonBit(countOnes int, countZeroes int) string {
	if countOnes >= countZeroes {
		return "0"
	} else {
		return "1"
	}
}

func GetRating(s []string, selBitFunc SelectCommonBit) int64 {
	i := 0
	for len(s) > 1 {
		countOnes, countZeroes := CountBitsByPos(s, i)
		chosenVal := selBitFunc(countOnes, countZeroes)
		s = SelectValues(s, i, chosenVal)
		i++
	}
	rating, _ := strconv.ParseInt(s[0], 2, 64)
	return rating
}

func SelectValues(s []string, pos int, posValue string) []string {
	var ret []string
	for _, vals := range s {
		binValue := string(vals[pos])
		if binValue == posValue {
			ret = append(ret, vals)
		}
	}
	return ret
}
func CountBitsByPos(s []string, pos int) (int, int) {
	var countOnes int
	var countZeroes int
	for _, vals := range s {
		binValue := string(vals[pos])
		if binValue == "1" {
			countOnes++
		} else {
			countZeroes++
		}
	}
	return countOnes, countZeroes
}
