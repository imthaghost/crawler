package queue

import (
	"github.com/enriquebris/goconcurrentqueue"
)

// Queue is an abstract collection of entities
// that are maintained in a sequence and can
// be modified by the addition of entities
// at one end of the sequence and the removal
// of entities from the other end of the sequence
type Queue interface {
	// Enqueue element
	Enqueue(interface{}) error
	// Dequeue element
	Dequeue() (interface{}, error)
}

// NewQueue instantiates a queue
func NewQueue() Queue {
	return &queue{
		q: goconcurrentqueue.NewFIFO(),
	}
}

type queue struct {
	q goconcurrentqueue.Queue
}

// Enqueue enqueues an element. Returns error if queue is locked.
func (f *queue) Enqueue(i interface{}) error {
	err := f.q.Enqueue(i)
	if err != nil {
		return err
	}
	return nil
}

// Dequeue dequeues an element. Returns error if queue is locked or empty.
func (f* queue) Dequeue() (interface{}, error) {
	item, err := f.q.Dequeue()
	if err != nil {
		return nil, err
	}
	return item, nil
}


