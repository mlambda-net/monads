package monad

import (
	"github.com/mlambda-net/monads"
)

type Task interface {
	Continue(func(any types.Any) Task) Task
	Result() (types.Any, error)
}

type task struct {
	result chan Mono
}

func (t *task) Continue(fn func(any types.Any) Task) Task {
	r, e := t.Result()
	if e == nil {
		return fn(r)
	}
	return t
}

func (t *task) Result() (types.Any, error) {
	resp := <-t.result
	return resp.Unwrap()
}

func RunTask(fn func() types.Any) Task {
	v := &task{result: make(chan Mono)}
	go func() {
		v.result <- ToMono(fn())
	}()

	return v
}
