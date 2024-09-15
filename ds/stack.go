package ds

type Stack[T any] struct {
	a []T
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		a: []T{},
	}
}

func (s *Stack[T]) Push(i T) {
	s.a = append(s.a, i)
}

func (s *Stack[T]) Pop() (T, error) {
	var res T
	if len(s.a) == 0 {
		return res, ErrPopEmpty
	}
	res = s.a[len(s.a)-1]
	s.a = s.a[:len(s.a)-1]
	return res, nil
}
