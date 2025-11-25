package util

import (
	"fmt"
	"regexp"
	"strconv"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Sign[T Number](val T) T {
	if val == 0 {
		return 0
	}
	return val / Abs(val)
}

func Abs[T Number](val T) T {
	if val < T(0) {
		return -val
	}
	return val
}

func IIf[T any](condition bool, yes, no T) T {
	if condition {
		return yes
	}
	return no
}

func Humanize(val int) string {
	if val < 1000 {
		return fmt.Sprint(val)
	}

	val = val / 1000
	if val < 1000 {
		return fmt.Sprintf("%vK", val)
	}

	val = val / 1000
	if val < 1000 {
		return fmt.Sprintf("%vM", val)
	}

	return fmt.Sprintf("%vG", val/1000)
}

func Max[T Number](values ...T) T {
	if len(values) == 0 {
		panic("need at least one value")
	}
	best := values[0]

	if len(values) == 1 {
		return best
	}

	for _, value := range values[1:] {
		if value > best {
			best = value
		}
	}

	return best
}

func Min[T Number](values ...T) T {
	if len(values) == 0 {
		panic("need at least one value")
	}
	best := values[0]

	if len(values) == 1 {
		return best
	}

	for _, value := range values[1:] {
		if value < best {
			best = value
		}
	}

	return best
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
			t := b
			b = a % b
			a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func ExtractInts(line string) []int {
	values, err := StringsToInts(ExtractRegex("[-+]?[0-9]+", line))
	if err != nil {
		panic(err)
	}
	return values
}

func ExtractRegex(pattern, line string) []string {
	return regexp.MustCompile(pattern).FindAllString(line, -1)
}

func StringsToInts(values []string) ([]int, error) {
	ints := []int{}
	for _, value := range values {
		v, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		ints = append(ints, v)
	}
	return ints, nil
}