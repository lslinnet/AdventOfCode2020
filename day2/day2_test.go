package day2

import (
	"testing"
)

func TestExampleA(t *testing.T) {
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
	expected := []bool{
		true,
		false,
		true,
	}

	result := Day2a(input)
	for i := 0; i < len(input); i++ {
		t.Logf("Processing index %d", i)
		t.Logf("Result is %t", result[i])
		if result[i] != expected[i] {
			t.Errorf("Expected and return values did not match, at index: %d for value %s", i, input[i])
		}
	}
}

/*
func TestExampleB(t *testing.T) {
	expenses := []int{1721, 979, 366, 299, 675, 1456}

	total := Day1b(expenses)
	if total != 241861950 {
		t.Errorf("Multiplication was wrong, got: %d, want: %d", total, 241861950)
	}
}
*/
