package main

import "fmt"

type VertexID int
type EdgeID int

func main() {
	gen0 := NewVertexIDGenerator()
	gen1 := NewEdgeIDGenerator()

	fmt.Println("-------------------")
	fmt.Println(gen0(), gen0(), gen0())
	fmt.Println(gen1(), gen1())
	fmt.Println("-------------------")
	NewVertex(gen0())
	NewEdge(gen1())
	fmt.Println("-------------------")

}

func NewVertexIDGenerator() func() VertexID {
	var next VertexID
	return func() VertexID {
		next++
		return next
	}
}

func NewEdgeIDGenerator() func() EdgeID {
	var next EdgeID
	return func() EdgeID {
		next++
		return next
	}
}

func NewVertex(vid VertexID) {
	fmt.Println(vid)
}

func NewEdge(eid EdgeID) {
	fmt.Println(eid)
}
