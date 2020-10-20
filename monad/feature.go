package monad

import "github.com/mlambda-net/monads/mono"

type Future interface {
	SetResult(data interface{})
	Result() (interface{}, error)
}

type future struct {
	pipe chan mono.Mono
}

func (f *future) SetResult(data interface{}) {
	f.pipe <- mono.ToMono(data)
}

func (f *future) Result() (interface{}, error) {
	response := <-f.pipe
	return response.Unwrap()
}

func NewFuture() Future {
	return &future{pipe: make(chan mono.Mono)}
}
