package maybe

import (
  "github.com/stretchr/testify/assert"
  "mlambda.net/monads"
  "testing"
)

func Test_Just(t *testing.T) {
  some :=  Unit(2)

  some = some.Bind(func(i types.Any) Monad {
    value := i.(int)
    return Unit(value + 2)
  })

  some.Bind(func(value types.Any) Monad {
    assert.Equal(t, value, 4)
    return Empty()
  })

}

func Test_Nothing(t *testing.T)  {

  val := Empty()

  val.Bind(func(i types.Any) Monad {
    return Unit(i.(int) + 2)
  })

  val.Bind(func(i types.Any) Monad {
    assert.Fail(t, "Nothing monad shouldn't be called")
    return Empty()
  })

}

