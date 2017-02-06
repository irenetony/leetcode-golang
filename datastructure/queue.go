package datastructure

import "errors"

type Queue struct {
	head *Element
	tail *Element
	size int
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Enqueue(v interface{}) {
	newTailElement := &Element{
		value:v,
	}
	if q.tail == nil {
		q.tail = newTailElement
	} else {
		q.tail.next, q.tail = newTailElement, newTailElement
	}
	if q.head == nil {
		q.head = newTailElement
	}
	q.size++
}

func (q *Queue) Dequeue() (value interface{}, err error) {
	if !q.IsEmpty() {
		value = q.head.value
		q.head = q.head.next
		q.size--
		return value, nil
	} else {
		return nil, errors.New("Dequeue an empty queue!")
	}
}

func (q *Queue) Len() int {
	return q.size
}

func NewQueue() *Queue {
	return &Queue{size:0}
}
