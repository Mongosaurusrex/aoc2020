package day3

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//DayThreeTwo Day three task two
func DayThreeTwo() {
	input, err := ioutil.ReadFile("./3/input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sliceData := strings.Split(string(input), "\n")
	slope := make([][]string, len(sliceData))

	for index, xString := range sliceData {
		slope[index] = strings.Split(xString, "")
	}

	first := calcCrash(slope, 1, 1)
	second := calcCrash(slope, 3, 1)
	third := calcCrash(slope, 5, 1)
	fourth := calcCrash(slope, 7, 1)
	fifth := calcCrash(slope, 1, 2)

	fmt.Println(first * second * third * fourth * fifth)
}
