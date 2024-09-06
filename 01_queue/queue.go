package queue

// Queue - First In First Out(FIFO)
// In case of Go no need for startIdx and endIdx, because it's uses slices
type Queue struct {
	// startIdx int
	// endIdx   int
	values []int
	total  int
}

var QUEUE_MAX_LEN = 4

func NewQueue() *Queue {
	return &Queue{
		total:  0,
		values: make([]int, 0, QUEUE_MAX_LEN),
	}
}

func (q *Queue) Enqueue(n int) {
	q.values = append(q.values, n)
	q.total++
}

func (q *Queue) Dequeue() int {
	e := q.values[0]
	q.values = append(q.values[:0], q.values[0+1:]...)
	q.total--
	return e
}

func (q *Queue) IsEmpty() bool {
	return q.total == 0
}

func (q *Queue) IsFull() bool {
	return q.total == QUEUE_MAX_LEN
}
