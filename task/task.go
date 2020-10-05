
package task

import (
  "github.com/mlambda-net/monads"
  "github.com/mlambda-net/monads/mono"
)

type Task interface {
  Continue(func(any types.Any) Task) Task
  Result() (types.Any, error)
}

type future struct {
  result chan mono.Mono
}

func (f *future) Continue(fn func(any types.Any) Task) Task {
  r, e := f.Result()
  if e == nil {
    return fn(r)
  }
  return f
}

func (f *future) Result() (types.Any, error)  {
  resp :=  <- f.result
  return resp.Unwrap()
}

func RunTask(fn func() types.Any) *future {
  v := &future{ result: make(chan mono.Mono)}
  go func() {
    v.result <- mono.ToMono(fn())
  }()

  return v
}
