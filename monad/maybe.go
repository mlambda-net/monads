package monad

import "github.com/mlambda-net/monads"

type Maybe interface {
	Bind(func(types.Any) Maybe) Maybe
}

type Just interface {
  Value() interface{}
}

type Nothing interface {
}

type just struct {
	value interface{}
}

func (j *just) Value() interface{} {
  return j.value
}

func (j just) Bind(fn func(value types.Any) Maybe) Maybe {
	return fn(j.value)
}

type nothing struct {
}

func (n nothing) Bind(_ func(types.Any) Maybe) Maybe {
	return nothing{}
}

func Unit(value types.Any) Maybe {
	if value != nil {
		return &just{value}
	}

	return nothing{}
}

func Empty() Maybe {
	return nothing{}
}
