package maybe

import "mlambda.net/monads"

type Monad interface {
  Bind(func(types.Any) Monad) Monad
}

type just struct {
  value interface{}
}

func (j just) Bind(fn func(value types.Any) Monad) Monad {
  return fn(j.value)
}

type nothing struct {
}

func (n nothing) Bind(_ func(types.Any) Monad) Monad {
  return nothing{}
}

func Unit(value types.Any ) Monad {
  if value != nil {
    return &just{value}
  }

  return nothing{}
}

func Empty() Monad {
  return nothing{}
}

