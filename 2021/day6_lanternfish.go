package main

import (
	"fmt"
)

func day6Part1(lanternFish []int) int {
	queue := make([][]int, 0)

	// fish = number of days for this fish to create new fish
	for _, fish := range lanternFish {
		queue = append(queue, []int{fish, 0})
	}

	// fmt.Println(queue)

	for len(queue) > 0 {

		fishDetails := queue[0]
		fmt.Println(fishDetails[1])

		if fishDetails[1] == 256 {
			return len(queue)
		}

		queue[0] = nil
		queue = queue[1:]

		if fishDetails[0] == 0 {
			fishDetails[0] = 6
			fishDetails[1]++

			newFishDetails := []int{8, fishDetails[1]}

			queue = append(queue, fishDetails)
			queue = append(queue, newFishDetails)

		} else {

			fishDetails[0]--
			fishDetails[1]++

			queue = append(queue, fishDetails)
		}

	}

	return len(queue)
}

// Depth First + Graph -> Really SLOWWWW
// The queue is growing exponentially, so this will take too long
func day6Part2(lanternFish []int) int {

	dayToFish := make(map[int][]int)
	dayToFish[0] = make([]int, 0)

	// fish = number of days for this fish to create new fish
	for _, fish := range lanternFish {
		dayToFish[0] = append(dayToFish[0], fish)
	}

	queue := make([][]int, 0)

	// fmt.Println(queue)

	for len(queue) > 0 {

		fishDetails := queue[0]
		fmt.Println(fishDetails[1])

		if fishDetails[1] == 256 {
			return len(queue)
		}

		queue[0] = nil
		queue = queue[1:]

		if fishDetails[0] == 0 {
			fishDetails[0] = 6
			fishDetails[1]++

			newFishDetails := []int{8, fishDetails[1]}

			queue = append(queue, fishDetails)
			queue = append(queue, newFishDetails)

		} else {

			fishDetails[0]--
			fishDetails[1]++

			queue = append(queue, fishDetails)
		}

	}

	return len(queue)
}

func day6Part2DP(lanternFish []int) int64 {
	// currentDay:daysToZeroForFish:count
	dp := make(map[int]map[int]int64)
	var count int64
	// fish = number of days for this fish to create new fish
	for _, fish := range lanternFish {
		count += findCountAtGivenDay(256, 0, fish, dp)
		fmt.Println(fish, count)
	}

	return count
}

// 3,4,3,1,2
// 5, 0, 3
// 0 + 3 + 1 => day 4: 6, 8

// map[0][3] = number of fish that have been spawned by fish with an internal timer of 3
// map[4][6] =

// 256, 4, 6 = 4 + 6 + 1 = 11
func findCountAtGivenDay(givenDay, currentDay, internalTimer int, dp map[int]map[int]int64) int64 {
	_, dayExists := dp[currentDay]
	if !dayExists {
		dp[currentDay] = make(map[int]int64)
	}

	_, daysToZeroExists := dp[currentDay][internalTimer]
	if !daysToZeroExists {
		nextValidDay := currentDay + internalTimer + 1

		if nextValidDay > givenDay {
			return 1
		}

		if nextValidDay == givenDay {
			return 2
		}

		var count int64
		count += findCountAtGivenDay(givenDay, nextValidDay, 6, dp)
		count += findCountAtGivenDay(givenDay, nextValidDay, 8, dp)

		dp[currentDay][internalTimer] = count
	}

	return dp[currentDay][internalTimer]
}

func main() {
	lanternFish := []int{
		4, 5, 3, 2, 3, 3, 2, 4, 2, 1, 2, 4, 5, 2, 2, 2, 4, 1, 1, 1, 5, 1, 1, 2, 5, 2, 1, 1, 4, 4, 5,
		5, 1, 2, 1, 1, 5, 3, 5, 2, 4, 3, 2, 4, 5, 3, 2, 1, 4, 1, 3, 1, 2, 4, 1, 1, 4, 1, 4, 2, 5, 1,
		4, 3, 5, 2, 4, 5, 4, 2, 2, 5, 1, 1, 2, 4, 1, 4, 4, 1, 1, 3, 1, 2, 3, 2, 5, 5, 1, 1, 5, 2, 4,
		2, 2, 4, 1, 1, 1, 4, 2, 2, 3, 1, 2, 4, 5, 4, 5, 4, 2, 3, 1, 4, 1, 3, 1, 2, 3, 3, 2, 4, 3, 3,
		3, 1, 4, 2, 3, 4, 2, 1, 5, 4, 2, 4, 4, 3, 2, 1, 5, 3, 1, 4, 1, 1, 5, 4, 2, 4, 2, 2, 4, 4, 4,
		1, 4, 2, 4, 1, 1, 3, 5, 1, 5, 5, 1, 3, 2, 2, 3, 5, 3, 1, 1, 4, 4, 1, 3, 3, 3, 5, 1, 1, 2, 5,
		5, 5, 2, 4, 1, 5, 1, 2, 1, 1, 1, 4, 3, 1, 5, 2, 3, 1, 3, 1, 4, 1, 3, 5, 4, 5, 1, 3, 4, 2, 1,
		5, 1, 3, 4, 5, 5, 2, 1, 2, 1, 1, 1, 4, 3, 1, 4, 2, 3, 1, 3, 5, 1, 4, 5, 3, 1, 3, 3, 2, 2, 1,
		5, 5, 4, 3, 2, 1, 5, 1, 3, 1, 3, 5, 1, 1, 2, 1, 1, 1, 5, 2, 1, 1, 3, 2, 1, 5, 5, 5, 1, 1, 5,
		1, 4, 1, 5, 4, 2, 4, 5, 2, 4, 3, 2, 5, 4, 1, 1, 2, 4, 3, 2, 1,
	}

	fmt.Println(day6Part2DP(lanternFish))
}
