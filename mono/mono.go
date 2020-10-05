package mono

import "github.com/mlambda-net/monads"

type Mono interface {
  Bind(func(any types.Any) Mono) Mono
  Unwrap() (types.Any, error)
}

type mono struct {
  value types.Any
}

func (m mono) Bind(fn func(types.Any) Mono) Mono  {
  return fn(m.value)
}

func (m mono) Unwrap() (types.Any, error)  {
  return m.value, nil
}

func ToMono(m types.Any) Mono {
  switch t :=  m.(type) {
  case error: return fail{error: t}
  default:
     return mono{value: t}
  }
}

type fail struct {
  error error
}

func (f fail)Bind(fn func(a types.Any) Mono) Mono  {
  return f
}

func (f fail) Unwrap() (types.Any, error)  {
  return nil, f.error
}



