package day8

import (
	"fmt"
	"slices"

	"github.com/wthys/advent-of-code-2025/solver"
	"github.com/wthys/advent-of-code-2025/util"
	L "github.com/wthys/advent-of-code-2025/location"
	S "github.com/wthys/advent-of-code-2025/collections/set"
)

type solution struct{}

func init() {
	solver.Register(solution{})
}

func (s solution) Day() string {
	return "8"
}

func (s solution) Part1(input []string, opts solver.Options) (string, error) {
	junctions, err := readInput(input)
	if err != nil {
		return solver.Error(err)
	}

	example := len(input) < 25

	circuits, dists := initCircuits(junctions)

	findCircuit := func(junc L.Location3) (int, *S.Set[L.Location3]) {
		for idx, circ := range circuits {
			if circ.Has(junc) {
				return idx, circ
			}
		}
		return -1, nil
	}

	connections := 0
	for _, shortest := range dists {
		lidx, lcirc := findCircuit(shortest.from)
		ridx, rcirc := findCircuit(shortest.to)

		if lidx == ridx {
			connections += 1
			continue
		}
		
		connections += 1
		newcircs := []*S.Set[L.Location3]{}
		for idx, circ := range circuits {
			if idx == lidx || idx == ridx {
				continue
			}

			newcircs = append(newcircs, circ)
		}

		circuits = append(newcircs, lcirc.Union(rcirc))

		opts.Debugf("_ made connection #%v: %v -> %v | %v\n", connections, shortest.from, shortest.to, asSizes(circuits))
		if example && connections >= 10-1 || connections >= 1000-1 {
			break
		}
	}

	sizes := asSizes(circuits)
	slices.SortFunc(sizes, func(a, b int) int {
		return a-b
	})
	slices.Reverse(sizes)

	prod := 1
	for _, s := range sizes[:3] {
		prod *= s
	}

	opts.Debugf(" sizes: %v\n circuits: %v\n", sizes, len(circuits))
	return solver.Solved(prod)
}

func (s solution) Part2(input []string, opts solver.Options) (string, error) {
	junctions, err := readInput(input)
	if err != nil {
		return solver.Error(err)
	}

	circuits, dists := initCircuits(junctions)

	findCircuit := func(junc L.Location3) (int, *S.Set[L.Location3]) {
		for idx, circ := range circuits {
			if circ.Has(junc) {
				return idx, circ
			}
		}
		return -1, nil
	}

	connections := 0
	lastX1 := -1
	lastX2 := -1
	for _, shortest := range dists {
		lidx, lcirc := findCircuit(shortest.from)
		ridx, rcirc := findCircuit(shortest.to)

		if lidx == ridx {
			connections += 1
			continue
		}
		
		connections += 1
		newcircs := []*S.Set[L.Location3]{}
		for idx, circ := range circuits {
			if idx == lidx || idx == ridx {
				continue
			}

			newcircs = append(newcircs, circ)
		}

		circuits = append(newcircs, lcirc.Union(rcirc))

		opts.Debugf("_ made connection #%v: %v -> %v | %v\n", connections, shortest.from, shortest.to, asSizes(circuits))
		if len(circuits) == 1 {
			lastX1 = shortest.from.X
			lastX2 = shortest.to.X
			break
		}
	}

	return solver.Solved(lastX1 * lastX2)
}


func readInput(input []string) (L.Locations3, error) {
	junctions := L.Locations3{}

	for no, line := range input {
		nums := util.ExtractInts(line)

		if len(nums) == 0 {
			continue
		}

		if len(nums) < 3 {
			return L.Locations3{}, fmt.Errorf("line %v: not enough numbers, expected 3", no + 1)
		}

		junctions = append(junctions, L.New3(nums[0], nums[1], nums[2]))
	}

	if len(junctions) == 0 {
		return L.Locations3{}, fmt.Errorf("no junctions provided")
	}

	return junctions, nil
}

type (
	edge struct {
		from L.Location3
		to L.Location3
		dist float64
	}
)

func mapFunc[T any, R any](input []T, transform func(T) R) []R {
	result := []R{}
	for _, v := range input {
		result = append(result, transform(v))
	}
	return result
}

func asSizes (cs []*S.Set[L.Location3]) []int {
	return mapFunc(cs, func(circ *S.Set[L.Location3]) int {
		return circ.Len()
	})
}

func initCircuits(junctions L.Locations3) ([]*S.Set[L.Location3], []edge) {
	circuits := []*S.Set[L.Location3]{}
	dists := []edge{}
	for iA, juncA := range junctions[:len(junctions)-1] {
		circuits = append(circuits, S.New(juncA))

		for _, juncB := range junctions[:iA+1] {
			if juncA == juncB {
				continue
			}
			e := edge{juncA, juncB, juncA.Subtract(juncB).Magnitude()}
			dists = append(dists, e)
		}
	}

	slices.SortFunc(dists, func(a, b edge) int {
		if a.dist < b.dist {
			return -1
		}
		if a.dist > b.dist {
			return 1
		}
		return 0
	})

	return circuits, dists
}