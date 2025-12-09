package day9

import (
	"fmt"

	"github.com/wthys/advent-of-code-2025/solver"
	"github.com/wthys/advent-of-code-2025/util"
	L "github.com/wthys/advent-of-code-2025/location"
)

type solution struct{}

func init() {
	solver.Register(solution{})
}

func (s solution) Day() string {
	return "9"
}

func (s solution) Part1(input []string, opts solver.Options) (string, error) {
	tiles, err := readInput(input)
	if err != nil {
		return solver.Error(err)
	}

	largest := -1
	extra := L.New(1, 1)

	for idx, a := range tiles[:len(tiles) - 1] {
		for _, b := range tiles[idx + 1:] {
			rect := b.Subtract(a).Abs().Add(extra)
			area := rect.X * rect.Y
			if area > largest {
				opts.Debugf("NEW largest rect: %v-%v, %v (%v tiles)\n", a, b, rect, area)
				largest = area
			}
		}
	}

	return solver.Solved(largest)
}

func (s solution) Part2(input []string, opts solver.Options) (string, error) {
	tiles, err := readInput(input)
	if err != nil {
		return solver.Error(err)
	}
	
	return solver.NotImplemented()
}

func readInput(input []string) (L.Locations, error) {
	tiles := L.Locations{}

	for _, line := range input {
		nums := util.ExtractInts(line)
		if len(nums) != 2 {
			continue
		}

		tiles = append(tiles, L.New(nums[0], nums[1]))
	}

	if len(tiles) == 0 {
		return L.Locations{}, fmt.Errorf("no tiles found")
	}
	return tiles, nil
}
