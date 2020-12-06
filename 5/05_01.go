package day5

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//DayFiveOne Day five task one
func DayFiveOne() {
	amountOfRows := createNumberArray(0, 127)
	amountOfColumns := createNumberArray(0, 7)
	highestSeatId := 0

	input, err := ioutil.ReadFile("./5/input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	boardingCards := strings.Split(string(input), "\n")

	for i := 0; i < len(boardingCards); i++ {
		currentSeatId := 0
		ourRows := amountOfRows
		ourColumns := amountOfColumns
		firstSeven := boardingCards[i][0:7]
		for j := 0; j < len(firstSeven); j++ {
			switch letter := string(firstSeven[j]); letter {
			case "B":
				ourRows = ourRows[len(ourRows)/2 : len(ourRows)]
			case "F":
				ourRows = ourRows[0 : len(ourRows)/2]
			}
		}
		lastThree := boardingCards[i][7:]
		for z := 0; z < len(lastThree); z++ {
			switch letter := string(lastThree[z]); letter {
			case "R":
				ourColumns = ourColumns[len(ourColumns)/2 : len(ourColumns)]
			case "L":
				ourColumns = ourColumns[0 : len(ourColumns)/2]
			}
		}

		currentSeatId = ourRows[0]*8 + ourColumns[0]

		if currentSeatId > highestSeatId {
			highestSeatId = currentSeatId
		}
	}

	fmt.Println(highestSeatId)
}

func createNumberArray(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
