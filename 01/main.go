package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	readFile, err := os.Open("input.txt")

  if err != nil {
  	fmt.Println(err)
  }

	fileScanner := bufio.NewScanner(readFile)

  fileScanner.Split(bufio.ScanLines)

	var mostCaloriesCarrying = 0
	var caloriesAccumulation = 0
  for fileScanner.Scan() {
		var line = fileScanner.Text()
		if line == "" {
			if caloriesAccumulation > mostCaloriesCarrying {
				mostCaloriesCarrying = caloriesAccumulation
			}

			caloriesAccumulation = 0
		}
		currentValue, _ := strconv.Atoi(line)
		caloriesAccumulation = caloriesAccumulation + currentValue
  }
  readFile.Close()

	fmt.Println("Most calories part one", mostCaloriesCarrying)
}

func partTwo() {
	readFile, err := os.Open("input.txt")

  if err != nil {
  	fmt.Println(err)
  }

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var topThree []int
	var caloriesAccumulation = 0
  for fileScanner.Scan() {
		var line = fileScanner.Text()
		if line == "" {
			var replaceNeeded = false
			var indexToReplace = 0
			var diff = 0
			if len(topThree) == 3 {
				for i , v := range topThree {
					if caloriesAccumulation > v {
						replaceNeeded = true
						if diff < caloriesAccumulation - v {
							indexToReplace = i
							diff = caloriesAccumulation - v
						}
					}
				}

				if replaceNeeded == true {
					topThree[indexToReplace] = caloriesAccumulation
					replaceNeeded = false
				}
			} else {
				topThree = append(topThree, caloriesAccumulation)
			}
			caloriesAccumulation = 0
		}

		currentValue, _ := strconv.Atoi(line)
		caloriesAccumulation = caloriesAccumulation + currentValue
	}

	readFile.Close()
	fmt.Println("top three of calories", topThree)
	var finalAccumulation = 0
	for _, v := range topThree {
		finalAccumulation = finalAccumulation + v
	}

	fmt.Println("cumulative calories part two", finalAccumulation)
}
