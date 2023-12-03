package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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
	games := strings.Split(text, "\n")

	// Part 1
	// Each row represents an elf grabbing several handfuls of colored cubes, separated by ";"
	// Find the sum of the IDs of the rows where all handfuls are possible with 12 red, 13 green, and 14 blue cubes
	sum := 0

	colorMap := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

games:
	for _, game := range games {
		data := strings.Split(game, ":")
		gameNum, err := strconv.Atoi(strings.Split(data[0], " ")[1])
		if err != nil {
			log.Fatal(err)
		}
		gameData := data[1]

		rounds := strings.Split(gameData, ";")
		for _, round := range rounds {
			numColorPairs := strings.Split(round, ",")
			for _, numColorPair := range numColorPairs {
				numColorSlice := strings.Split(strings.Trim(numColorPair, " "), " ")
				numStr := numColorSlice[0]
				color := numColorSlice[1]

				num, err := strconv.Atoi(numStr)
				if err != nil {
					log.Fatal(err)
				}

				if num > colorMap[color] {
					continue games
				}
			}
		}
		sum += gameNum
	}

	fmt.Println("Part 1:", sum)

	// Part 2
	// Find the minimum number of cubes of each color required for each game
	// For each game, multiply the counts together, and find the sum across all games
	sum2 := 0

	for _, game := range games {
		gameData := strings.Split(game, ":")[1]
		rounds := strings.Split(gameData, ";")
		gameMaxes := make(map[string]int)

		for _, round := range rounds {
			numColorPairs := strings.Split(round, ",")
			for _, numColorPair := range numColorPairs {
				numColorSlice := strings.Split(strings.Trim(numColorPair, " "), " ")
				numStr := numColorSlice[0]
				color := numColorSlice[1]

				num, err := strconv.Atoi(numStr)
				if err != nil {
					log.Fatal(err)
				}

				max, ok := gameMaxes[color]
				if !ok || num > max {
					gameMaxes[color] = num
				}
			}
		}

		power := gameMaxes["red"] * gameMaxes["green"] * gameMaxes["blue"]
		sum2 += power
	}

	fmt.Println("Part 2:", sum2)
}
