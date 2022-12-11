package day11

type queue []int

func (s queue) Push(v int) queue {
	return append(s, v)
}

func (queue queue) Pop() (int, queue) {
	element := queue[0] // The first element is the one to be dequeued.
	if len(queue) == 1 {
		var tmp = []int{}
		return element, tmp
	}
	return element, queue[1:] // Slice off the element once it is dequeued.
}
