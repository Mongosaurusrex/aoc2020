package day5

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//DayFiveTwo day five task two
func DayFiveTwo() {
	ourSeatId := 0
	amountOfRows := createNumberArray(0, 127)
	amountOfColumns := createNumberArray(0, 7)
	bigBadArray := [128][8]bool{}

	for x := 0; x < len(amountOfRows); x++ {
		bigBadArray[x] = [8]bool{}
		for y := 0; y < len(amountOfColumns); y++ {
			bigBadArray[x][y] = false
		}
	}

	input, err := ioutil.ReadFile("./5/input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	boardingCards := strings.Split(string(input), "\n")

	for i := 0; i < len(boardingCards); i++ {
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

		row := ourRows[0]
		column := ourColumns[0]
		bigBadArray[row][column] = true
	}

	foundRow := -1
	foundColumn := -1

	//Fuck it
	for r := 1; r < len(bigBadArray)-6; r++ {
		for c := 0; c < len(bigBadArray[r]); c++ {
			if bigBadArray[r][c] == false {
				foundRow = r
				foundColumn = c
				break
			}
		}

		if foundColumn != -1 {
			break
		}
	}

	ourSeatId = foundRow*8 + foundColumn

	fmt.Println(ourSeatId)
}
