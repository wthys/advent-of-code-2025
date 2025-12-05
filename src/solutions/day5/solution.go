package day5

import (
	"fmt"

	"github.com/wthys/advent-of-code-2025/solver"
	"github.com/wthys/advent-of-code-2025/util"
	I "github.com/wthys/advent-of-code-2025/util/interval"
)

type solution struct{}

func init() {
	solver.Register(solution{})
}

func (s solution) Day() string {
	return "5"
}

func (s solution) Part1(input []string, opts solver.Options) (string, error) {
	freshdb, ingredients, err := readInput(input)
	if err != nil {
		return solver.Error(err)
	}

	not_spoiled := 0
	for _, ing := range ingredients {

		fresh := freshdb.Contains(ing)
		opts.Debugf("%v is fresh? %t\n", ing, fresh)
		if fresh {
			not_spoiled += 1
		}
	}

	return solver.Solved(not_spoiled)
}

func (s solution) Part2(input []string, opts solver.Options) (string, error) {
	freshdb, _, err := readInput(input)
	if err != nil {
		return solver.Error(err)
	}

	total := 0
	freshdb = freshdb.Compact()

	for _, ivl := range freshdb {
		total += ivl.Len()
	}

	return solver.Solved(total)
}

func readInput(input []string) (I.Intervals, []int, error) {
	fresh := I.Intervals{}
	ingredients := []int{}

	for _, line := range input {
		nums, _ := util.StringsToInts(util.ExtractRegex("[0-9]+", line))

		if len(nums) == 2 {
			fresh = append(fresh, I.New(nums[0], nums[1]))
		} else if len(nums) == 1 {
			ingredients = append(ingredients, nums[0])
		}

	}

	if len(fresh) == 0 && len(ingredients) == 0 {
		return I.Intervals{}, []int{}, fmt.Errorf("no db contents found")
	}

	return fresh, ingredients, nil
}