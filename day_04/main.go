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

	text := string(data)
	lines := strings.Split(text, "\n")
	sum := float64(0)
	lineCounts := make(map[int]int)

	for _, line := range lines {
		splitPipe := strings.Split(line, "|")
		splitColon := strings.Split(splitPipe[0], ":")
		cardNum, err := strconv.Atoi(regexp.MustCompile(`\d+`).FindString(splitColon[0]))
		if err != nil {
			log.Fatal(err)
		}
		// add original card
		lineCounts[cardNum]++

		winningNums := regexp.MustCompile(`\d+`).FindAllString(splitColon[1], -1)
		myNums := regexp.MustCompile(`\d+`).FindAllString(splitPipe[1], -1)

		wins := float64(0)
		for _, w := range winningNums {
			for _, m := range myNums {
				if w == m {
					wins++
				}
			}
		}

		if wins == 0 {
			continue
		}

		sum += math.Pow(2, wins-1)

		// add one copy of next i cards, where i is number of wins
		for i := 0; i < int(wins); i++ {
			// do this for each copy of your current card
			for j := 0; j < lineCounts[cardNum]; j++ {
				lineCounts[cardNum+1+i]++
			}
		}
	}

	fmt.Println("Part 1:", sum)

	totalCount := 0
	for _, count := range lineCounts {
		totalCount += count
	}

	fmt.Println("Part 2:", totalCount)
}
