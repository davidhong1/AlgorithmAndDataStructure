package queue

type CircularQueue struct {
	items []string
	head  int
	tail  int
	n     int
}

func InitCircularQueue(n int) *CircularQueue {
	return &CircularQueue{
		items: make([]string, n),
		head:  0,
		tail:  0,
		n:     n,
	}
}

func (c *CircularQueue) Enqueue(item string) bool {
	if (c.tail+1)%c.n == c.head {
		// 循环队列满时的条件，循环队列会浪费一个数组的存储空间。申请10个存储空间的数据，并不能存储10个数据
		return false
	}
	c.items[c.tail] = item
	c.tail = (c.tail + 1) % c.n
	return true
}

func (c *CircularQueue) Dequeue() string {
	if c.tail == c.head {
		// 循环队列为空的条件
		return ""
	}
	ret := c.items[c.head]
	c.head = (c.head + 1) % c.n
	return ret
}
