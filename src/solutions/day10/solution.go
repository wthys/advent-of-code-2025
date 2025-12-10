package day10

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/wthys/advent-of-code-2025/solver"
	"github.com/wthys/advent-of-code-2025/util"
	PF "github.com/wthys/advent-of-code-2025/pathfinding"
)

type solution struct{}

func init() {
	solver.Register(solution{})
}

func (s solution) Day() string {
	return "10"
}

func (s solution) Part1(input []string, opts solver.Options) (string, error) {
	machines, err := readInput(input)
	if err != nil {
		return solver.Error(err)
	}

	totalpresses := 0
	for _, machine := range machines {
		neejberFunc := func(node Indicators) []Indicators {
			neejbers := []Indicators{}
			for _, but := range machine.Buttons {
				neejbers = append(neejbers, node.Toggle(but))
			}
			opts.Debugf("  neejbers for [%v] => %v\n", node, neejbers)
			return neejbers
		}
		path, err := PF.ShortestPath(machine.NewState(), machine.Indicators, neejberFunc)
		opts.Debugf("%v\npath (%v) : %v\nerr : %v\n", machine, len(path), path, err)
		if err != nil {
			return solver.Error(err)
		}
		totalpresses += len(path)
	}

	return solver.Solved(totalpresses)
}

func (s solution) Part2(input []string, opts solver.Options) (string, error) {
	return solver.NotImplemented()
}

type (
	Machine struct {
		Indicators Indicators
		Buttons []Button
		Joltages []int
	}
	Machines []Machine

	Indicators string

	Button []int
)

func IndicatorsOff(n int) Indicators {
	return Indicators(strings.Repeat(".", n))
}

func IndicatorsB(onoffs ...bool) Indicators {
	bldr := strings.Builder{}
	for _, onoff := range onoffs {
		if onoff {
			fmt.Fprint(&bldr, "#")
		} else {
			fmt.Fprint(&bldr, ".")
		}
	}
	return Indicators(bldr.String())
}

func (m Machine) NewState() Indicators {
	return IndicatorsOff(len(m.Indicators))
}

func (ind Indicators) ForEach(doer func (int, bool)) {
	for idx, r := range ind {
		doer(idx, string(r) == "#")
	}
}

func (ind Indicators) Flip(n int) Indicators {
	states := []bool{}
	ind.ForEach(func(idx int, state bool) {
		if idx == n {
			states = append(states, !state)
		} else {
			states = append(states, state)
		}
	})
	return IndicatorsB(states...)
}

func (ind Indicators) Toggle(button Button) Indicators {
	toggled := ind
	for _, idx := range button {
		toggled = toggled.Flip(idx)
	}
	return toggled
}

func (this Indicators) Equals(that Indicators) bool {
	return this == that
}

var (
	re_machine = regexp.MustCompile(`\[([^\]]+)\]\ (\([^)]+\)(\ \([^)]+\))*)\ \{([^\}]+)\}`)
	re_button = regexp.MustCompile(`\([^)]+\)`)
)

func readInput(input []string) (Machines, error){
	machines := Machines{}

	for lno, line := range input {
		nums := re_machine.FindStringSubmatch(line)
		if nums == nil {
			continue
		}

		indicators := extractIndicators(nums[1])
		buttons := extractButtons(nums[2])
		joltages := extractJoltages(nums[4])

		if len(indicators) == 0 {
			return Machines{}, fmt.Errorf("#%v : no indicators found", lno)
		}

		if len(buttons) == 0 {
			return Machines{}, fmt.Errorf("#%v : no buttons found", lno)
		}

		if len(joltages) == 0 {
			return Machines{}, fmt.Errorf("#%v : no joltages found", lno)
		}

		machines = append(machines, Machine{indicators, buttons, joltages})
	}

	if len(machines) == 0 {
		return Machines{}, fmt.Errorf("no machines found")
	}

	return machines, nil
}

func extractIndicators(in string) Indicators {
	indicators := []bool{}
	for _, r := range in {
		indicators = append(indicators, string(r) == "#")
	}
	return IndicatorsB(indicators...)
}

func extractButtons(in string) []Button {
	buttons := []Button{}

	buts := re_button.FindAllString(in, -1)
	for _, but := range buts {
		buttons = append(buttons, Button(util.ExtractInts(but)))
	}

	return buttons
}

func extractJoltages(in string) []int {
	return util.ExtractInts(in)
}