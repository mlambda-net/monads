package mono

import (
  "errors"
  "github.com/stretchr/testify/assert"
  "mlambda.net/monads"
  "testing"
)

func Test_Mono(t *testing.T) {

  m := ToMono(2)

  m = m.Bind(func(any types.Any) Mono {
    return ToMono(2 + any.(int))
  })
  r, f := m.Unwrap()
  assert.Equal(t, r, 4)
  assert.Nil(t, f)
}

func Test_Fail(t *testing.T) {
  m := ToMono(errors.New("not implemented"))

  m = m.Bind(func(any types.Any) Mono {
    assert.Fail(t, "it should not be called")
    return ToMono(nil)
  })

  v, f := m.Unwrap()

  assert.Nil(t, v)
  assert.Error(t, f)
}
