package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day8Part1() {
	file, _ := os.Open("2021/day8_seven_segment_search.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([][]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " | ")
		lines = append(lines, lineSplit)
	}

	// In the output values, how many times do digits 1, 4, 7, or 8 appear?
	// 1 -> len(2), 4 -> len(4), 7 -> len(3), 8 -> len(7)
	count := 0

	for _, lineSplit := range lines {
		fmt.Println(lineSplit)
		outputPart := strings.Split(lineSplit[1], " ")
		for _, output := range outputPart {
			switch len(output) {
			case 2, 4, 3, 7:
				count++
			}
		}
	}

	fmt.Println(count)
}

func day8Part2() {
	file, _ := os.Open("2021/day8_seven_segment_search.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([][]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " | ")
		lines = append(lines, lineSplit)
	}

	numToCharSet := make(map[int]map[string]bool)
	var result float64

	for _, lineSplit := range lines {
		patterns := strings.Split(lineSplit[0], " ")

		// Extract, 1, 4, 7 and 8
		for _, pattern := range patterns {
			switch len(pattern) {
			case 2:
				numToCharSet[1] = getCharSet(pattern)
			case 4:
				numToCharSet[4] = getCharSet(pattern)
			case 3:
				numToCharSet[7] = getCharSet(pattern)
			case 7:
				numToCharSet[8] = getCharSet(pattern)
			}
		}

		// 2, 3 and 5 have 5 segments
		// 0, 6 and 9 have 6 segments
		for _, pattern := range patterns {
			switch len(pattern) {
			case 5:
				set := getCharSet(pattern)
				if diff(numToCharSet[7], set) == 2 { // diff btw 3 and 7 is 2
					numToCharSet[3] = set
				} else if diff(numToCharSet[4], set) == 5 { // diff btw 2 and 4 is 4
					numToCharSet[2] = set
				} else { // has to be 5 :)
					numToCharSet[5] = set
				}
			case 6:
				set := getCharSet(pattern)
				if diff(numToCharSet[4], set) == 2 { // diff btw 9 and 4 is 2
					numToCharSet[9] = set
				} else if diff(numToCharSet[7], set) == 3 { // diff btw 0 and 7 is 3
					numToCharSet[0] = set
				} else { // has to be 6
					numToCharSet[6] = set
				}
			}
		}


		output := strings.Split(lineSplit[1], " ")

		// It is 4-digit. So first * 1000, second * 100, third * 10, fourth * 1
		currentResult := 0.0
		multiplier := 1000.0

		for _, pattern := range output {
			for k := range numToCharSet {
				set := getCharSet(pattern)
				if diff(set, numToCharSet[k]) == 0 {
					currentResult += multiplier * float64(k)
					multiplier /= 10
				}
			}
		}

		result += currentResult
	}

	fmt.Println(int64(result))
}

func getCharSet(pattern string) map[string]bool {
	charSet := make(map[string]bool)
	for i := range pattern {
		charSet[string(pattern[i])] = true
	}

	return charSet
}

func diff(setA map[string]bool, setB map[string]bool) int {
	diffCount := 0
	for k := range setB {
		_, exists := setA[k]
		if !exists {
			diffCount++
		}
	}

	for k := range setA {
		_, exists := setB[k]
		if !exists {
			diffCount++
		}
	}

	return diffCount
}

func main() {
	day8Part2()
}
