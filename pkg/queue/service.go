package queue

import (
	"github.com/enriquebris/goconcurrentqueue"
)

// Service is an abstract collection of entities
// that are maintained in a sequence and can
// be modified by the addition of entities
// at one end of the sequence and the removal
// of entities from the other end of the sequence
type Service interface {
	// Enqueue element
	Enqueue(interface{}) error
	// Dequeue element
	Dequeue() (interface{}, error)
}

type queue struct {
	q goconcurrentqueue.Queue
}


// New instantiates a queue
func New() Service {
	return &queue{
		q: goconcurrentqueue.NewFIFO(),
	}
}

// Enqueue an element. Returns error if queue is locked.
func (f *queue) Enqueue(i interface{}) error {
	err := f.q.Enqueue(i)
	if err != nil {
		return err
	}
	return nil
}

// Dequeue an element. Returns error if queue is locked or empty.
func (f* queue) Dequeue() (interface{}, error) {
	item, err := f.q.Dequeue()
	if err != nil {
		return nil, err
	}
	return item, nil
}


