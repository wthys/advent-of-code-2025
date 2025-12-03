package day3

import (
	"fmt"
	"strings"

	"github.com/wthys/advent-of-code-2025/solver"
	"github.com/wthys/advent-of-code-2025/util"
)

type solution struct{}

func init() {
	solver.Register(solution{})
}

func (s solution) Day() string {
	return "3"
}

func (s solution) Part1(input []string, opts solver.Options) (string, error) {
	banks, err := readInput(input)
	if err != nil {
		return solver.Error(err)
	}

	total := 0
	for _, bank := range banks {
		joltage := bank.MaxJoltage()
		opts.Debugf("checking %v => max %v\n", bank, joltage)
		total += joltage
	}

	return solver.Solved(total)
}

func (s solution) Part2(input []string, opts solver.Options) (string, error) {
	banks, err := readInput(input)
	if err != nil {
		return solver.Error(err)
	}

	total := 0
	for _, bank := range banks {
		joltage := bank.MaxJoltageN(12)
		opts.Debugf("checking %v => max %v\n", bank, joltage)
		total += joltage
	}

	return solver.Solved(total)
}

type (
	Bank []int
	Banks []Bank
)

func readInput(input []string) (Banks, error) {
	banks := Banks{}

	for _, line := range input {
		batteries, err := util.StringsToInts(util.ExtractRegex("[0-9]", line))
		if err != nil {
			return Banks{}, err
		}

		if len(batteries) == 0 {
			continue
		}

		banks = append(banks, Bank(batteries))
	}

	return banks, nil
}

func (b Bank) String() string {
	builder := strings.Builder{}
	for _, lbl := range b {
		fmt.Fprint(&builder, lbl)
	}
	return builder.String()
}

func (b Bank) MaxJoltage() int {
	maxFirst := -1
	maxFirstIdx := -1

	for idx, lbl := range b {
		if lbl > maxFirst {
			maxFirst = lbl
			maxFirstIdx = idx
		}
	}

	maxSecond := -1
	if maxFirstIdx == len(b) - 1 {
		for _, lbl := range b[:len(b)-1] {
			if lbl > maxSecond {
				maxSecond = lbl
			}
		}
		return 10 * maxSecond + maxFirst
	}

	for _, lbl := range b[maxFirstIdx+1:] {
		if lbl > maxSecond {
			maxSecond = lbl
		}
	}
	return 10 * maxFirst + maxSecond
}

func (b Bank) MaxJoltageN(n int) int {
	earliestIdx := 0
	joltage := 0

	for digit := 0; digit < n; digit += 1 {
		max := -1
		maxIdx := -1
		for idx, lbl := range b[earliestIdx:len(b)-(n-digit-1)] {
			if lbl > max {
				max = lbl
				maxIdx = idx
			}
		}
		joltage = 10 * joltage + max
		earliestIdx += maxIdx + 1
	}

	return joltage
}