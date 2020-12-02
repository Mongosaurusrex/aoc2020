package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")

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
		charPos := strings.Split(tempKey[0], "-")
		first, errorFrom := strconv.Atoi(charPos[0])
		second, errorTo := strconv.Atoi(charPos[1])

		if errorFrom != nil || errorTo != nil {
			fmt.Println("Some error :/")
			os.Exit(1)
		}

		if string(pwd[first]) == charToCheck && string(pwd[second]) == charToCheck {
			continue
		}
		if string(pwd[first]) == charToCheck || string(pwd[second]) == charToCheck {
			amountOfCorrectPwds++
		}
	}

	fmt.Println(amountOfCorrectPwds)
}
