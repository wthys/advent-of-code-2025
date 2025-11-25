package solver

import (
    "errors"
    "fmt"
    "io"
    "context"
)


const (
    Unknown = "unknown"
    Unsolved = "unsolved"
    Undefined = "undefined"
    InProgress = "in progress"
)

var (
    ErrNotImplemented = errors.New("Not implemented")
)

type (
    Day int

    Options struct {
        debug bool
    }
)

func NotImplemented() (string, error) {
    return Unsolved, ErrNotImplemented
}

func Solved[T any](value T) (string, error) {
    return fmt.Sprintf("%v", value), nil
}

func Error(err error) (string, error) {
    return Unsolved, err
}

func DefaultOptions() Options {
    return Options{false}
}

func (opts Options) Debugf(format string, values ...any) {
    if opts.debug {
        fmt.Printf(format, values...)
    }
}

func (opts Options) IfDebugDo(printer func(opts Options)) {
    if opts.debug {
        printer(opts)
    }
}

func (opts Options) IsDebug() bool {
    return opts.debug
}

type Solver interface{
    Part1(input []string, opts Options) (string, error)
    Part2(input []string, opts Options) (string, error)
    Day() string
}

var (
    solvers = make(map[string]Solver)
)


func Register(solver Solver) {
    if solver == nil {
        panic("puzzle: Register solver is nil")
    }

    name := solver.Day()

    if _, dup := solvers[name]; dup {
        panic(fmt.Errorf("puzzle: Register called twice for solver [%s]", name))
    }

    solvers[name] = solver
}

func GetSolver(day string) (Solver, error) {
    if day == "" {
        return nil, errors.New("empty puzzle day")
    }

    solver, exist := solvers[day]
    if !exist {
        return nil, fmt.Errorf("%s: %w", day, errors.New("unknown puzzle day"))
    }

    return solver, nil
}

func Solve(solver Solver, input io.Reader, ctx context.Context) (Result, error) {
    res := Result{
        Name: solver.Day(),
        Part1: Unsolved,
        Part2: Unsolved,
        Elapsed: nil,
    }

    lines, err := ReadLines(input)

    if err != nil {
        return Result{}, fmt.Errorf("failed to read: %w", err)
    }

    if err := res.AddAnswers(solver, lines, ctx); err != nil {
        return Result{}, fmt.Errorf("failed to add answers: %w", err)
    }

    return res, nil
}
