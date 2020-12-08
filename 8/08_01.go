package day8

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//DayEightOne Day eight task one
func DayEightOne() {
	input, err := ioutil.ReadFile("./8/input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	accumulator := 0
	rowsAlreadyChecked := []int{}
	row := 0
	instrucions := strings.Split(string(input), "\n")

	infiniteLoopDetected := false

	for !infiniteLoopDetected {
		if IsNumberInList(row, rowsAlreadyChecked) {
			infiniteLoopDetected = true
			break
		}

		formattedCommand := FormatCommand(instrucions[row])
		valueOfCommand := formattedCommand[1]
		switch command := formattedCommand[0]; command {
		case "acc":
			if strings.Contains(valueOfCommand, "-") {
				numberToSubtract, err := strconv.Atoi(valueOfCommand[1:])
				if err != nil {
					os.Exit(1)
				}
				rowsAlreadyChecked = append(rowsAlreadyChecked, row)
				accumulator -= numberToSubtract
				row++
			} else {
				numberToAdd, err := strconv.Atoi(valueOfCommand[1:])
				if err != nil {
					os.Exit(1)
				}
				rowsAlreadyChecked = append(rowsAlreadyChecked, row)
				accumulator += numberToAdd
				row++
			}
		case "jmp":
			if strings.Contains(valueOfCommand, "-") {
				numberToSubtract, err := strconv.Atoi(valueOfCommand[1:])
				if err != nil {
					os.Exit(1)
				}
				rowsAlreadyChecked = append(rowsAlreadyChecked, row)
				row -= numberToSubtract
			} else {
				numberToAdd, err := strconv.Atoi(valueOfCommand[1:])
				if err != nil {
					os.Exit(1)
				}
				rowsAlreadyChecked = append(rowsAlreadyChecked, row)
				row += numberToAdd
			}
		case "nop":
			rowsAlreadyChecked = append(rowsAlreadyChecked, row)
			row++
		}
	}
	fmt.Println(rowsAlreadyChecked)
	fmt.Println(accumulator)
}

//FormatCommand
func FormatCommand(command string) []string {
	return strings.Split(command, " ")
}

//IsNumberInList
func IsNumberInList(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}
