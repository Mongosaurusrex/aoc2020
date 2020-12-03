package day3

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// DayThreOne Day 3
func DayThreeOne() {
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

	ans := calcCrash(slope, 3, 1)
	fmt.Println(ans)
}

func calcCrash(slope [][]string, dx int, dy int) int {
	x := dx
	tree := 0
	for i := dy; i < len(slope); i += dy {
		if slope[i][x] == "#" {
			tree++
		}
		x = (dx + x) % (len(slope[i]))
	}

	return tree
}
