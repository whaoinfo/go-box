package safetyqueue

import "sync"

type IQueue interface {
	Put(item interface{}) bool
	Puts(items ...interface{}) int
	Pop() (interface{}, bool)
	Pops(count int) (retList []interface{})
	IsEmpty() bool
	IsFull() bool
	Length() int
}

func NewSafetyQueue(maxLength int, newInstFunc func(maxLength int) IQueue) *SafetyQueue {
	return &SafetyQueue{
		inst: newInstFunc(maxLength),
	}
}

type SafetyQueue struct {
	lock sync.RWMutex
	inst IQueue
}

func (t *SafetyQueue) Put(item interface{}) bool {
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.inst.Put(item)
}

func (t *SafetyQueue) Puts(items ...interface{}) int {
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.inst.Puts(items...)
}

func (t *SafetyQueue) Pop() (interface{}, bool) {
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.inst.Pop()
}

func (t *SafetyQueue) Pops(count int) []interface{} {
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.inst.Pops(count)
}

func (t *SafetyQueue) IsEmpty() bool {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.inst.IsEmpty()
}

func (t *SafetyQueue) IsFull() bool {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.inst.IsFull()
}

func (t *SafetyQueue) Length() int {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.inst.Length()
}
