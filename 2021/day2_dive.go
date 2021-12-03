package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func day2Part1(){
	// Read file
	// Split items by new line to create an array
	// Do some math

	bytes, err := ioutil.ReadFile("2021/day2_dive.txt")
	if err != nil {
		log.Fatal(err)
	}

	content := string(bytes)
	contentSlice := strings.Split(content, "\n")

	// contentSlice := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}

	horPos := 0
	depth := 0

	for _, cmd := range contentSlice {
		cmdSlice := strings.Split(cmd, " ")
		dir := cmdSlice[0]
		unit, _ := strconv.Atoi(cmdSlice[1])

		switch dir {
		case "forward":
			horPos = horPos + unit
		case "down":
			depth = depth + unit
		case "up":
			depth = depth - unit
		}
	}

	fmt.Println(horPos, depth)
	fmt.Println(horPos*depth)
}

func day2Part2()  {
	bytes, err := ioutil.ReadFile("2021/day2_dive.txt")
	if err != nil {
		log.Fatal(err)
	}

	content := string(bytes)
	contentSlice := strings.Split(content, "\n")

	// contentSlice := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}

	horPos := 0
	depth := 0
	aim := 0

	for _, cmd := range contentSlice {
		cmdSlice := strings.Split(cmd, " ")
		dir := cmdSlice[0]
		unit, _ := strconv.Atoi(cmdSlice[1])

		switch dir {
		case "forward":
			horPos = horPos + unit
			depth = depth + (aim * unit)
		case "down":
			aim = aim + unit
		case "up":
			aim = aim - unit
		}
	}

	fmt.Println(horPos, depth)
	fmt.Println(horPos*depth)
}

func main() {
	day2Part1()
	day2Part2()
}
