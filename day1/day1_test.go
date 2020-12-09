package day1

import "testing"

func TestExampleA(t *testing.T) {
	expenses := []string{"1721", "979", "366", "299", "675", "1456"}

	total := Day1a(expenses)
	if total != 514579 {
		t.Errorf("Multiplication was wrong, got: %d, want: %d", total, 514579)
	}
}
func TestExampleB(t *testing.T) {
	expenses := []string{"1721", "979", "366", "299", "675", "1456"}

	total := Day1b(expenses)
	if total != 241861950 {
		t.Errorf("Multiplication was wrong, got: %d, want: %d", total, 241861950)
	}
}
