package stack

// 扩缩容因子，默认2倍
var EXPANSION_FACTOR = 2

type ArrayStack struct {
	items []string `json:"items"` // 数组
	count int      `json:"count"` // 栈中元素个数
	n     int      `json:"n"`     // 栈的大小
}

// 初始化
func InitArrayStack(n int) *ArrayStack {
	return &ArrayStack{
		items: make([]string, n),
		count: 0,
		n:     n,
	}
}

// 入栈操作
func (s *ArrayStack) Push(item string) bool {
	if s.count == s.n && s.count > 0 {
		exp := make([]string, s.n*(EXPANSION_FACTOR-1))
		s.items = append(s.items, exp...)
		s.n += s.n * (EXPANSION_FACTOR - 1)
	}
	s.items[s.count] = item
	s.count++
	return true
}

// 出栈操作
func (s *ArrayStack) Pop() string {
	if s.count == 0 {
		return ""
	}
	ret := s.items[s.count-1]
	s.count--
	return ret
}
