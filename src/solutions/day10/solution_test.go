package day10

import (
	"testing"
)

func TestIndicatorsToggle(t *testing.T) {
	machine := Machine{
		extractIndicators(".##."),
		extractButtons("(3) (1,3) (2) (2,3) (0,2) (0,1)"),
		extractJoltages("3,5,4,7"),
	}

	state := machine.NewState()

	pressed := state.Toggle(Button{3})

	expected := extractIndicators("...#")
	if !pressed.Equals(expected) {
		t.Fatalf("%v.Toggle( (3) ) should yield %v, got %v", state, expected, pressed)
	}

	repressed := pressed.Toggle(Button{3})
	reexp := extractIndicators("....")
	if !repressed.Equals(reexp) {
		t.Fatalf("%v.Toggle( (3) ) should yield %v, got %v", pressed, reexp, repressed)
	}
}