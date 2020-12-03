package day1

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//DayOneOne Day one task one
func DayOneOne() {
	input, err := ioutil.ReadFile("./1/input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sliceData := strings.Split(string(input), "\n")

	ans := 0

	for i := 0; i < len(sliceData); i++ {
		for j := i + 1; j < len(sliceData); j++ {
			first, _ := strconv.Atoi(sliceData[i])
			second, _ := strconv.Atoi(sliceData[j])
			if first+second == 2020 {
				ans = first * second
			}
		}
	}

	fmt.Printf("Result: %d", ans)

}
