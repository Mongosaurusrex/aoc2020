package day8

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//DayEightTwo Day eight task two
func DayEightTwo() {
	input, err := ioutil.ReadFile("./8/input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	instructions := strings.Split(string(input), "\n")
	for i := 0; i < len(instructions); i++ {
		instruction := strings.Split(instructions[i], " ")[0]
		arg := strings.Split(instructions[i], " ")[1]
		alteredInstructions := make([]string, len(instructions))
		copy(alteredInstructions, instructions)
		if instruction == "jmp" {
			alteredInstructions[i] = "nop " + arg
		} else if instruction == "nop" {
			alteredInstructions[i] = "jmp " + arg
		} else {
			continue
		}
		result := OneOnceAgain(alteredInstructions)
		fmt.Println(result)
	}
}

func OneOnceAgain(instructions []string) int {

	accumulator := 0
	rowsAlreadyChecked := []int{}
	row := 0

	infiniteLoopDetected := false

	for !infiniteLoopDetected {
		if row >= len(instructions) {
			break
		}
		if IsNumberInList(row, rowsAlreadyChecked) {
			infiniteLoopDetected = true
			break
		}

		formattedCommand := FormatCommand(instructions[row])
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

	return accumulator
}
