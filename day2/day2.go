package day2

import (
	"fmt"
	"regexp"
)

// Day2a validate "passwords"
func Day2a(in []string) []bool {
	var result []bool
	for _, s := range in {
		min, max, char, password := parseStringKeys(s)
		re1, _ := regexp.Compile(char)
		count := len(re1.FindAllString(password, 0))
		fmt.Printf("min: %d, max: %d", min, max)
		if min >= count && max <= count {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return result
}

func parseStringKeys(input string) (min int, max int, char string, password string) {
	re1, _ := regexp.Compile("(\\d)-(\\d) ([a-z]): ([a-z]+)")
	matches := re1.FindAllStringSubmatch(input, 0)

	fmt.Println(input)
	fmt.Println(matches)

	return 0, 0, "c", "dasd"
}
