package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
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

	// Part 1 - Find the sum1 of all numbers adjacent (horizontally, verically, or diagonally) to a symbol (anything except a period).
	// Part 2 - Gears are * that are adjacent to 2 numbers. Find the sum of the products of each gear's numbers.
	sum1 := 0
	sum2 := 0
	symbolCoordinates := []coordinate{}
	gearCoordinates := []coordinate{}
	parts := []part{}
	currNum := ""
	currCoordinates := []coordinate{}

	array := [][]string{}
	for _, line := range lines {
		chars := strings.Split(line, "")
		array = append(array, chars)
	}

	for y, row := range array {
		for x, char := range row {
			if strings.Contains("0123456789", char) {
				currNum += char
				currCoordinates = append(currCoordinates, coordinate{x: x, y: y})
			} else if char != "." {
				symbolCoordinates = append(symbolCoordinates, coordinate{x: x, y: y})
				if char == "*" {
					gearCoordinates = append(gearCoordinates, coordinate{x: x, y: y})
				}
			}

			if !strings.Contains("0123456789", char) || x == len(row)-1 {
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

parts:
	for _, p := range parts {
		for _, s := range symbolCoordinates {
			if isAdjacent(p, s) {
				sum1 += p.num
				continue parts
			}
		}
	}

	fmt.Println("Part 1:", sum1)

	for _, g := range gearCoordinates {
		gearParts := []part{}
		for _, p := range parts {
			if isAdjacent(p, g) {
				gearParts = append(gearParts, p)
			}
		}

		if len(gearParts) == 2 {
			gearRatio := gearParts[0].num * gearParts[1].num
			sum2 += gearRatio
		}
	}

	fmt.Println("Part 2:", sum2)
}

func isAdjacent(part part, symbol coordinate) bool {
	for _, coordinate := range part.coordinates {
		if math.Abs(float64(coordinate.x-symbol.x)) <= 1 && math.Abs(float64(coordinate.y-symbol.y)) <= 1 {
			return true
		}
	}
	return false
}
