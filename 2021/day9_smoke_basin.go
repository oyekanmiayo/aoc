package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func day9Part1() {
	bytes, err := ioutil.ReadFile("2021/day9_smoke_basin.txt")
	if err != nil {
		log.Fatal(err)
	}

	content := string(bytes)
	contentSlice := strings.Split(content, "\n")

	matrix := make([][]int, 0)
	for i := range contentSlice {
		subMatrixStr := strings.Split(contentSlice[i], "")

		subMatrix := make([]int, 0)
		for _, numStr := range subMatrixStr {
			num, _ := strconv.Atoi(numStr)
			subMatrix = append(subMatrix, num)
		}

		matrix = append(matrix, subMatrix)
		contentSlice[i] = ""
	}

	// matrix := [][]int{
	// 	{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
	// 	{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
	// 	{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
	// 	{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
	// 	{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	// }

	riskLevelSum := 0

	// 1, 3
	// up = 0, 3, down= 2,3 left=1,2 right=1,4
	for i := range matrix {
		for j := range matrix[i] {
			currVal := matrix[i][j]

			if i-1 >= 0 && matrix[i-1][j] <= currVal {
				continue
			}

			if i+1 < len(matrix) && matrix[i+1][j] <= currVal {
				continue
			}

			if j-1 >= 0 && matrix[i][j-1] <= currVal {
				continue
			}

			if j+1 < len(matrix[i]) && matrix[i][j+1] <= currVal {
				continue
			}

			fmt.Println(i, j, currVal)
			riskLevelSum += currVal + 1
		}
	}

	fmt.Println(riskLevelSum)
}

func day9Part2() {
	bytes, err := ioutil.ReadFile("2021/day9_smoke_basin.txt")
	if err != nil {
		log.Fatal(err)
	}

	content := string(bytes)
	contentSlice := strings.Split(content, "\n")

	matrix := make([][]int, 0)
	for i := range contentSlice {
		subMatrixStr := strings.Split(contentSlice[i], "")

		subMatrix := make([]int, 0)
		for _, numStr := range subMatrixStr {
			num, _ := strconv.Atoi(numStr)
			subMatrix = append(subMatrix, num)
		}

		matrix = append(matrix, subMatrix)
		contentSlice[i] = ""
	}

	// matrix = [][]int{
	// 	{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
	// 	{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
	// 	{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
	// 	{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
	// 	{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	// }

	basinSize1 := math.MinInt32
	basinSize2 := math.MinInt32
	basinSize3 := math.MinInt32

	seen := make(map[int]map[int]bool)

	// 1, 3
	// up = 0, 3, down= 2,3 left=1,2 right=1,4
	for i := range matrix {
		for j := range matrix[i] {

			if !isLow(matrix, i, j) {
				continue
			}

			basinSize := findSizeOfBasin(i, j, matrix, seen)

			if basinSize > basinSize1 {
				basinSize3 = basinSize2
				basinSize2 = basinSize1
				basinSize1 = basinSize
			} else if basinSize > basinSize2 {
				basinSize3 = basinSize2
				basinSize2 = basinSize
			} else if basinSize > basinSize3 {
				basinSize3 = basinSize
			}
		}
	}

	fmt.Println(basinSize1, basinSize2, basinSize3)
	fmt.Println(basinSize1 * basinSize2 * basinSize3)
}

func isLow(matrix [][]int, i, j int) bool {
	if matrix[i][j] == 9 {
		return false
	}

	currVal := matrix[i][j]

	if i-1 >= 0 && matrix[i-1][j] < currVal {
		return false
	}

	if i+1 < len(matrix) && matrix[i+1][j] < currVal {
		return false
	}

	if j-1 >= 0 && matrix[i][j-1] < currVal {
		return false
	}

	if j+1 < len(matrix[i]) && matrix[i][j+1] < currVal {
		return false
	}

	return true
}

func findSizeOfBasin(i, j int, matrix [][]int, seen map[int]map[int]bool) int {
	// if seen before get out
	if seen[i][j] || matrix[i][j] == 9 {
		return 0
	}

	// Add to seen
	_, exists := seen[i]
	if !exists {
		seen[i] = make(map[int]bool)
	}
	seen[i][j] = true

	count := 1
	// val := matrix[i][j]

	// up
	if i-1 >= 0 {
		count += findSizeOfBasin(i-1, j, matrix, seen)
	}

	// down
	if i+1 < len(matrix) {
		count += findSizeOfBasin(i+1, j, matrix, seen)
	}

	// left
	if j-1 >= 0 {
		count += findSizeOfBasin(i, j-1, matrix, seen)
	}

	// right
	if j+1 < len(matrix[i]) {
		count += findSizeOfBasin(i, j+1, matrix, seen)
	}

	return count
}

func main() {
	day9Part2()
}
