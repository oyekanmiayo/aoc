package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func isIllegal(lastStr string, char string) bool {
	if lastStr == "(" && char != ")" {
		return true
	}

	if lastStr == "[" && char != "]" {
		return true
	}

	if lastStr == "{" && char != "}" {
		return true
	}

	if lastStr == "<" && char != ">" {
		return true
	}

	return false
}

func day10Part1(contentSlice []string) {
	errMap := map[string]int64{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	var errScore int64

	for i := range contentSlice {
		charSeq := strings.Split(contentSlice[i], "")
		stack := make([]string, 0)

		for _, char := range charSeq {
			if char == "(" || char == "[" || char == "{" || char == "<" {
				stack = append(stack, char)
			} else {
				// check if stack is empty

				lastIdx := len(stack) - 1
				lastStr := stack[lastIdx]
				stack = stack[:lastIdx]

				if isIllegal(lastStr, char) {
					fmt.Println(i, char)
					errScore += errMap[char]
					break
				}
			}
		}
	}

	fmt.Println(errScore)
}

func day10Part2(contentSlice []string) {

	allScores := make([]int, 0)

	for i := range contentSlice {
		charSeq := strings.Split(contentSlice[i], "")
		stack := make([]string, 0)
		for _, char := range charSeq {
			if char == "(" || char == "[" || char == "{" || char == "<" {
				stack = append(stack, char)
			} else {
				lastIdx := len(stack) - 1
				lastStr := stack[lastIdx]
				stack = stack[:lastIdx]

				if isIllegal(lastStr, char) {
					// Make stack empty if line is corrupt
					stack = make([]string, 0)
					break
				}
			}
		}

		if len(stack) != 0 {
			score := findScore(stack)
			allScores = append(allScores, score)
			fmt.Println(i, stack)
		}
	}

	sort.Ints(allScores)
	midIdx := len(allScores) / 2
	fmt.Println(allScores[midIdx])
}

func findScore(stack []string) int {

	var score int
	scoreMap := map[string]int{
		"(": 1,
		"[": 2,
		"{": 3,
		"<": 4,
	}

	for i := len(stack) - 1; i >= 0; i-- {
		score *= 5
		score += scoreMap[stack[i]]
	}

	return score
}

func main() {

	bytes, _ := ioutil.ReadFile("2021/day10_syntax_scoring.txt")
	content := string(bytes)
	contentSlice := strings.Split(content, "\n")

	day10Part2(contentSlice)

}
