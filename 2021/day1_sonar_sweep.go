package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func part1() int {
	bytes, err := ioutil.ReadFile("2021/day1_sonar_sweep.txt")
	if err != nil {
		log.Fatal(err)
	}

	content := string(bytes)
	depthIncrements := 0

	// Insert numbers into a slice
	contentSlice := strings.Split(content, "\n")
	for i := range contentSlice {
		if i == 0 {
			continue
		}

		currNum, _ := strconv.Atoi(contentSlice[i])
		prevNum, _ := strconv.Atoi(contentSlice[i - 1])
		if currNum > prevNum {
			depthIncrements++
		}
	}

	return depthIncrements
}

func part2() int {
	bytes, err := ioutil.ReadFile("2021/day1_sonar_sweep.txt")
	if err != nil {
		log.Fatal(err)
	}

	content := string(bytes)
	depthIncrements := 0

	// Insert numbers into a slice
	contentSlice := strings.Split(content, "\n")
	for i := range contentSlice {
		j := i + 1
		k := i + 2
		l := i + 3

		if l >= len(contentSlice) {
			break
		}

		iNum, _ := strconv.Atoi(contentSlice[i])
		jNum, _ := strconv.Atoi(contentSlice[j])
		kNum, _ := strconv.Atoi(contentSlice[k])
		lNum, _ := strconv.Atoi(contentSlice[l])

		firstWindow := iNum + jNum + kNum
		secondWindow := jNum + kNum + lNum

		if secondWindow > firstWindow {
			depthIncrements++
		}
	}

	return depthIncrements
}

func main() {

	fmt.Println(part1())
	fmt.Println(part2())
}
