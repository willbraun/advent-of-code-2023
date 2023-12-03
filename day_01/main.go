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
	lines := strings.Split(text, "\n")

	// Part 1 - find the sum of the two-digit numbers made from the first and last digit on each line
	sum1 := 0

	for _, line := range lines {
		chars := strings.Split(line, "")
		numChars := []string{}
		for _, char := range chars {
			_, err := strconv.Atoi(char)
			if err == nil {
				numChars = append(numChars, char)
			}
		}

		digit1 := numChars[0]
		digit2 := numChars[len(numChars)-1]
		numStr := digit1 + digit2

		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal(err)
		}

		sum1 += num
	}

	fmt.Println("Part 1:", sum1)

	// Part 2 - same as part 1, but include the spelled out numbers as digits
	sum2 := 0

	numberMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for _, line := range lines {
		indexMap := make(map[int]string)
		for word, num := range numberMap {
			index := strings.Index(line, word)
			if index != -1 {
				indexMap[index] = num
			}

			lastIndex := strings.LastIndex(line, word)
			if lastIndex != -1 {
				indexMap[lastIndex] = num
			}
		}

		chars := strings.Split(line, "")
		for index, char := range chars {
			_, err := strconv.Atoi(char)
			if err == nil {
				indexMap[index] = char
			}
		}

		min := 1000
		max := 0
		for key := range indexMap {
			if key < min {
				min = key
			}
			if key > max {
				max = key
			}
		}

		digit1 := indexMap[min]
		digit2 := indexMap[max]
		numStr := digit1 + digit2

		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal(err)
		}

		sum2 += num
	}

	fmt.Println("Part 2:", sum2)
}