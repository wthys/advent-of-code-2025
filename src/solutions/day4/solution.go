package day4

import (
	"github.com/wthys/advent-of-code-2025/solver"
	G "github.com/wthys/advent-of-code-2025/grid"
	S "github.com/wthys/advent-of-code-2025/collections/set"
	L "github.com/wthys/advent-of-code-2025/location"
)

type solution struct{}

func init() {
	solver.Register(solution{})
}

func (s solution) Day() string {
	return "4"
}

func (s solution) Part1(input []string, opts solver.Options) (string, error) {
	rolls := readInput(input)

	reachable := S.NewFor(L.New(0, 0))
	rolls.ForEach(func (pos L.Location) {
		if S.New(pos.Neejbers()...).Intersect(rolls).Len() < 4 {
			reachable.Add(pos)
		}
	})

	return solver.Solved(reachable.Len())
}

func (s solution) Part2(input []string, opts solver.Options) (string, error) {
	rolls := readInput(input)

	table := map[L.Location]*S.Set[L.Location]{}

	rolls.ForEach(func (pos L.Location) {
		table[pos] = S.New(pos.Neejbers()...).Intersect(rolls)
	})

	removed := 0

	for true {
		candidates := S.NewFor(L.New(0, 0))
		for roll, neejbers := range table {
			if neejbers.Len() < 4 {
				candidates.Add(roll)
			}
		}

		if candidates.Len() == 0 {
			break
		}

		newTable := map[L.Location]*S.Set[L.Location]{}
		for roll, neejbers := range table {
			if candidates.Has(roll) {
				removed += 1
				continue
			}

			newTable[roll] = neejbers.Subtract(candidates)
		}

		table = newTable
	}

	return solver.Solved(removed)
}

func readInput(input []string) *S.Set[L.Location] {
	rolls := S.NewFor(L.New(0,0))
	for y, line := range input {
		for x, c := range line {
			if string(c) == "@" {
				rolls.Add(L.New(x, y))
			}
		}
	}

	return rolls
}

func visualiseRolls(rolls *S.Set[L.Location], poi *S.Set[L.Location]) {
	grid := G.WithDefault(".")

	if rolls != nil {
		rolls.ForEach(func (loc L.Location) {
			grid.Set(loc, "@")
		})
	}

	if poi != nil {
		poi.ForEach(func (loc L.Location) {
			grid.Set(loc, "x")
		})
	}

	grid.Print()
}