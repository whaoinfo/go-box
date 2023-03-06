package ringqueue

type RingQueue struct {
	maxSize int
	items   []interface{}
	front   int
	rear    int
}

func NewRingQueue(maxSize int) *RingQueue {
	return &RingQueue{
		maxSize: maxSize,
		items:   make([]interface{}, maxSize),
		front:   0,
		rear:    0,
	}
}

func (t *RingQueue) Push(item interface{}) bool {
	if t.IsFull() {
		return false
	}

	t.items[t.rear] = item
	t.rear = (t.rear + 1) % t.maxSize
	return true
}

func (t *RingQueue) Pop() (interface{}, bool) {
	if t.IsEmpty() {
		return nil, false
	}

	retItem := t.items[t.front]
	t.front = (t.front + 1) % t.maxSize
	return retItem, false
}

func (t *RingQueue) IsEmpty() bool {
	return t.front == t.rear
}

func (t *RingQueue) IsFull() bool {
	return t.front == ((t.rear + 1) % t.maxSize)
}

func (t *RingQueue) Length() int {
	return (t.rear - t.front + t.maxSize) % t.maxSize
}
