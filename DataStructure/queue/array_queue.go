package queue

type ArrayQueue struct {
	items []string
	n     int
	head  int // 队头下标
	tail  int // 队尾下标
}

func InitArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{
		items: make([]string, n),
		n:     n,
		head:  0,
		tail:  0,
	}
}

func (q *ArrayQueue) Enqueue(item string) bool {
	if q.tail == q.n && q.head == 0 {
		return false
	}

	if q.tail == q.n {
		// tail已到数组尾部，迁移数组内容
		for i := 0; i < q.tail-q.head; i++ {
			q.items[i] = q.items[q.head+i]
		}
	}
	q.tail -= q.head
	q.head = 0

	// 队尾添加新内容
	q.items[q.tail] = item
	q.tail++
	return true
}

func (q *ArrayQueue) Dequeue() string {
	if q.head == q.tail {
		return ""
	}
	ret := q.items[q.head]
	q.head++
	return ret
}
