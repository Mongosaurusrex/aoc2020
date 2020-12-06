package day6

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//DaySixOne Day six task one
func DaySixOne() {
	totalCustomPoints := 0
	input, err := ioutil.ReadFile("./6/input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	customDeclarationPerGroup := strings.Split(string(input), "\n\n")

	for i := 0; i < len(customDeclarationPerGroup); i++ {
		groupSplit := strings.Split(string(customDeclarationPerGroup[i]), "\n")
		oneStringToRuleThemAll := strings.Join(groupSplit, "")

		perLetterSplit := strings.Split(oneStringToRuleThemAll, "")
		totalCustomPoints += len(Unique(perLetterSplit))
	}

	fmt.Println(totalCustomPoints)
}

func Unique(slice []string) []string {
	// create a map with all the values as key
	uniqMap := make(map[string]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	// turn the map keys into a slice
	uniqSlice := make([]string, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}
	return uniqSlice
}
