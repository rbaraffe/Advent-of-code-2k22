package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func determineStrokeValue(stroke string) int {
	if stroke == "A" || stroke == "X" {
		return 1
	} else if stroke == "B" || stroke == "Y" {
		return 2
	} else {
		return 3
	}
}


func main() {
	partOne()
	partTwo()
}

func partTwo() {
	readFile, err := os.Open("input.txt")

  if err != nil {
  	fmt.Println(err)
  }

	fileScanner := bufio.NewScanner(readFile)

  fileScanner.Split(bufio.ScanLines)

	var ownScore = 0
	for fileScanner.Scan() {
		var line = fileScanner.Text()

		line_splited := strings.Split(line, " ")
		var opponentStroke = line_splited[0]
		var ouwnStroke = line_splited[1]

		var additionValue = 0
		if ouwnStroke == "Y" {
			additionValue = 3 + determineStrokeValue(opponentStroke)
		} else if ouwnStroke == "Z" {
			var letter = "X"
			if opponentStroke == "A" {
				letter = "Y"
			} else if opponentStroke == "B" {
				letter = "Z"
			}
			additionValue = 6 + determineStrokeValue(letter)
		} else {
			var letter = "Y"
			if opponentStroke == "A" {
				letter = "Z"
			} else if opponentStroke == "B" {
				letter = "X"
			}
			additionValue = determineStrokeValue(letter)
		}
		ownScore = ownScore + additionValue
	}

	fmt.Println("Total", ownScore)
}
func partOne() {
	readFile, err := os.Open("input.txt")

  if err != nil {
  	fmt.Println(err)
  }

	fileScanner := bufio.NewScanner(readFile)

  fileScanner.Split(bufio.ScanLines)

	var ownScore = 0
	for fileScanner.Scan() {
		var line = fileScanner.Text()

		line_splited := strings.Split(line, " ")
		var opponentStroke = line_splited[0]
		var ouwnStroke = line_splited[1]

		var gameStateAddition = 0
		if (opponentStroke == "A" && ouwnStroke == "Y") || (opponentStroke == "B" && ouwnStroke == "Z") || (opponentStroke == "C" && ouwnStroke == "X") {
			gameStateAddition = 6
		} else if determineStrokeValue(opponentStroke) == determineStrokeValue(ouwnStroke) {
			gameStateAddition = 3
		}
		ownScore = ownScore + determineStrokeValue(ouwnStroke) + gameStateAddition
	}

	fmt.Println("Total", ownScore)
}
