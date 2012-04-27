package graphs

import (
	"fmt"
)

type AdjacentList struct {
	n   int
	adj map[int]map[int]bool
}

func NewAdjacentList(n0 int) *AdjacentList {
	a := &AdjacentList{n: n0, adj: make(map[int]map[int]bool, n0)}
	for i := 0; i < n0; i++ {
		a.adj[i] = make(map[int]bool)
	}

	return a
}

func (a *AdjacentList) AddEdge(i, j int) {
	a.adj[i][j] = true
}

func (a *AdjacentList) OutEdges(i int) map[int]bool {
	return a.adj[i]
}

func (a *AdjacentList) Bfs(r int) int {
	seen := make(map[int]bool)
	quee := make([]int, 0, 100)
	quee = append(quee, r)
	seen[r] = true

	for len(quee) != 0 {
		i := quee[0]
		quee = quee[1:]
		mapp := a.OutEdges(i)
		for k, v := range mapp {
			if v == true && !seen[k] {
				fmt.Printf(" %d => %d\n", i ,k)
				quee = append(quee, k)
				seen[k] = true
			}
		}
	}
	return len(seen)

}

func NewMeshGraph(n int) *AdjacentList {
	g := NewAdjacentList(n * n)
	for k := 0; k < n*n; k++ {
		if k%n > 0 {
			g.AddEdge(k, k-1)
		}
		if k >= n {
			g.AddEdge(k, k-n)
		}
		if k%n != n-1 {
			g.AddEdge(k, k+1)
		}
		if k < n*(n-1) {
			g.AddEdge(k, k+n)
		}
	}
	return g
}

func main() {
	a := NewMeshGraph(4)
	fmt.Println(a.Bfs(0))
	fmt.Println(a)
}
