package day2

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//DayTwoOne Day two task one
func DayTwoOne() {
	input, err := ioutil.ReadFile("./2/input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pwdData := strings.Split(string(input), "\n")

	amountOfCorrectPwds := 0

	for i := 0; i < len(pwdData); i++ {
		keyAndPwd := strings.Split(pwdData[i], ":")
		pwd := keyAndPwd[1]
		tempKey := strings.Fields(keyAndPwd[0])
		charToCheck := tempKey[1]
		amountRange := strings.Split(tempKey[0], "-")
		from, errorFrom := strconv.Atoi(amountRange[0])
		to, errorTo := strconv.Atoi(amountRange[1])

		if errorFrom != nil || errorTo != nil {
			fmt.Println("Some error :/")
			os.Exit(1)
		}

		re := regexp.MustCompile(fmt.Sprintf("[^%s]", charToCheck))
		parsedPwd := strings.Join(re.Split(pwd, -1), "")

		if len(parsedPwd) >= from && len(parsedPwd) <= to {
			amountOfCorrectPwds++
		}
	}

	fmt.Println(amountOfCorrectPwds)
}
