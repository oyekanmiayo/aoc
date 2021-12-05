package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func day5Part1() {
	// file, _ := os.Open("2021/day5_hydrothermal_venture.txt")
	// defer file.Close()
	//
	// scanner := bufio.NewScanner(file)
	// var lines []string
	// for scanner.Scan() {
	// 	lines = append(lines, scanner.Text())
	// }

	lines := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}

	// x:y:freq
	coordinates := make(map[int]map[int]int)
	overlapPts := 0

	// treat the 2d space like a matrix. no -ves
	for _, line := range lines {
		ranges := strings.Split(line, " -> ")
		// len(ranges) should be 2 e.g. ["599,531", "599,32"]

		firstRange := strings.Split(ranges[0], ",")
		secondRange := strings.Split(ranges[1], ",")

		x1, _ := strconv.Atoi(firstRange[0])
		y1, _ := strconv.Atoi(firstRange[1])
		x2, _ := strconv.Atoi(secondRange[0])
		y2, _ := strconv.Atoi(secondRange[1])

		if x1 != x2 && y1 != y2 {
			continue
		}

		if x1 == x2 {
			// x is the first key in the map, so this makes sense
			_, exists := coordinates[x1]
			if !exists {
				coordinates[x1] = make(map[int]int)
			}

			y := y1
			// assumes we are going top-down
			bound := y2 + 1
			// if we are going bottom-up, modify
			if y1 > y2 {
				bound = y2 - 1
			}

			for y != bound {
				currFreq, exists := coordinates[x1][y]
				if !exists {
					coordinates[x1][y] = 1
				} else {
					currFreq = currFreq + 1
					// At least two lines have crossed this point
					if currFreq == 2 {
						overlapPts++
					}
					coordinates[x1][y] = currFreq
				}

				if y1 < y2 {
					y++
				} else {
					y--
				}
			}
		}

		if y1 == y2 {

			x := x1

			// assumes we are going left-right
			bound := x2 + 1
			// if we going right-left, modify
			if x1 > x2 {
				bound = x2 - 1
			}

			for x != bound {
				// y can only be checked in the context of an x, so check if x exists
				_, exists := coordinates[x]
				if !exists {
					coordinates[x] = make(map[int]int)
				}

				currFreq, exists := coordinates[x][y1]
				if !exists {
					coordinates[x][y1] = 1
				} else {
					currFreq = currFreq + 1
					// At least two lines have crossed this point
					if currFreq == 2 {
						overlapPts++
					}
					coordinates[x][y1] = currFreq
				}

				if x1 < x2 {
					x++
				} else {
					x--
				}
			}
		}

		fmt.Println("Hi")
	}

	fmt.Println(overlapPts)
}

func day5Part2() {
	file, _ := os.Open("2021/day5_hydrothermal_venture.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// lines := []string{
	// 	"0,9 -> 5,9",
	// 	"8,0 -> 0,8",
	// 	"9,4 -> 3,4",
	// 	"2,2 -> 2,1",
	// 	"7,0 -> 7,4",
	// 	"6,4 -> 2,0",
	// 	"0,9 -> 2,9",
	// 	"3,4 -> 1,4",
	// 	"0,0 -> 8,8",
	// 	"5,5 -> 8,2",
	// }

	// x:y:freq
	coordinates := make(map[int]map[int]int)
	overlapPts := 0

	// treat the 2d space like a matrix. no -ves
	for _, line := range lines {
		ranges := strings.Split(line, " -> ")
		// len(ranges) should be 2 e.g. ["599,531", "599,32"]

		firstRange := strings.Split(ranges[0], ",")
		secondRange := strings.Split(ranges[1], ",")

		x1, _ := strconv.Atoi(firstRange[0])
		y1, _ := strconv.Atoi(firstRange[1])
		x2, _ := strconv.Atoi(secondRange[0])
		y2, _ := strconv.Atoi(secondRange[1])

		xDiffAbs := math.Abs(float64(x2 - x1))
		yDiffAbs := math.Abs(float64(y2 - y1))

		// for diagonal 45deg lines, |x2 - x1| == |y2 - y1|
		if x1 != x2 && y1 != y2 && (xDiffAbs != yDiffAbs) {
			continue
		}

		if x1 == x2 {
			// x is the first key in the map, so this makes sense
			_, exists := coordinates[x1]
			if !exists {
				coordinates[x1] = make(map[int]int)
			}

			y := y1
			bound := y2 + 1
			if y1 > y2 {
				bound = y2 - 1
			}

			for y != bound {

				if x1 == 9 && y == 0 {
					fmt.Println("See what's wrong")
				}

				currFreq, exists := coordinates[x1][y]
				if !exists {
					coordinates[x1][y] = 1
				} else {
					currFreq = currFreq + 1
					// At least two lines have crossed this point
					if currFreq == 2 {
						overlapPts++
					}
					coordinates[x1][y] = currFreq
				}

				if y1 < y2 {
					y++
				} else {
					y--
				}
			}
		}

		if y1 == y2 {

			x := x1

			bound := x2 + 1
			if x1 > x2 {
				bound = x2 - 1
			}

			for x != bound {
				// y can only be checked in the context of an x, so check if x exists
				_, exists := coordinates[x]
				if !exists {
					coordinates[x] = make(map[int]int)
				}

				currFreq, exists := coordinates[x][y1]
				if !exists {
					coordinates[x][y1] = 1
				} else {
					currFreq = currFreq + 1
					// At least two lines have crossed this point
					if currFreq == 2 {
						overlapPts++
					}
					coordinates[x][y1] = currFreq
				}

				if x1 < x2 {
					x++
				} else {
					x--
				}
			}
		}

		if xDiffAbs == yDiffAbs {
			x := x1
			y := y1

			// Assumes left to right
			xBound := x2 + 1
			// Change if it is opp
			if x1 > x2 {
				xBound = x2 - 1
			}
			// Assumes top-down
			yBound := y2 + 1
			// Change if it is opp
			if y1 > y2 {
				yBound = y2 - 1
			}
			for x != xBound && y != yBound {
				_, exists := coordinates[x]
				if !exists {
					coordinates[x] = make(map[int]int)
				}

				currFreq, exists := coordinates[x][y]
				if !exists {
					coordinates[x][y] = 1
				} else {
					currFreq = currFreq + 1
					// At least two lines have crossed this point
					if currFreq == 2 {
						overlapPts++
					}
					coordinates[x][y] = currFreq
				}

				if x2 > x1 {
					x++
				} else {
					x--
				}

				if y2 > y1 {
					y++
				} else {
					y--
				}
			}
		}
	}

	fmt.Println(overlapPts)
}

func main() {
	day5Part2()
}
