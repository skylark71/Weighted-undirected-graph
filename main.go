package main

import (
	"fmt"
	"time"
)

type Graph struct {
	NumNodes int
	Edges    [][]Edge
}

type Edge struct {
	From   int
	To     int
	Weight int
}

// NewGraph: Create graph with n nodes.
func NewGraph(n int) *Graph {
	return &Graph{
		NumNodes: n,
		Edges:    make([][]Edge, n),
	}
}

// AddEdge: Add an edge from u to v.
func (g *Graph) AddEdge(u, v, w int) {
	g.Edges[u] = append(g.Edges[u], Edge{From: u, To: v, Weight: w})

	// For undirected graph add edge from v to u.
	// g.Edges[v] = append(g.Edges[v], Edge{From: v, To: u, Weight: w})
}

func (g *Graph) adjacentEdgesExample(size, start, end int) {
	var visit = make([]bool, size)
	var m, max int
	u := start
	fmt.Printf("Node:%v - ", start)
	for u != end {
		if visit[u] {
			fmt.Printf("Node %v Unreachable", end)
			break
		}
		visit[u] = true
		m += max
		max = 0
		var temp int
		for _, e := range g.Edges[u] {
			temp = e.To
		}
		for _, e := range g.Edges[u] {
			//fmt.Printf("Edge: %d -> %d (%d)\n", e.From, e.To, e.Weight)
			v := e.To
			w := e.Weight
			if w > max && !visit[v] {
				max = w
				u = v
			}

			if temp == v {
				fmt.Printf("Weight:%v  - %v - ", max, u)
				break
			}
		}
	}

}

var start time.Time

func init() {
	start = time.Now()
}

func main() {
	size := 3
	fmt.Println("main execution started at time", time.Since(start))
	var g *Graph = NewGraph(size)
	/*g.AddEdge(0, 1, 5)
	g.AddEdge(0, 2, 10)
	g.AddEdge(1, 2, 20)
	g.AddEdge(1, 3, 40)
	g.AddEdge(4, 1, 50)
	g.AddEdge(0, 3, 60)
	g.AddEdge(3, 4, 70)*/
	g.AddEdge(0, 1, 5)
	g.AddEdge(0, 2, 6)
	g.AddEdge(1, 0, 5)
	g.AddEdge(1, 2, 4)
	g.AddEdge(2, 1, 4)
	g.AddEdge(2, 0, 6)
	g.adjacentEdgesExample(size, 0, 1)
	fmt.Println("\nmain execution stopped at time", time.Since(start))

}
