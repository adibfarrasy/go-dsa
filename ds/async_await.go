package ds

import (
	"math/rand"
	"time"
)

type Future[T any] struct {
	resChan chan Result[T]
}

type Result[T any] struct {
	Val T
	Err error
}

func (f *Future[T]) Set(val T, err error) {
	f.resChan <- Result[T]{val, err}
	close(f.resChan)
}

func (f *Future[T]) Await() (T, error) {
	result := <-f.resChan
	return result.Val, result.Err
}

func Async[T any](fn func() (T, error)) *Future[T] {
	future := &Future[T]{
		resChan: make(chan Result[T]),
	}
	go func(f *Future[T]) {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		f.Set(fn())
	}(future)

	return future
}
