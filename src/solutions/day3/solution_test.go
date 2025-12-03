package day3

import (
	"testing"
)

func TestBankMaxJoltage2_MaxJoltage_Equivalent(t *testing.T) {
	banks := Banks{
		Bank{9,8,7,6,5,4,3,2,1,1,1,1,1,1,1},
		Bank{8,1,1,1,1,1,1,1,1,1,1,1,1,1,9},
		Bank{2,3,4,2,3,4,2,3,4,2,3,4,2,7,8},
		Bank{8,1,8,1,8,1,9,1,1,1,1,2,1,1,1},
	}

	for _, bank := range banks {
		p1 := bank.MaxJoltage()
		p2 := bank.MaxJoltageN(2)
		if p1 != p2 {
			t.Errorf("Expected %v to return %v, got %v", bank, p1, p2)
		}
	}
}

type maxJoltage12case struct {
	bank Bank
	expected int
}

func TestBankMaxJoltage12_example(t *testing.T) {
	cases := []maxJoltage12case{
		{Bank{9,8,7,6,5,4,3,2,1,1,1,1,1,1,1}, 987654321111},
		{Bank{8,1,1,1,1,1,1,1,1,1,1,1,1,1,9}, 811111111119},
		{Bank{2,3,4,2,3,4,2,3,4,2,3,4,2,7,8}, 434234234278},
		{Bank{8,1,8,1,8,1,9,1,1,1,1,2,1,1,1}, 888911112111},
	}

	for _, cs := range cases {
		actual := cs.bank.MaxJoltageN(12)
		if actual != cs.expected {
			t.Errorf("Expeced %v.MaxJoltageN(12) to return %v, got %v", cs.bank, cs.expected, actual)
		}
	}
}