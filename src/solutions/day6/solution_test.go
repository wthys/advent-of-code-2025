package day6

import (
	"testing"
)

func TestReadInput2(t *testing.T) {
	input := []string{
		"123 ",
		" 45 ",
		"  6 ",
		"*   ",
	}

	problems, err := readInput2(input)

	if len(problems) != 1 {
		t.Errorf("Expected 1 problem, got %v", len(problems))
	}

	if err != nil {
		t.Errorf("Did not expect an error, got %v", err)
	}

	assertEquals(t, problems[0].values[0], 1, "first value")
	assertEquals(t, problems[0].values[1], 24, "second value")
	assertEquals(t, problems[0].values[2], 356, "third value")
	assertEquals(t, problems[0].oper, "*", "oper")
}

func assertEquals[T comparable](t *testing.T, actual T, expected T, label string) {
	if actual != expected {
		t.Errorf("Expected %v to be %v, got %v", label, expected, actual)
	}
}