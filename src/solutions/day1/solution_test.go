package day1

import (
	"testing"
)

func TestTurnCount(t *testing.T) {
	dial := Dial(25)
	match := Dial(75)

	turnCount(t, dial, 2, match, 0)
	turnCount(t, dial, 50, match, 0)
	turnCount(t, dial, 100, match, 1)
	turnCount(t, dial, 151, match, 2)
	turnCount(t, dial, -2, match, 0)
	turnCount(t, dial, -50, match, 0)
	turnCount(t, dial, -100, match, 1)
	turnCount(t, dial, -155, match, 2)
}

func turnCount(t *testing.T, dial Dial, amount int, match Dial, expected int) {
	actual := dial.TurnCount(amount, match)
	if actual != expected {
		t.Errorf(`Dial(%v).TurnCount(%v, Dial(%v)) should be %v, got %v`, dial, amount, match, expected, actual)
	}
}