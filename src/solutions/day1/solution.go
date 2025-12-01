package day1

import (
	"fmt"
	"regexp"

	"github.com/wthys/advent-of-code-2025/solver"
	"github.com/wthys/advent-of-code-2025/util"
)

type solution struct{}

func init() {
	solver.Register(solution{})
}

func (s solution) Day() string {
	return "1"
}

func (s solution) Part1(input []string, opts solver.Options) (string, error) {
	moves, err := readInput(input)
	if err != nil {
		return solver.Error(err)
	}

	to := Dial(50)
	zeros := 0

	for nr, move := range moves {
		to = to.Turn(move)
		opts.Debugf("move %2v => %2v\n", nr, to)
		if to == Dial(0) {
			zeros += 1
		}
	}

	return solver.Solved(zeros)
}

func (s solution) Part2(input []string, opts solver.Options) (string, error) {
	moves, err := readInput(input)
	if err != nil {
		return solver.Error(err)
	}

	dial := Dial(50)
	zero := Dial(0)
	zeroes := 0

	for nr, move := range moves {
		newdial := dial.Turn(move)
		count := dial.TurnCount(move, zero)
		opts.Debugf("move %4v | %2v ==[ %v ]==> %2v (hit 0 %vx)\n", nr, dial, move, newdial, count)
		zeroes += count
		if newdial == zero {
			zeroes += 1
		}
		dial = newdial
	}

	return solver.Solved(zeroes)
}

type (
	Dial int
)

var (
	re_move = regexp.MustCompile("([RL])([0-9]+)")
)

func (from Dial) Left(left int) Dial {
	return from.Turn(-left)
}

func (from Dial) Right(right int) Dial {
	return from.Turn(right)
}

func (from Dial) Turn(amount int) Dial {
	return Dial(((int(from) + amount) % 100 + 100) % 100)
}

func (from Dial) TurnCount(amount int, match Dial) int {
	times := util.Abs(amount) / 100
	to := from.Turn(amount)

	if match == to || to == from {
		return times
	}

	if amount > 0 {
		if from < to {
			if from < match && match < to {
				return times + 1
			}
			return times
		} else {
			if to < match || match < from {
				return times + 1
			}
			return times
		}
	} else {
		if from < to {
			if to < match || match < from {
				return times + 1
			}
			return times
		} else {
			if match < to && from < match {
				return times + 1
			}
			return times
		}
	}

	return times
}

func dir2int(dir string) int {
	if dir == "L" {
		return -1
	} else if dir == "R" {
		return 1
	}
	
	panic(fmt.Sprintf("unknown direction %q", dir))
}

func readInput(input []string) ([]int, error) {
	moves := []int{}

	for nr, line := range input {
		m := re_move.FindStringSubmatch(line)
		if len(m) == 0 {
			continue
		}
		if len(m) != 3 {
			return []int{}, fmt.Errorf("Wrong line format on #%v : %q", nr, line)
		}
		amount, _ := util.StringToInt(m[2])
		dir := dir2int(m[1])
		moves = append(moves, dir * amount)
	}

	return moves, nil
}