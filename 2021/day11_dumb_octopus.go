package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func day11Part1() {
	bytes, _ := ioutil.ReadFile("2021/day11_dumb_octopus.txt")
	content := string(bytes)
	contentSlice := strings.Split(content, "\n")

	octMatrix := make([][]int, 0)
	for _, line := range contentSlice {
		row := make([]int, 0)
		rowStr := strings.Split(line, "")

		for _, energyStr := range rowStr {
			energy, _ := strconv.Atoi(energyStr)
			row = append(row, energy)
		}

		octMatrix = append(octMatrix, row)
	}

	var flashes int
	steps := 100
	for i := 0; i < steps; i++ {
		seen := make(map[int]map[int]bool)
		for r := range octMatrix {
			for c := range octMatrix[r] {
				currFlash := increaseEnergyLevels(r, c, octMatrix, seen)
				flashes += currFlash
			}
		}
	}

	fmt.Println(octMatrix)
	fmt.Println(flashes)
}

func day11Part2() {
	bytes, _ := ioutil.ReadFile("2021/day11_dumb_octopus.txt")
	content := string(bytes)
	contentSlice := strings.Split(content, "\n")

	octMatrix := make([][]int, 0)
	for _, line := range contentSlice {
		row := make([]int, 0)
		rowStr := strings.Split(line, "")

		for _, energyStr := range rowStr {
			energy, _ := strconv.Atoi(energyStr)
			row = append(row, energy)
		}

		octMatrix = append(octMatrix, row)
	}

	found := false
	for i := 0; ; i++ {
		seen := make(map[int]map[int]bool)
		for r := range octMatrix {
			for c := range octMatrix[r] {
				currFlash := increaseEnergyLevels(r, c, octMatrix, seen)
				if currFlash == (len(octMatrix) * len(octMatrix[r])) {
					fmt.Println("Step:", i+1, "Flash:", currFlash)
					found = true
					break
				}
			}

			if found {
				break
			}
		}

		if found {
			break
		}
	}

	fmt.Println(octMatrix)
}

func main() {
	day11Part2()
}

func increaseEnergyLevels(row, col int, octMatrix [][]int, seen map[int]map[int]bool) int {
	if seen[row][col] {
		return 0
	}

	if row < 0 || row >= len(octMatrix) || col < 0 || col >= len(octMatrix[row]) {
		return 0
	}

	octMatrix[row][col]++

	// If it won't flash, leave
	if octMatrix[row][col] <= 9 {
		return 0
	}

	flashCnt := 1

	_, exists := seen[row]
	if !exists {
		seen[row] = make(map[int]bool)
	}
	seen[row][col] = true

	// top
	// bottom
	// left
	// right
	// top left
	// top right
	// bottom left
	// bottom right
	flashCnt += increaseEnergyLevels(row-1, col, octMatrix, seen)
	flashCnt += increaseEnergyLevels(row+1, col, octMatrix, seen)
	flashCnt += increaseEnergyLevels(row, col-1, octMatrix, seen)
	flashCnt += increaseEnergyLevels(row, col+1, octMatrix, seen)
	flashCnt += increaseEnergyLevels(row-1, col-1, octMatrix, seen)
	flashCnt += increaseEnergyLevels(row-1, col+1, octMatrix, seen)
	flashCnt += increaseEnergyLevels(row+1, col-1, octMatrix, seen)
	flashCnt += increaseEnergyLevels(row+1, col+1, octMatrix, seen)

	octMatrix[row][col] = 0

	return flashCnt
}
