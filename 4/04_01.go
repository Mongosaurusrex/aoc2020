package day4

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//DayFourOne Day four task one
func DayFourOne() {
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
		}

		if currentSoFarValid {
			validPassports++
		}

		currentSoFarValid = true
	}

	fmt.Println(validPassports)
}
