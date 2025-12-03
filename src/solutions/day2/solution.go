package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wthys/advent-of-code-2025/solver"
	"github.com/wthys/advent-of-code-2025/util"
	"github.com/wthys/advent-of-code-2025/util/interval"
	"github.com/wthys/advent-of-code-2025/collections/set"
)

type solution struct{}

func init() {
	solver.Register(solution{})
}

func (s solution) Day() string {
	return "2"
}


func (s solution) Part1(input []string, opts solver.Options) (string, error) {
	ranges, err := readInput(input)
	if err != nil {
		return solver.Error(err)
	}

	total := 0

	for _, rng := range ranges {
		low := fmt.Sprint(rng.Lower())
		high := fmt.Sprint(rng.Upper())
		begin := string(low[:len(low)/2])
		end := string(high[:len(high)-len(low)/2])

		nbegin, _ := strconv.Atoi(begin)
		nend, _ := strconv.Atoi(end)
		
		for n := nbegin; n <= nend; n += 1 {
			repeat, _ := strconv.Atoi(fmt.Sprintf("%v%v", n, n))
			if rng.Contains(repeat) {
				total += repeat
			}
		}
	}

	return solver.Solved(total)
}

func (s solution) Part2(input []string, opts solver.Options) (string, error) {
	ranges, err := readInput(input)
	if err != nil {
		return solver.Error(err)
	}

	invalids := set.NewFor(1)
	for _, rng := range ranges {
		opts.Debugf("== checking range %v ==\n", rng)
		low := fmt.Sprint(rng.Lower())
		high := fmt.Sprint(rng.Upper())

		for length := 1; length <= util.Max(len(low), len(high))/2; length+= 1 {
			lowreps := len(low) / length
			hireps := len(high) / length
			if lowreps != hireps {
				opts.Debugf("  checking multiple reps in %v/%v: %v -> %v\n", rng, length, lowreps, hireps)
			}
			
			for n := pow(10, length - 1); n < pow(10, length); n += 1 {
				pattern := fmt.Sprint(n)
				for reps := lowreps; reps <= hireps; reps += 1 {
					if reps < 2 {
						continue
					}
					repeated := strings.Repeat(pattern, reps)
					// opts.Debugf("  checking %v in %v/%vx%v\n", repeated, rng, length, reps)
					num, _ := strconv.Atoi(repeated)
					if rng.Contains(num) && !invalids.Has(num) {
						invalids.Add(num)
						opts.Debugf("found %v (%v x %v)\n", num, pattern, reps)
					}
				}
			}
		}
	}

	total := 0
	invalids.ForEach(func(id int) {
		total += id
	})

	return solver.Solved(total)
}


func readInput(input []string) (interval.Intervals, error) {
	ends := []int{}

	for _, line := range input {
		nums, err := util.StringsToInts(util.ExtractRegex("[0-9]+", line))
		if err != nil {
			return interval.Intervals{}, err
		}
		ends = append(ends, nums...)
	}

	if len(ends) % 2 != 0 {
		return interval.Intervals{}, fmt.Errorf("not enough range ends (got %v)", len(ends))
	}

	ranges := interval.Intervals{}
	for idx := 0; idx < len(ends); idx += 2 {
		ranges = append(ranges, interval.New(ends[idx], ends[idx + 1]))
	}

	return ranges, nil
}

func pow(n, e int) int {
	result := 1
	for e > 0 {
		result *= n
		e -= 1
	}
	return result
}

func divisors(n int) *set.Set[int] {
	divs := set.NewFor(n)

	for d := 1; d < n; d += 1 {
		if n % d == 0 {
			divs.Add(d)
		}
	}

	return divs
}