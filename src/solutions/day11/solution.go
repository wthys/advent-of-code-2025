package day11

import (
	"fmt"

	"github.com/wthys/advent-of-code-2025/solver"
	"github.com/wthys/advent-of-code-2025/util"
	PF "github.com/wthys/advent-of-code-2025/pathfinding"
)

type solution struct{}

func init() {
	solver.Register(solution{})
}

func (s solution) Day() string {
	return "11"
}

func (s solution) Part1(input []string, opts solver.Options) (string, error) {
	connections, err := readInput(input)
	if err != nil {
		return solver.Error(err)
	}

	bfs := PF.ConstructBreadthFirst("you", connections.NeejberFunc())

	paths := bfs.AllPathsTo("out")

	return solver.Solved(len(paths))
}

func (s solution) Part2(input []string, opts solver.Options) (string, error) {
	connections, err := readInput(input)
	if err != nil {
		return solver.Error(err)
	}

	if len(connections) < 50 {
		extraExample := []string{
			"svr: aaa bbb",
			"aaa: fft",
			"fft: ccc",
			"bbb: tty",
			"tty: ccc",
			"ccc: ddd eee",
			"ddd: hub",
			"hub: fff",
			"eee: dac",
			"dac: fff",
			"fff: ggg hhh",
			"ggg: out",
			"hhh: out",
		}
		connections, _ = readInput(extraExample)
	}






	bfs := PF.ConstructBreadthFirst("svr", connections.NeejberFunc())
	dacbfs := PF.ConstructBreadthFirst("dac", connections.NeejberFunc())
	fftbfs := PF.ConstructBreadthFirst("fft", connections.NeejberFunc())


	svrfft := 0
	bfs.AllPathsFunc("fft", func(_ []string) {
		svrfft += 1
		if svrfft % 10 == 0 {
			opts.Debugf("svr ... fft : %v found\n", svrfft)
		}
	})

	fftdac := 0
	fftbfs.AllPathsFunc("dac", func(_ []string) {
		fftdac += 1
		if fftdac % 10 == 0 {
			opts.Debugf("fft ... dac : %v found\n", fftdac)
		}
	})

	dacout := 0
	dacbfs.AllPathsFunc("out", func(_ []string) {
		dacout += 1
		if dacout % 10 == 0 {
			opts.Debugf("dac ... out : %v found\n", dacout)
		}
	})

	return solver.Solved(svrfft * fftdac * dacout)
}

type (
	Connections map[string][]string
)

func (con Connections) Neejbers(device string) []string {
	n, ok := con[device]
	if ok {
		return n
	}
	return []string{}
}

func (con Connections) NeejberFunc() PF.NeejberFunc[string] {
	return func (node string) []string {
		return con.Neejbers(node)
	}
}

func readInput(input []string) (Connections, error) {
	connections := Connections{}

	for _, line := range input {
		labels := util.ExtractRegex("[a-zA-Z]+", line)
		if len(labels) == 0 {
			continue
		}

		connections[labels[0]] = labels[1:]
	}

	if len(connections) == 0 {
		return Connections{}, fmt.Errorf("no connections found")
	}

	return connections, nil
}
