package day11

import (
	"fmt"
	"slices"

	"github.com/wthys/advent-of-code-2025/solver"
	"github.com/wthys/advent-of-code-2025/util"
	PF "github.com/wthys/advent-of-code-2025/pathfinding"
	S "github.com/wthys/advent-of-code-2025/collections/set"
	D "github.com/wthys/advent-of-code-2025/collections/defaultmap"
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


	nextnodes := S.New("svr")
	toplevels := D.ForKey("svr", func() *S.Set[int] {
		return S.NewFor(0)
	})
	toplevels.Get("svr").Add(0)
	for nextnodes.Len() > 0 {
		newnext := S.NewFor("svr")
		nextnodes.ForEach(func(node string) {
			lvls := toplevels.Get(node)
			lvl := slices.Max(lvls.Values())
			for _, neejber := range connections.Neejbers(node) {
				toplevels.Get(neejber).Add(lvl + 1)
				newnext.Add(neejber)
			}
		})

		nextnodes = newnext
	}
	
	levels := map[string]int{}
	toplevels.ForEach(func(key string, lvls *S.Set[int]) {
		if lvls.Len() > 0 {
			lvl := slices.Max(lvls.Values())
			levels[key] = lvl
		}
	})


	for _, node := range []string{"svr", "fft", "dac", "out"} {
		opts.Debugf("%v @ %v / %v\n", node, toplevels.Get(node), levels[node])
	}



	bfs := PF.ConstructBreadthFirst("svr", connections.NeejberFunc())
	dacbfs := PF.ConstructBreadthFirst("dac", connections.NeejberFunc())
	fftbfs := PF.ConstructBreadthFirst("fft", connections.NeejberFunc())

	svrfft := 0
	fftlvl, _ := levels["fft"]
	bfs.AllPathsFuncVisited("fft", func(path []string) {
		svrfft += 1
	}, func(_ []string, next string) bool {
		return levels[next] > fftlvl
	})

	opts.Debugf("svr -> fft: %v\n", svrfft)

	fftdac := 0
	daclvl, _ := levels["dac"]
	fftbfs.AllPathsFuncVisited("dac", func(path []string) {
		fftdac += 1
	}, func (_ []string, next string) bool {
		return levels[next] > daclvl
	})

	opts.Debugf("fft -> dac: %v\n", fftdac)

	dacout := 0
	dacbfs.AllPathsFunc("out", func(path []string) {
		dacout += 1
	})

	opts.Debugf("dac -> out: %v\n", dacout)

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

func (con Connections) Nodes() []string {
	nodes := S.NewFor("svr")
	for node, nexts := range con {
		nodes.Add(node).AddAll(nexts)
	}
	return nodes.Values()
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
