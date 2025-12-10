package grid

import (
	"fmt"

	L "github.com/wthys/advent-of-code-2025/location"
)

type (
	Grid[T any] struct {
		defaultFunc DefaultFunction[T]
		data        map[L.Location]T
	}

	DefaultFunction[T any] func(loc L.Location) (T, error)
	ForEachFunction[T any] func(loc L.Location, value T)
)

// `DefaultValue` creates a `DefaultFunction` that always returns the provided
// value.
func DefaultValue[T any](value T) DefaultFunction[T] {
	return func(_ L.Location) (T, error) {
		return value, nil
	}
}

// `DefaultZero` creates a `DefaultFunction` that always returns the 'zero' value.
func DefaultZero[T any]() DefaultFunction[T] {
	return DefaultValue(*new(T))
}

// `DefaultError` creates a `DefaultFunction` that always returns an error "no
// value at <loc>"
func DefaultError[T any]() DefaultFunction[T] {
	return func(loc L.Location) (T, error) {
		return *new(T), fmt.Errorf("no value at %v", loc)
	}
}

// `New` creates a `Grid` using the `DefaultError` `DefaultFunction` for unknown
// `Location`s. Equivalent to `WithDefaultFunc(DefaultError())`.
func New[T any]() *Grid[T] {
	return WithDefaultFunc(DefaultError[T]())
}

// `WithDefault` creates a `Grid` using the `DefaultValue` `DefaultFunction` for
// unknown `Location`s. Equivalent to `WithDefaultFunc(DefaultValue(value))`.
func WithDefault[T any](value T) *Grid[T] {
	return WithDefaultFunc(DefaultValue(value))
}

// `WithDefaultFunc` creates a `Grid` using the provided `DefaultFunction` for
// unknown `Location`s.
func WithDefaultFunc[T any](defaultFunc DefaultFunction[T]) *Grid[T] {
	return &Grid[T]{defaultFunc, map[L.Location]T{}}
}

// `Get` retrieves the value stored at `loc`. If there is no value stored, the
// `Grid`'s `DefaultFunction` is called. If no `DefaultFunction` was set,
// `DefaultError[T]()` is used.
func (g *Grid[T]) Get(loc L.Location) (T, error) {
	val, ok := g.data[loc]
	if ok {
		return val, nil
	}
	if g.defaultFunc == nil {
		return DefaultError[T]()(loc)
	}

	return g.defaultFunc(loc)
}

// `Set` stores a value at `loc`.
func (g *Grid[T]) Set(loc L.Location, value T) {
	g.data[loc] = value
}

// `Remove` removes the stored value at `loc`, if any.
func (g *Grid[T]) Remove(loc L.Location) {
	delete(g.data, loc)
}

// `ForEach` applies a function to all stored values. Both the `Location` and the
// value are provided to the given `ForEachFunction`.
func (g *Grid[T]) ForEach(forEach ForEachFunction[T]) {
	for loc, value := range g.data {
		forEach(loc, value)
	}
}

// `Bounds` finds the bounding box of the `Location`s of the stored values.
// Returns an error when there are no stored values.
func (g *Grid[T]) Bounds() (Bounds, error) {
	if len(g.data) == 0 {
		return BoundsEmpty(), fmt.Errorf("no values in grid")
	}

	bounds := BoundsEmpty()
	found := false
	apply := func(loc L.Location, _ T) {
		bounds = bounds.Accomodate(loc)

		if !found {
			found = true
		}
	}
	g.ForEach(apply)

	return bounds, nil
}

// `Len` returns the number of stored values.
func (g *Grid[T]) Len() int {
	return len(g.data)
}

func (g *Grid[T]) Print() {
	g.PrintFunc(func(val T, err error) string {
		if err != nil {
			return "."
		}
		return fmt.Sprint(val)
	})
}

func (g *Grid[T]) PrintFunc(stringer func(T, error) string) {
	g.PrintFuncWithLoc(func (_ L.Location, v T, err error) string {
		return stringer(v, err)
	})
}

func (g *Grid[T]) PrintFuncWithLoc(stringer func(L.Location, T, error) string) {
	bounds, err := g.Bounds()

	if err != nil {
		fmt.Println()
		return
	}

	g.PrintBoundsFuncWithLoc(bounds, stringer)
}

func (g *Grid[T]) PrintBoundsFuncWithLoc(bounds Bounds, stringer func(L.Location, T, error) string) {
	min := bounds.TopLeft()
	max := bounds.BottomRight()
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			pos := L.New(x, y)
			val, err := g.Get(pos)
			fmt.Print(stringer(pos, val, err))
		}
		fmt.Println()
	}
}