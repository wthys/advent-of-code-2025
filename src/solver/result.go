package solver

import (
    "fmt"
    "time"
    "context"
    "errors"
)

type Result struct{
    Name string
    Part1 string
    Part2 string
    Elapsed []time.Duration
}

func (r Result) String() string {
    if r.Part1 == "" {
        r.Part1 = Unsolved
    }

    if r.Part2 == "" {
        r.Part2 = Unsolved
    }

    if r.Name == "" {
        r.Name = Unknown
    }

    if r.Elapsed != nil && len(r.Elapsed) == 2 {
        return fmt.Sprintf("%v\t%v\t%v\t%v\t%v", r.Name, r.Part1, r.Part2, r.Elapsed[0], r.Elapsed[1])
    }

    return fmt.Sprintf("%v\t%v\t%v", r.Name, r.Part1, r.Part2)
}

func (r *Result) AddAnswers(s Solver, input []string, ctx context.Context) error {
    elapsed, ok := ctx.Value("elapsed").(bool)
    if !ok {
        elapsed = false
    }

    durations := []time.Duration{}

    debug, dok := ctx.Value("debug").(bool)
    opts := Options{dok && debug}

    var start time.Time

    opts.Debugf("===== BEGIN DAY %v PART 1 =====\n", s.Day())
    if (elapsed) {
        start = time.Now()
    }
    part1, err := s.Part1(input, opts)
    if (elapsed) {
        durations = append(durations, time.Since(start))
    }
    opts.Debugf("===== END DAY %v PART 1 =====\n", s.Day())
    if err != nil && !errors.Is(err, ErrNotImplemented) {
        return fmt.Errorf("failed to solve Part1: %w", err)
    }

    opts.Debugf("===== BEGIN DAY %v PART 2 =====\n", s.Day())
    if (elapsed) {
        start = time.Now()
    }
    part2, err := s.Part2(input, opts)
    if (elapsed) {
        durations = append(durations, time.Since(start))
    }
    opts.Debugf("===== END DAY %v PART 2 =====\n", s.Day())
    if err != nil && !errors.Is(err, ErrNotImplemented) {
        return fmt.Errorf("failed to solve Part2: %w", err)
    }

    if !elapsed {
        durations = nil
    }

    r.Part1 = part1
    r.Part2 = part2
    r.Elapsed = durations

    return nil
}
