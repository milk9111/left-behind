package engine

type Queue[T any] struct {
	queue []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(t T) {
	q.queue = append(q.queue, t)
}

func (q *Queue[T]) Dequeue() T {
	var t T
	if q.Empty() {
		return t
	}

	t = q.queue[0]

	if len(q.queue) > 1 {
		q.queue = q.queue[1:]
	} else {
		q.Reset()
	}

	return t
}

func (q *Queue[T]) Empty() bool {
	return len(q.queue) == 0
}

func (q *Queue[T]) Len() int {
	return len(q.queue)
}

func (q *Queue[T]) Reset() {
	q.queue = q.queue[:0]
}
