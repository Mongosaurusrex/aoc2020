package day4

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//DayFourTwo Day four task two
func DayFourTwo() {
	input, err := ioutil.ReadFile("./4/input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	passports := strings.Split(string(input), "\n\n")

	keys := [...]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	validPassports := 0
	currentSoFarValid := true

	for i := 0; i < len(passports); i++ {
		for j := 0; j < len(keys); j++ {
			if !strings.Contains(passports[i], keys[j]) {
				currentSoFarValid = false
				break
			}

			comboSplit := strings.Fields(passports[i])

			for z := 0; z < len(comboSplit); z++ {
				keyValSplit := strings.Split(comboSplit[z], ":")

				valTemp := string(keyValSplit[1])
				val := strings.Replace(valTemp, "\n", "", -1)
				switch key := keyValSplit[0]; key {
				case "byr":
					intVal, err := strconv.Atoi(val)
					if err != nil {
						fmt.Println("well, what the hell can I do 'bout it")
					}
					currentSoFarValid = intVal <= 2002 && intVal >= 1920
					if !currentSoFarValid {
						break
					}

				case "iyr":
					intVal, err := strconv.Atoi(val)
					if err != nil {
						fmt.Println("well, what the hell can I do 'bout it")
					}
					currentSoFarValid = intVal <= 2020 && intVal >= 2010
					if !currentSoFarValid {
						break
					}

				case "eyr":
					intVal, err := strconv.Atoi(val)
					if err != nil {
						fmt.Println("well, what the hell can I do 'bout it")
					}
					currentSoFarValid = intVal <= 2030 && intVal >= 2020
					if !currentSoFarValid {
						break
					}

				case "hgt":
					if strings.Contains(val, "cm") {
						trimdVal := strings.Trim(val, "cm")
						intVal, err := strconv.Atoi(trimdVal)
						if err != nil {
							fmt.Println("well, what the hell can I do 'bout it")
						}
						currentSoFarValid = intVal >= 150 && intVal <= 193
						if !currentSoFarValid {
							break
						}

					} else if strings.Contains(val, "in") {
						trimdVal := strings.Trim(val, "in")
						intVal, err := strconv.Atoi(trimdVal)
						if err != nil {
							fmt.Println("well, what the hell can I do 'bout it")
						}
						currentSoFarValid = intVal >= 59 && intVal <= 76
						if !currentSoFarValid {
							break
						}

					} else {
						currentSoFarValid = false
						break
					}
				case "hcl":
					matched, err := regexp.MatchString(`^#(?:[0-9a-f]{3}){2}$`, val)
					if err != nil {
						fmt.Println("well, what the hell can I do 'bout it")
					}
					currentSoFarValid = matched
					if !currentSoFarValid {
						break
					}
				case "ecl":
					if len(val) < 3 {
						currentSoFarValid = false
						break
					}
					oneOf := "amb blu brn gry grn hzl oth"
					count := strings.Count(oneOf, val)
					currentSoFarValid = count == 1
					if !currentSoFarValid {
						break
					}

				case "pid":
					matched, err := regexp.MatchString(`\d{9}`, val)
					if err != nil {
						fmt.Println("well, what the hell can I do 'bout it")
					}
					currentSoFarValid = string(val[0]) == "0" && matched && len(val) == 9
				}
			}
		}

		if currentSoFarValid {
			validPassports++
		}

		currentSoFarValid = true
	}

	fmt.Println(validPassports)
}
