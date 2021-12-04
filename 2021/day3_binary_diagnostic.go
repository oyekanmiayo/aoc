package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func day3Part1() {
	bytes, _ := ioutil.ReadFile("2021/day3_binary_diagnostic.txt")
	contents := string(bytes)
	contentSlice := strings.Split(contents, "\n")

	oneList := make([]int, 12)
	zeroList := make([]int, 12)

	for _, binaryRep := range contentSlice {
		for i, ch := range binaryRep {
			currChar := string(ch)
			switch currChar {
			case "1":
				oneList[i]++
			case "0":
				zeroList[i]++
			}
		}
	}

	// in binary form
	gammaRateStr := ""
	epsilonRateStr := ""

	// oneList and zeroList have the same length
	// I realize that there's an assumption in the question that there will always be a greater bit
	for i := range oneList {
		if oneList[i] > zeroList[i] {
			gammaRateStr = gammaRateStr + "1"
			epsilonRateStr = epsilonRateStr + "0"
		} else {
			gammaRateStr = gammaRateStr + "0"
			epsilonRateStr = epsilonRateStr + "1"
		}
	}

	// in decimal form
	gammaRate, _ := strconv.ParseInt(gammaRateStr, 2, 64)
	epsilonRate, _ := strconv.ParseInt(epsilonRateStr, 2, 64)
	powerConsumption := gammaRate * epsilonRate

	fmt.Println(oneList, zeroList)
	fmt.Println(gammaRateStr, epsilonRateStr)
	fmt.Println(gammaRate, epsilonRate)
	fmt.Println(powerConsumption)
}

func day3Part2() {
	bytes, _ := ioutil.ReadFile("2021/day3_binary_diagnostic.txt")
	contents := string(bytes)
	contentSlice := strings.Split(contents, "\n")


	// Insert binaries into a set for both values
	// The set is used to keep track of binaries that have been discarded
	ogrSet := make(map[string]bool)
	csrSet := make(map[string]bool)
	for _, binaryRep := range contentSlice {
		ogrSet[binaryRep] = true
		csrSet[binaryRep] = true
	}

	// Generate ogr
	// ogr - Oxygen Generator Rating
	// ogr -> keep only numbers with most common bit (in each pos). If 0 and 1 are eq, keep 1
	// 	   -> stop when only one number is left
	oneList, zeroList := oneAndZeroList(contentSlice)
	ogrStr := ratingValueStr(contentSlice, zeroList, oneList, ogrSet, true)
	ogr, _ := strconv.ParseInt(ogrStr, 2, 64)
	fmt.Println(ogr)

	// Generate ogr
	// csr -> C02 Scrubber Rating
	// csr -> keep only numbers with least common bit (in each pos). If 0 and 1 are eq, keep 0
	//     -> stop when only one number is left
	oneList, zeroList = oneAndZeroList(contentSlice)
	csrStr := ratingValueStr(contentSlice, zeroList, oneList, csrSet, false)
	csr, _ := strconv.ParseInt(csrStr, 2, 64)
	fmt.Println(csr)

	// lsr -> Life Support Rating
	lsr := ogr * csr
	fmt.Println(lsr)
}

func oneAndZeroList(contentSlice []string) ([]int, []int) {
	binaryLen := len(contentSlice[0])
	oneList := make([]int, binaryLen)
	zeroList := make([]int, binaryLen)

	for _, binaryRep := range contentSlice {
		for i, ch := range binaryRep {
			currChar := string(ch)
			switch currChar {
			case "1":
				oneList[i]++
			case "0":
				zeroList[i]++
			}
		}
	}

	return oneList, zeroList
}

// This reduces the frequency of bits (in the binary) from either zeroList or oneList
func deleteBitFrequencyFromLists(zeroList[]int, oneList[]int, binary string) ([]int, []int) {
	// need to delete from multiple positions
	for j, ch := range binary {
		chStr := string(ch)

		// reduce count in oneList or zeroList
		switch chStr {
		case "0":
			zeroList[j]--
		case "1":
			oneList[j]--
		}
	}

	return zeroList, oneList
}

// set can be ogrSet or csrSet
// useMCB means "use most common bit?"
func ratingValueStr(contentSlice []string, zeroList[]int, oneList[]int, set map[string]bool, useMCB bool) string {

	binaryLen := len(contentSlice[0])

	for i := 0; i < binaryLen; i++ {
		bitToCompare := "0"
		if useMCB{
			if oneList[i] >= zeroList[i] {
				bitToCompare = "1"
			}
		} else {
			if zeroList[i] > oneList[i] {
				bitToCompare = "1"
			}
		}


		for _, binaryRep := range contentSlice {
			// if there's just one left, stop
			if len(set) == 1 {
				break
			}

			// ignore binaries that have been discarded before
			exists := set[binaryRep]
			if !exists {
				continue
			}

			currChar := string(binaryRep[i])

			if currChar != bitToCompare {
				delete(set, binaryRep)
				zeroList, oneList = deleteBitFrequencyFromLists(zeroList, oneList, binaryRep)
			}
		}
	}

	ratingStr := ""
	for binary := range set {
		ratingStr = binary
		break
	}

	return ratingStr
}

func main() {
	day3Part1()
	day3Part2()
}
