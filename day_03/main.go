package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
}

type part struct {
	num         int
	coordinates []coordinate
}

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

	// Part 1 - find the sum of all numbers adjacent (horizontally, verically, or diagonally) to a symbol (anything except a period).

	// find all symbols
	// check the 8 adjacent spots from the symbol
	// if one is a number, track left and right until you find a period
	// don't double count a number if it takes up more than 1 of the 8 spots
	// add that number to the total

	// get array of symbol coordinates
	// get array of number coordinates - number, array of coordinates for each number
	// loop through symbol coordinates, find number if

	sum := 0

	array := [][]string{}
	for _, line := range lines {
		chars := strings.Split(line, "")
		array = append(array, chars)
	}

	symbolCoordinates := []coordinate{}
	parts := []part{}

	currNum := ""
	currCoordinates := []coordinate{}

	for y, row := range array {
		for x, char := range row {
			if isSymbol(char) {
				symbolCoordinates = append(symbolCoordinates, coordinate{x, y})
			}

			if isNumber(char) {
				currNum += char
				currCoordinates = append(currCoordinates, coordinate{x: x, y: y})
			}

			if !isNumber(char) || x == len(row)-1 {
				if len(currNum) > 0 {
					num, err := strconv.Atoi(currNum)
					if err != nil {
						log.Fatal(err)
					}
					parts = append(parts, part{num: num, coordinates: currCoordinates})
				}
				currNum = ""
				currCoordinates = []coordinate{}
			}
		}
	}

	for _, part := range parts {
		for _, symbol := range symbolCoordinates {
			if isAdjacent(part, symbol) {
				sum += part.num
			}
		}
	}

	fmt.Println("Part 1:", sum)
}

func isSymbol(char string) bool {
	notSymbols := []string{".", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	return !slices.Contains(notSymbols, char)
}

func isNumber(char string) bool {
	numbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	return slices.Contains(numbers, char)
}

func isAdjacent(part part, symbol coordinate) bool {
	for _, coordinate := range part.coordinates {
		if math.Abs(float64(coordinate.x-symbol.x)) <= 1 && math.Abs(float64(coordinate.y-symbol.y)) <= 1 {
			return true
		}
	}
	return false
}

// func isAdjacent(part part, array [][]string) bool {
// 	ring := []coordinate{}
// 	for i, c := range part.coordinates {
// 		ring = append(ring, coordinate{x: c.x, y: c.y - 1}, coordinate{x: c.x, y: c.y + 1})
// 		if i == 0 {
// 			ring = append(ring, coordinate{x: c.x - 1, y: c.y - 1}, coordinate{x: c.x - 1, y: c.y}, coordinate{x: c.x - 1, y: c.y + 1})
// 		} else if i == len(part.coordinates)-1 {
// 			ring = append(ring, coordinate{x: c.x + 1, y: c.y - 1}, coordinate{x: c.x + 1, y: c.y}, coordinate{x: c.x + 1, y: c.y + 1})
// 		}
// 	}

// 	for _, c := range ring {
// 		if c.x > -1 && c.x < len(array[0]) && c.y > -1 && c.y < len(array) && isSymbol(array[c.y][c.x]) {
// 			return true
// 		}
// 	}

// 	return false
// }
