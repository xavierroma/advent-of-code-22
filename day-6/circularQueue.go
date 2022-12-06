package day6

type ICircularQueue[T any] struct {
	queue []T
	last  int
}

func CircularQueue[T any](size int) ICircularQueue[T] {
	return ICircularQueue[T]{make([]T, size), 0}
}

func (q *ICircularQueue[T]) enqueue(element T) {
	q.last++
	if q.last >= len(q.queue) {
		q.last = 0
	}
	q.queue[q.last] = element
}
