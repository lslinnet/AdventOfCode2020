package day1

import "strconv"

// Day1a calculates the secret 2020 keys 2SUM
func Day1a(in []string) int {
	expenses := convertStringArrayToIntArray(in)

	for x := 0; x < len(expenses); x++ {
		for y := 0; y < len(expenses); y++ {
			if x != y && expenses[x]+expenses[y] == 2020 {
				return expenses[x] * expenses[y]
			}
		}
	}
	return 0
}

// Day1b calculates the secret 2020 keys 3SUM
func Day1b(in []string) int {
	expenses := convertStringArrayToIntArray(in)

	for x := 0; x < len(expenses); x++ {
		for y := 0; y < len(expenses); y++ {
			for z := 0; z < len(expenses); z++ {
				if x != y && x != z && y != z && expenses[x]+expenses[y]+expenses[z] == 2020 {
					return expenses[x] * expenses[y] * expenses[z]
				}
			}
		}
	}
	return 0
}

func convertStringArrayToIntArray(in []string) []int {
	var i2 = []int{}

	for _, i := range in {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		i2 = append(i2, j)
	}

	return i2
}
