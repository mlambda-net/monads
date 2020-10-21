package monad

type Future interface {
	SetResult(data interface{})
	Result() (interface{}, error)
}

type future struct {
	pipe chan Mono
}

func (f *future) SetResult(data interface{}) {
	f.pipe <- ToMono(data)
}

func (f *future) Result() (interface{}, error) {
	response := <-f.pipe
	return response.Unwrap()
}

func NewFuture() Future {
	return &future{pipe: make(chan Mono)}
}
