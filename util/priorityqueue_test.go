package util

import (
	"fmt"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

type N1 struct {
	value int
}

func (n *N1) Gt(s interface{}) bool {
	node := s.(*N1)
	if n.value > node.value {
		return true
	}
	return false
}

func (n *N1) Lt(s interface{}) bool {
	node := s.(*N1)
	if n.value < node.value {
		return true
	}
	return false
}

func (n *N1) Eq(s interface{}) bool {
	node := s.(*N1)
	if n == node {
		return true
	}
	return false
}

func TestPriorityQueue(t *testing.T) {
	q := NewPriorityQueue()
	assert.Equal(t, 0, q.Size())

	q.Push(&N1{1})
	v2 := &N1{4}
	q.Push(v2)
	v3 := &N1{3}
	q.Push(v3)
	assert.Equal(t, 3, q.Size())
	p := q.Pop().(*N1)
	assert.Equal(t, 2, q.Size())
	assert.Equal(t, 4, p.value)

	assert.True(t, q.Contain(v3))
	assert.False(t, q.Contain(v2))

}

func NewNode(edgesTo []int, edgesCost []int) *Node {
	return &Node{
		edgesTo:   edgesTo,
		edgesCost: edgesCost,
		done:      false,
		cost:      0,
	}
}

type Node struct {
	edgesTo   []int
	edgesCost []int
	done      bool
	cost      int
}

func (n *Node) String() string {
	message := ""
	if n.done {
		message = "done(true)"
	} else {
		message = "done(false)"
	}
	message += fmt.Sprintf(", cost(%d), to(%v), edgecost(%v)", n.cost, n.edgesTo, n.edgesCost)
	return "{" + message + "}"
}

func (n *Node) Cost() int {
	return n.cost
}

func (n *Node) SetCost(cost int) {
	n.cost = cost
}

func (n *Node) Gt(s interface{}) bool {
	node := s.(*Node)
	if n.cost > node.cost {
		return true
	}
	return false
}

func (n *Node) Lt(s interface{}) bool {
	node := s.(*Node)
	if n.cost < node.cost {
		return true
	}
	return false
}

func (n *Node) Eq(s interface{}) bool {
	node := s.(*Node)
	if n == node {
		return true
	}
	return false
}

func TestDijkstra(t *testing.T) {
	nodes := make([]*Node, 6)
	nodes[0] = NewNode([]int{1, 2, 3}, []int{2, 4, 5})
	nodes[1] = NewNode([]int{0, 2, 4}, []int{2, 3, 6})
	nodes[2] = NewNode([]int{0, 1, 3, 4}, []int{4, 3, 2, 2})
	nodes[3] = NewNode([]int{0, 2, 5}, []int{5, 2, 6})
	nodes[4] = NewNode([]int{1, 2, 5}, []int{6, 2, 4})
	nodes[5] = NewNode([]int{}, []int{})

	for _, node := range nodes {
		node.cost = -1
		node.done = false
	}

	start := 0
	last := 5

	nodes[start].cost = 0

	q := NewPriorityQueue()
	q.Push(nodes[start])

	for !q.Empty() {
		doneNode := q.Pop().(*Node)

		doneNode.done = true
		for i := 0; i < len(doneNode.edgesTo); i++ {
			to := doneNode.edgesTo[i]
			cost := doneNode.Cost() + doneNode.edgesCost[i]
			if nodes[to].Cost() < 0 || cost < nodes[to].Cost() {
				nodes[to].SetCost(cost)
				if !q.Contain(nodes[to]) {
					q.Push(nodes[to])
				}
			}
		}
	}

	assert.Equal(t, 10, nodes[last].cost)
}

func BenchmarkDijkstra(b *testing.B) {
	nodes := make([]*Node, 6)
	nodes[0] = NewNode([]int{1, 2, 3}, []int{2, 4, 5})
	nodes[1] = NewNode([]int{0, 2, 4}, []int{2, 3, 6})
	nodes[2] = NewNode([]int{0, 1, 3, 4}, []int{4, 3, 2, 2})
	nodes[3] = NewNode([]int{0, 2, 5}, []int{5, 2, 6})
	nodes[4] = NewNode([]int{1, 2, 5}, []int{6, 2, 4})
	nodes[5] = NewNode([]int{}, []int{})

	for _, node := range nodes {
		node.cost = -1
		node.done = false
	}

	start := 0
	//last := 5

	nodes[start].cost = 0

	q := NewPriorityQueue()
	q.Push(nodes[start])

	for !q.Empty() {
		doneNode := q.Pop().(*Node)

		doneNode.done = true
		for i := 0; i < len(doneNode.edgesTo); i++ {
			to := doneNode.edgesTo[i]
			cost := doneNode.Cost() + doneNode.edgesCost[i]
			if nodes[to].Cost() < 0 || cost < nodes[to].Cost() {
				nodes[to].SetCost(cost)
				if !q.Contain(nodes[to]) {
					q.Push(nodes[to])
				}
			}
		}
	}

	//assert.Equal(t, 10, nodes[last].cost)
}
