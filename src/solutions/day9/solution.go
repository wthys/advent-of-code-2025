package day9

import (
	"fmt"

	"github.com/wthys/advent-of-code-2025/solver"
	"github.com/wthys/advent-of-code-2025/util"
	L "github.com/wthys/advent-of-code-2025/location"
	G "github.com/wthys/advent-of-code-2025/grid"
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

	roundtiles := append(tiles, tiles[0])

	edges := []G.Bounds{}
	util.PairWiseDo(roundtiles, func(a, b L.Location) {
		edges = append(edges, G.BoundsFromLocations(a, b))
	})
	countEdgeTransitions := func(a, b L.Location) int {
		count := 0
		line := G.BoundsFromLocations(a, b)
		for _, edge := range edges {
			if line.Intersects(edge) {
				count += 1
			}
		}
		return count
	}

	
	largest := -1
	for idx, a := range tiles[:len(tiles) - 1] {
		for _, b := range tiles[idx + 1:] {
			rect := G.BoundsFromLocations(a, b)

			// check inner edge of rectangle for edge transitions
			tl := rect.TopLeft().Add(L.New(1,1))
			tr := rect.TopRight().Add(L.New(-1,1))
			bl := rect.BottomLeft().Add(L.New(1,-1))
			br := rect.BottomRight().Add(L.New(-1,-1))
			if countEdgeTransitions(tl, tr) > 0 {
				continue
			}

			if countEdgeTransitions(tr, br) > 0 {
				continue
			}

			if countEdgeTransitions(br, bl) > 0 {
				continue
			}

			if countEdgeTransitions(br, tl) > 0 {
				continue
			}

			size := rect.Width() * rect.Height()
			if size > largest {
				opts.Debugf("NEW largest rect: %v-%v, %v (%v tiles)\n", a, b, rect, size)
				largest = size
			}
		}
	}

	return solver.Solved(largest)
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
