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

	sliceData := strings.Split(string(input), "\n")

	ans := 0

	for i := 0; i < len(sliceData); i++ {
		for j := i + 1; j < len(sliceData); j++ {
			for z := j + 1; z < len(sliceData); z++ {
				first, _ := strconv.Atoi(sliceData[i])
				second, _ := strconv.Atoi(sliceData[j])
				third, _ := strconv.Atoi(sliceData[z])
				if first+second+third == 2020 {
					ans = first * second * third
				}
			}
		}
	}

	fmt.Println(ans)

}
