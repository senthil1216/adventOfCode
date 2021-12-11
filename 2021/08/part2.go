package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

// 1 == > length = 2
// 4 == > length = 4
// 7 ==> length = 3
// 8 ==> length = 7

func findOverlapLen(pattern, curr string) int {
	var result int
	for _, s1 := range pattern {
		for _, s2 := range curr {
			if s1 == s2 {
				result++
			}
		}
	}
	return result
}

func findDecodedNum(curr string) int {
	if len(curr) == 2 {
		return 1
	} else if len(curr) == 3 {
		return 7
	} else if len(curr) == 4 {
		return 4
	} else if len(curr) == 7 {
		return 8
	}
	return -1
}

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	dirName := path.Dir(filename)
	file, _ := os.Open(dirName + "/input")
	//file, _ := os.Open(dirName + "/input_back")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var sum int
	for scanner.Scan() {
		lines := strings.TrimSpace(scanner.Text())
		inputs := strings.Split(lines, "|")
		decoder := strings.TrimSpace(inputs[0])
		decodedSet := map[int]string{
			0: "",
			1: "",
			2: "",
			3: "",
			4: "",
			5: "",
			6: "",
			7: "",
			8: "",
			9: "",
		}
		patterns := strings.Split(decoder, " ")
		for _, str := range patterns {
			str = strings.TrimSpace(str)
			decodedNum := findDecodedNum(str)
			if decodedNum >= 0 {
				decodedSet[decodedNum] = str
			}
		}
		/*
			at this stage, we have find out which codes are 1, 4, 7, 8
			rest of the numbers to find are 0, 6, 9 & 2, 3, 5
			Decoding logic is as below
				2, 3, 5 ==> length 5
				if len == 5:
					len of characters that overlap between
							2 & 7 = 2 || 2 & 1 = 2 || 2 & 4 = 2 || 2 & 8 = 2
					len of characters that overlap between
							3 & 7 = 3 || 3 & 1 = 2 || 3 & 4 = 2 || 3 & 8 = 4
					len of characters that overlap between
							5 & 7 = 2 || 5 & 1 = 1 || 5 & 4 = 1 || 5 & 8 = 5

				0, 6, 9 ==> length 6
				if len == 6:
					len of characters that overlap between
							0 & 7 = 3 || 0 & 1 = 2 || 0 & 4 = 3 || 0 & 8 = 6
					len of characters that overlap between
							6 & 7 = 2 || 6 & 1 = 1 || 6 & 4 = 3 || 6 & 8 = 6
					len of characters that overlap between
							9 & 7 = 3 || 9 & 1 = 2 || 9 & 4 = 4 || 9 & 8 = 7
		*/

		for _, str := range patterns {
			str = strings.TrimSpace(str)
			lenStr := len(str)
			decodedNum := findDecodedNum(str)
			if decodedNum < 0 {
				overlap7 := findOverlapLen(decodedSet[7], str)
				overlap4 := findOverlapLen(decodedSet[4], str)
				if lenStr == 5 {
					if overlap7 == 3 {
						decodedSet[3] = str
						continue
					}
					if overlap4 == 2 {
						decodedSet[2] = str
						continue
					}
					decodedSet[5] = str
				} else {
					if overlap7 == 2 {
						decodedSet[6] = str
						continue
					}
					if overlap4 == 3 {
						decodedSet[0] = str
						continue
					}
					decodedSet[9] = str
				}
			}
		}
		var total int
		index := 1000
		values := strings.Split(strings.TrimSpace(inputs[1]), " ")
		for _, i := range values {
			for number, s := range decodedSet {
				if len(s) == len(i) && findOverlapLen(s, i) == len(s) {
					total += number * index
					index /= 10
				}
			}
		}
		fmt.Println(total)
		sum += total
	}
	fmt.Println(sum)
}
