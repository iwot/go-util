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
	q.elms = push(q.elms, newElm)
}

func push(ps []PriorityQueueElement, elm PriorityQueueElement) []PriorityQueueElement {
	if len(ps) > 1 {
		idx := len(ps) / 2
		if ps[idx-1].Lt(elm) {
			result := ps[0:idx]
			result = append(result, push(ps[idx:], elm)...)
			return result
		} else if ps[idx].Gt(elm) {
			result := push(ps[0:idx], elm)
			result = append(result, ps[idx:]...)
			return result
		} else {
			result := append(ps[0:idx], elm)
			result = append(result, ps[idx:]...)
			return result
		}
	} else if len(ps) == 1 {
		ns := make([]PriorityQueueElement, 2)
		if ps[0].Lt(elm) {
			ns[0], ns[1] = elm, ps[0]
		} else {
			ns[0], ns[1] = ps[0], elm
		}
		return ns
	} else {
		ns := make([]PriorityQueueElement, 1)
		ns[0] = elm
		return ns
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
