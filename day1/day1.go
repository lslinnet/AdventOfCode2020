package day1

// Day1a calculates the secret 2020 keys 2SUM
func Day1a(expenses []int) int {
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
func Day1b(expenses []int) int {
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
