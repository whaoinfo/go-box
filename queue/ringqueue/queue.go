package ringqueue

import "sync"

type RingQueue struct {
	lock      sync.Mutex
	maxLength int
	items     []interface{}
	front     int
	rear      int
}

func NewRingQueue(maxLength int) *RingQueue {
	return &RingQueue{
		maxLength: maxLength,
		items:     make([]interface{}, maxLength),
		front:     0,
		rear:      0,
	}
}

func (t *RingQueue) Put(item interface{}) bool {
	if t.IsFull() {
		return false
	}

	t.items[t.rear] = item
	t.rear = (t.rear + 1) % t.maxLength
	return true
}

func (t *RingQueue) Puts(items ...interface{}) int {
	c := 0
	for n, item := range items {
		if !t.Put(item) {
			break
		}
		c = n + 1
	}

	return c
}

func (t *RingQueue) Pop() (interface{}, bool) {
	if t.IsEmpty() {
		return nil, false
	}

	retItem := t.items[t.front]
	t.items[t.front] = nil
	t.front = (t.front + 1) % t.maxLength
	return retItem, true
}

func (t *RingQueue) Pops(count int) (retList []interface{}) {
	for i := 0; i < count; i++ {
		item, ok := t.Pop()
		if !ok {
			return
		}
		retList = append(retList, item)
	}

	return
}

func (t *RingQueue) IsEmpty() bool {
	return t.front == t.rear
}

func (t *RingQueue) IsFull() bool {
	return t.front == ((t.rear + 1) % t.maxLength)
}

func (t *RingQueue) Length() int {
	return (t.rear - t.front + t.maxLength) % t.maxLength
}
