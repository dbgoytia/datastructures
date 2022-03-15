package fifo

import (
	"errors"
	"fmt"
)

func HelloWorld() string {
	return "hello"
}

type queue struct {
	Entries []int
}

func NewQueue() queue {
	entries := make([]int, 0)
	return queue{Entries: entries}
}

// Adds (stores) an item in the queue
func (q *queue) Enqueue(element int) {
	q.Entries = append(q.Entries, element)
}

// Removes the first element and makes the
// second element the head of the queue
func (q *queue) Dequeue() (int, error) {

	if q.IsEmpty() {
		return 0, errors.New("empty queue")
	}

	res := q.Entries[0]
	new := make([]int, 0)
	for i := 1; i < len(q.Entries); i++ {
		new = append(new, q.Entries[i])
	}
	q.Entries = new
	return res, nil
}

// Returns the element at the front of the queue
func (q *queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	return q.Entries[0], nil

}

// Returns wether the queue is empty or not (bool)
func (q *queue) IsEmpty() bool {
	return len(q.Entries) == 0
}

// Prints a readable queue
func (q *queue) String() {
	for _, element := range q.Entries {
		fmt.Printf("%d ", element)
	}
}
