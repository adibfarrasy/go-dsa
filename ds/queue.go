package ds

type Queue[T any] struct {
	arr []T
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		arr: []T{},
	}
}

func (q *Queue[T]) Push(i T) {
	q.arr = append(q.arr, i)
}

func (q *Queue[T]) Pop() (T, error) {
	var res T
	if len(q.arr) == 0 {
		return res, ErrPopEmpty
	}
	res = q.arr[0]
	q.arr = q.arr[1:]
	return res, nil
}

func (q *Queue[T]) Len() int {
	return len(q.arr)
}
