package day7

import (
	"fmt"

	"github.com/wthys/advent-of-code-2025/solver"
	L "github.com/wthys/advent-of-code-2025/location"
	G "github.com/wthys/advent-of-code-2025/grid"
	S "github.com/wthys/advent-of-code-2025/collections/set"
)

type solution struct{}

func init() {
	solver.Register(solution{})
}

func (s solution) Day() string {
	return "7"
}

func (s solution) Part1(input []string, opts solver.Options) (string, error) {
	start, splitters, err := readInput(input)
	if err != nil {
		return solver.Error(err)
	}

	lowestSplitter := lowest(splitters)
	splits := S.NewFor(start)

	beams := S.New(start)
	for lowestSplitter.Y > lowest(beams).Y {
		newbeams := S.NewFor(start)
		beams.ForEach(func (beam L.Location) {
			newbeam := beam.Add(L.New(0, 1))
			if !splitters.Has(newbeam) {
				newbeams.Add(newbeam)
				return
			}

			left := newbeam.Add(L.New(-1, 0))
			right := newbeam.Add(L.New(1, 0))
			newbeams.Add(left).Add(right)
			splits.Add(newbeam)
		})

		beams = newbeams.Subtract(splitters)
	}

	return solver.Solved(splits.Len())
}

func (s solution) Part2(input []string, opts solver.Options) (string, error) {
	start, splitters, err := readInput(input)
	if err != nil {
		return solver.Error(err)
	}

	lowestSplitter := lowest(splitters)
	splits := S.NewFor(start)
	grid := G.WithDefault(0)
	grid.Set(start, 1)

	incpos := func(pos L.Location, n int) {
		v, _ := grid.Get(pos)
		grid.Set(pos, v + n)
	}

	beams := S.New(start)
	for lowestSplitter.Y > lowest(beams).Y {
		newbeams := S.NewFor(start)
		beams.ForEach(func (beam L.Location) {
			beamnum, _ := grid.Get(beam)
			newbeam := beam.Add(L.New(0, 1))
			if !splitters.Has(newbeam) {
				newbeams.Add(newbeam)
				incpos(newbeam, beamnum)
				return
			}

			left := newbeam.Add(L.New(-1, 0))
			right := newbeam.Add(L.New(1, 0))
			newbeams.Add(left).Add(right)
			incpos(left, beamnum)
			incpos(right, beamnum)
			splits.Add(newbeam)
		})

		beams = newbeams
		opts.IfDebugDo(func (_ solver.Options) {
			opts.Debugf("-------\n")
			grid.PrintFuncWithLoc(func (pos L.Location, value int, _ error) string {
				if value > 0 {
					return fmt.Sprintf("%X", value)
				}
				if splitters.Has(pos) {
					return "^"
				}
				return "."
			})
		})
	}

	total := 0
	grid.ForEach(func(pos L.Location, value int) {
		if pos.Y == lowestSplitter.Y {
			total += value
		}
	})

	return solver.Solved(total)
}

func readInput(input []string) (L.Location, *S.Set[L.Location], error) {
	start := L.New(-1, -1)
	splitters := S.NewFor(start)

	for y, line := range input {
		for x, ch := range line {
			pos := L.New(x, y)
			if string(ch) == "S" {
				start = pos
				continue
			}

			if string(ch) == "^" {
				splitters.Add(pos)
			}
		}
	}

	return start, splitters, nil
}

func lowest(locs *S.Set[L.Location]) L.Location {
	low := L.New(0, -1)
	locs.ForEach(func (pos L.Location) {
		if low.Y < pos.Y {
			low = pos
		}
	})
	return low
}
