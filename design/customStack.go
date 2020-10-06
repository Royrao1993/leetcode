package design

type CustomStack struct {
	Stack []int
	Size  int
	Cap   int
}

func NewCustomStack(maxSize int) *CustomStack {
	return &CustomStack{[]int{}, 0, maxSize}
}

func (s *CustomStack) Push(x int) {
	if s.Size >= s.Cap {
		return
	}
	s.Stack = append(s.Stack, x)
	s.Size++
}

func (s *CustomStack) Pop() int {
	if s.Size <= 0 {
		return -1
	}
	v := s.Stack[len(s.Stack)-1]
	s.Stack = s.Stack[:len(s.Stack)-1]
	s.Size--
	return v
}

func (s *CustomStack) Increment(k, val int) {
	l := len(s.Stack)
	if k > l {
		k = l
	}
	for i := 0; i < k; i++ {
		s.Stack[i] += val
	}
}
