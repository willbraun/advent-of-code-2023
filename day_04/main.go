package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"regexp"
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

	for _, line := range lines {
		splitLine := strings.Split(line, "|")
		winString := strings.Split(splitLine[0], ":")[1]
		winningNums := regexp.MustCompile(`\d+`).FindAllString(winString, -1)
		myNums := regexp.MustCompile(`\d+`).FindAllString(splitLine[1], -1)

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
	}

	fmt.Println("Part 1:", sum)
}
