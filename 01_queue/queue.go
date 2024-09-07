package queue

// Queue - First In First Out(FIFO)
// In case of Go no need for startIdx and endIdx, because it's uses slices
type Queue[T any] struct {
	// startIdx int
	// endIdx   int
	values []T
	total  int
}

var QUEUE_MAX_LEN = 4

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		total:  0,
		values: make([]T, 0, QUEUE_MAX_LEN),
	}
}

func (q *Queue[T]) Enqueue(n T) {
	q.values = append(q.values, n)
	q.total++
}

func (q *Queue[T]) Dequeue() T {
	e := q.values[0]
	q.values = append(q.values[:0], q.values[0+1:]...)
	q.total--
	return e
}

func (q *Queue[T]) IsEmpty() bool {
	return q.total == 0
}

func (q *Queue[T]) IsFull() bool {
	return q.total == QUEUE_MAX_LEN
}
