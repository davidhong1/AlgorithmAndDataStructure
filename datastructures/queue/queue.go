package queue

type Queue interface {
	Enqueue(item string) bool
	Dequeue() string
}
