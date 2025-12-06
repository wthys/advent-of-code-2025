package day6

import (
	"fmt"
	"regexp"
	"strings"
	"strconv"
	
	"github.com/wthys/advent-of-code-2025/solver"
	"github.com/wthys/advent-of-code-2025/util"
)

type solution struct{}

func init() {
	solver.Register(solution{})
}

func (s solution) Day() string {
	return "6"
}

func (s solution) Part1(input []string, opts solver.Options) (string, error) {
	problems, err := readInput(input)
	if err != nil {
		return solver.Error(err)
	}

	total := 0
	for _, problem := range problems {
		result := problem.Solve()
		opts.Debugf("%v => %v\n", problem, result)
		total += result
	}

	return solver.Solved(total)
}

func (s solution) Part2(input []string, opts solver.Options) (string, error) {
	problems, err := readInput2(input)
	if err != nil {
		return solver.Error(err)
	}

	total := 0
	for _, problem := range problems {
		result := problem.Solve()
		opts.Debugf("%v => %v\n", problem, result)
		total += result
	}

	return solver.Solved(total)
}

var (
	re_opers = regexp.MustCompile("[*+]")
)

type (
	Problem struct {
		values []int
		oper string
	}
)

func readInput(input []string) ([]Problem, error) {
	rows := [][]int{}
	opers := []string{}

	for _, line := range input {
		if re_opers.MatchString(line) {
			opers = util.ExtractRegex("[*+]", line)
			continue
		}

		nums, _ := util.StringsToInts(util.ExtractRegex("[0-9]+", line))
		if len(nums) == 0 {
			continue
		}

		rows = append(rows, nums)
	}

	if len(opers) == 0 {
		return []Problem{}, fmt.Errorf("no operators found")
	}

	problems := []Problem{}
	length := len(opers)

	for idx := 0; idx < length; idx += 1 {
		values := []int{}
		for _, nums := range rows {
			values = append(values, nums[idx])
		}

		problems = append(problems, Problem{values, opers[idx]})
	}

	return problems, nil
}

func (p Problem) String() string {
	joiner := strings.Builder{}
	for i, n := range p.values {
		if i == 0 {
			fmt.Fprint(&joiner, n)
		} else {
			fmt.Fprintf(&joiner, " %v %v", p.oper, n)
		}
	}
	return joiner.String()
}

func (p Problem) Solve() int {
	if p.oper == "*" {
		prod := 1
		for _, n := range p.values {
			prod *= n
		}
		return prod
	}

	sum := 0
	for _, n := range p.values {
		sum += n
	}
	return sum
}

func readInput2(input []string) ([]Problem, error) {
	lastline := -1
	longest := -1

	for no, line := range input {
		if re_opers.MatchString(line) {
			lastline = no
			break
		}
		longest = util.Max(longest, len(line))
	}

	if lastline < 0 {
		return []Problem{}, fmt.Errorf("no opers found")
	}

	indexes := []int{}
	for idx, ch := range input[lastline] {
		if re_opers.MatchString(string(ch)) {
			indexes = append(indexes, idx)
		}
	}
	indexes = append(indexes, longest+1)

	problems := []Problem{}
	re_num := regexp.MustCompile("[0-9]")
	util.PairWiseDo(indexes, func(start, next int) {
		values := []int{}
		for idx := start; idx < next - 1; idx += 1 {
			num := 0
			for _, line := range input[:lastline] {
				lll := fmt.Sprintf("%v%v", line, strings.Repeat(" ", next - start))
				ds := string(lll[idx])
				if re_num.MatchString(ds) {
					d, _ := strconv.Atoi(ds)
					num = 10 * num + d
				}
			}
			values = append(values, num)
		}
		problems = append(problems, Problem{values, string(input[lastline][start])})
	})

	return problems, nil
}