package gs

import (
	"log"

	slice "github.com/gyuho/goraph/gosequence"
)

// Path returns true if there is a path between two Vertices.
func (g *Graph) Path(src, end *Vertex) bool {
	if src == nil || end == nil {
		log.Fatal("Wrong Vertex Passed!")
	}
	src.Color = "gray"

	queue := slice.NewSequence() // Q = ∅
	queue.PushBack(src)          // ENQUEUE(Q, s)

	for queue.Len() != 0 {
		u := queue.PopFront().(*Vertex)
		ovs := u.GetOutVertices()
		for _, vtx := range *ovs {
			if vtx == nil {
				continue
			}
			vt := vtx.(*Vertex)
			if vt == end {
				return true
			}
			if vt.Color == "white" {
				vt.Color = "gray"
				queue.PushBack(vt)
			}
		}
		u.Color = "black"
	}
	return false
}
