package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

func day12Part1() {
	bytes, _ := ioutil.ReadFile("2021/day12_passage_pathing.txt")
	content := string(bytes)
	contentSlice := strings.Split(content, "\n")

	// undirected
	graph := make(map[string][]string)
	for _, line := range contentSlice {
		path := strings.Split(line, "-")
		_, exists := graph[path[0]]
		if !exists {
			graph[path[0]] = make([]string, 0)
		}
		_, exists = graph[path[1]]
		if !exists {
			graph[path[1]] = make([]string, 0)
		}

		graph[path[0]] = append(graph[path[0]], path[1])
		graph[path[1]] = append(graph[path[1]], path[0])
	}

	// set for small caves, start and end
	seen := make(map[string]bool)
	pathCnt := findAllUniquePaths("start", graph, seen, "")

	fmt.Println(pathCnt)
}

func findAllUniquePaths(
	cave string,
	graph map[string][]string,
	seen map[string]bool,
	paths string) int {

	if cave == "end" {
		fmt.Println(paths + ",end")
		return 1
	}

	if seen[cave] {
		return 0
	}

	if cave == "start" || isLower(cave) {
		seen[cave] = true
	}

	if cave == "start" {
		paths += "start"
	} else {
		paths += "," + cave
	}

	pathCnt := 0
	for _, childCave := range graph[cave] {
		pathCnt += findAllUniquePaths(childCave, graph, seen, paths)
	}

	if cave != "start" {
		seen[cave] = false
	}

	return pathCnt
}

func isLower(cave string) bool {
	for _, r := range cave {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func day12Part2() {
	bytes, _ := ioutil.ReadFile("2021/day12_passage_pathing.txt")
	content := string(bytes)
	contentSlice := strings.Split(content, "\n")

	// undirected
	graph := make(map[string][]string)
	for _, line := range contentSlice {
		path := strings.Split(line, "-")
		_, exists := graph[path[0]]
		if !exists {
			graph[path[0]] = make([]string, 0)
		}
		_, exists = graph[path[1]]
		if !exists {
			graph[path[1]] = make([]string, 0)
		}

		graph[path[0]] = append(graph[path[0]], path[1])
		graph[path[1]] = append(graph[path[1]], path[0])
	}

	seen := make(map[string]bool)
	seenCnt := make(map[string]int)
	pathCnt := findAllUniquePathsII("start", graph, seen, "", seenCnt, false)

	fmt.Println(pathCnt)
}

func findAllUniquePathsII(
	cave string,
	graph map[string][]string,
	seen map[string]bool,
	paths string,
	seenCnt map[string]int,
	seenTwice bool) int {

	if cave == "end" {
		fmt.Println(paths + ",end")
		return 1
	}

	if seen[cave] {
		return 0
	}

	isLower := isLower(cave)
	if cave == "start" || (isLower && seenTwice) {
		seen[cave] = true
	}

	if isLower && cave != "start" {
		seenCnt[cave]++

		// If another cave is about to be visited twice and one has been visited twice before
		// reject it
		if seenCnt[cave] == 2 && seenTwice {
			seen[cave] = false
			seenCnt[cave]--
			return 0
		}

		// visit a cave twice
		if seenCnt[cave] == 2 {
			seenTwice = true
			seen[cave] = true
		}
	}

	if cave == "start" {
		paths += "start"
	} else {
		paths += "," + cave
	}

	pathCnt := 0
	for _, childCave := range graph[cave] {
		pathCnt += findAllUniquePathsII(childCave, graph, seen, paths, seenCnt, seenTwice)
	}

	if cave != "start" && isLower {
		seen[cave] = false
		seenCnt[cave]--
	}

	return pathCnt
}

func main() {
	day12Part2()
}
