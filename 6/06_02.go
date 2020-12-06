package day6

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//DaySixTwo Day six task two
func DaySixTwo() {
	total := 0
	input, err := ioutil.ReadFile("./6/input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	customDeclarationPerGroup := strings.Split(string(input), "\n\n")

	for i := 0; i < len(customDeclarationPerGroup); i++ {
		amountOfUnanimousYes := 0
		groupSplit := strings.Split(string(customDeclarationPerGroup[i]), "\n")
		oneStringToRuleThemAll := strings.Join(groupSplit, "")

		perLetterSplit := strings.Split(oneStringToRuleThemAll, "")
		parsedOccurrences := Unique(perLetterSplit)
		desiredLength := len(groupSplit)

		for j := 0; j < len(parsedOccurrences); j++ {
			count := strings.Count(oneStringToRuleThemAll, parsedOccurrences[j])

			if count == desiredLength {
				amountOfUnanimousYes++
			}
		}

		total += amountOfUnanimousYes
	}

	fmt.Println(total)
}
