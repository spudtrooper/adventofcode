package lib

import (
	"log"
	"strings"
	"unicode"

	"github.com/spudtrooper/adventofcode/common/must"
)

type graph map[string]map[string]bool

type path []string

func (p path) Contains(node string) bool {
	for _, n := range p {
		if n == node {
			return true
		}
	}
	return false
}

func (p path) Add(node string) path {
	return append(p, node)
}

// https://stackoverflow.com/questions/59293525/how-to-check-if-a-string-is-all-upper-or-lower-case-in-go
func isLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func FindPaths(g graph, consider func(path, string) bool) map[string]bool {
	paths := map[string]bool{}
	findPaths(g, []string{"start"}, consider, paths)
	return paths
}

func findPaths(g graph, in path, consider func(path, string) bool, res map[string]bool) {
	node := in[len(in)-1]
	for to := range g[node] {
		if to == "start" {
			continue
		}
		if to == "end" {
			res[strings.Join(in.Add(to), " ")] = true
			continue
		}
		if consider(in, to) {
			findPaths(g, in.Add(to), consider, res)
		}
	}
}

func Part1(input string) int {
	g := graph{}
	for _, line := range must.ReadLines(input) {
		parts := strings.Split(line, "-")
		src, dst := parts[0], parts[1]
		if _, ok := g[src]; !ok {
			g[src] = map[string]bool{}
		}
		g[src][dst] = true
		if _, ok := g[dst]; !ok {
			g[dst] = map[string]bool{}
		}
		g[dst][src] = true
	}

	log.Printf("graph: %+v", g)

	consider := func(p path, n string) bool {
		seen := map[string]bool{}
		for _, n := range p {
			if !isLower(n) {
				continue
			}
			seen[n] = true
		}
		return !seen[n]
	}
	paths := FindPaths(g, consider)

	for p := range paths {
		log.Printf("%+v", p)
	}

	return len(paths)
}

func Part2(input string) int {
	g := graph{}
	for _, line := range must.ReadLines(input) {
		parts := strings.Split(line, "-")
		src, dst := parts[0], parts[1]
		if _, ok := g[src]; !ok {
			g[src] = map[string]bool{}
		}
		g[src][dst] = true
		if _, ok := g[dst]; !ok {
			g[dst] = map[string]bool{}
		}
		g[dst][src] = true
	}

	log.Printf("graph: %+v", g)

	consider := func(p path, n string) bool {
		times := map[string]int{}
		dups := 0
		for _, s := range p.Add(n) {
			if !isLower(s) {
				continue
			}
			times[s]++
			if times[s] > 1 {
				dups++
			}
			if dups > 1 {
				return false
			}
		}
		return true
	}

	paths := FindPaths(g, consider)

	for p := range paths {
		log.Printf("%+v", p)
	}

	return len(paths)
}
