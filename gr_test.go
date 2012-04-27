package graphs

import (
	"fmt"
	"testing"
)

func TestGrapg(t *testing.T) {
	fmt.Println("MEsh")
	g := NewMeshGraph(4)
	fmt.Println("OutEdges")
	fmt.Println(g.OutEdges(2))
	fmt.Println(g.OutEdges(3))
	fmt.Println(g.OutEdges(4))
	fmt.Println(g.OutEdges(5))
	fmt.Println(g.OutEdges(6))
	fmt.Println(g.OutEdges(7))
	fmt.Println(g.OutEdges(8))
	fmt.Println(g.OutEdges(9))
	fmt.Println(g.OutEdges(10))
	fmt.Println(g.OutEdges(11))
	fmt.Println(g.OutEdges(12))
	fmt.Println(g.OutEdges(13))
	fmt.Println(g.OutEdges(14))
	fmt.Println(g.OutEdges(15))
	r := g.Bfs(0)
	fmt.Println(r)
}
