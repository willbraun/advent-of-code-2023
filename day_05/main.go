package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	dataPath, err := filepath.Abs("./data.txt")
	if err != nil {
		log.Fatal("Error getting current directory:", err)
		return
	}

	data, err := os.ReadFile(dataPath)
	if err != nil {
		log.Fatal(err)
	}

	// Part 1
	// get soil number for initial seed
	// get fertilizer number for that soil
	// and so on until I get location number
	// find lowest location number for the initial seeds

	text := string(data)

	sections := strings.Split(text, "map:")
	seeds := regexp.MustCompile(`\d+`).FindAllString(sections[0], -1)
	seedNums := []int{}
	for _, seed := range seeds {
		seedNum, err := strconv.Atoi(seed)
		if err != nil {
			log.Fatal(err)
		}
		seedNums = append(seedNums, seedNum)
	}

	fmt.Println("Part 1:", seedsToLocation(seedNums, sections))

	// in progress
	// Part 2 - seed numbers come in pairs, the start number and the range. Calculate again with new list of seeds.

	// closestLocation := math.Inf(1)

	// for i := 0; i < len(seedNums); i+= 2 {
	// 	fmt.Println(seedNums[i])
	// 	for j := 0; j < seedNums[i+1]; j++ {
	// 		currSeed := seedNums[i] + j
	// 		location := seedToLocation(currSeed, sections)
	// 		location64 := float64(location)
	// 		if location64 < closestLocation {
	// 			closestLocation = location64
	// 		}
	// 	}
	// }

	// fmt.Println("Part 2:", closestLocation)
}

func seedToLocation(seed int, sections []string) int {
	mapResult := seed
	sections:
		for _, section := range sections[1:] {
			lines := strings.Split(section, "\n")
			for _, line := range lines[2:] {
				rowDataStrings := regexp.MustCompile(`\d+`).FindAllString(line, -1)
				if len(rowDataStrings) == 0 {
					continue
				} else {
					rowDataNums := []int{}
					for _, str := range rowDataStrings {
						num, err := strconv.Atoi(str)
						if err != nil {
							log.Fatal(err)
						}
						rowDataNums = append(rowDataNums, num)
					}

					destination := rowDataNums[0]
					source := rowDataNums[1]
					dataRange := rowDataNums[2]

					if (mapResult >= source) && (mapResult < (source + dataRange)) {
						diff := mapResult - source
						mapResult = destination + diff
						continue sections
					}
				}
			}
		}
	return mapResult
}

func seedsToLocation(seeds []int, sections []string) int {
	closestLocation := math.Inf(1)

	for _, seed := range seeds {
		mapResult := seed

	sections:
		for _, section := range sections[1:] {
			lines := strings.Split(section, "\n")
			for _, line := range lines[2:] {
				rowDataStrings := regexp.MustCompile(`\d+`).FindAllString(line, -1)
				if len(rowDataStrings) == 0 {
					continue
				} else {
					rowDataNums := []int{}
					for _, str := range rowDataStrings {
						num, err := strconv.Atoi(str)
						if err != nil {
							log.Fatal(err)
						}
						rowDataNums = append(rowDataNums, num)
					}

					destination := rowDataNums[0]
					source := rowDataNums[1]
					dataRange := rowDataNums[2]

					if (mapResult >= source) && (mapResult < (source + dataRange)) {
						diff := mapResult - source
						mapResult = destination + diff
						continue sections
					}
				}
			}
		}

		mapResult64 := float64(mapResult)
		if mapResult64 < closestLocation {
			closestLocation = mapResult64
		}
	}

	return int(closestLocation)
}
