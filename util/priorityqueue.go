package util

type PriorityQueueElement interface {
	// this.priority greater than ELEMENT
	Gt(interface{}) bool
	// this.priority less than ELEMENT
	Lt(interface{}) bool
	// this.priority equal ELEMENT
	Eq(interface{}) bool
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		elms: make([]PriorityQueueElement, 0),
	}
}

type PriorityQueue struct {
	elms []PriorityQueueElement
}

func (q *PriorityQueue) Push(newElm PriorityQueueElement) {
	ns := make([]PriorityQueueElement, len(q.elms)+1)
	for index, elm := range q.elms {
		if newElm.Lt(elm) {
			continue
		} else {
			if index > 0 {
				copy(ns, q.elms[0:index])
			}
			ns[index] = newElm
			i := index
			for _, v := range q.elms[index:] {
				i++
				ns[i] = v
			}
			break
		}
	}
	if len(q.elms) == 0 {
		ns[0] = newElm
		q.elms = ns
	} else {
		q.elms = ns
	}
}

func (q *PriorityQueue) Pop() PriorityQueueElement {
	if q.Empty() {
		return nil
	} else {
		result := q.elms[0]
		q.elms = q.elms[1:]
		return result
	}
}

func (q *PriorityQueue) Empty() bool {
	if len(q.elms) == 0 {
		return true
	} else {
		return false
	}
}

func (q *PriorityQueue) Contain(v PriorityQueueElement) bool {
	for _, c := range q.elms {
		if c.Eq(v) {
			return true
		}
	}
	return false
}

func (q *PriorityQueue) Size() int {
	return len(q.elms)
}
