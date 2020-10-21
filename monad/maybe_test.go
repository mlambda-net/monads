package monad

import (
	"github.com/mlambda-net/monads"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Just(t *testing.T) {
	some := Unit(2)

	some = some.Bind(func(i types.Any) Maybe {
		value := i.(int)
		return Unit(value + 2)
	})

	some.Bind(func(value types.Any) Maybe {
		assert.Equal(t, value, 4)
		return Empty()
	})

}

func Test_Nothing(t *testing.T) {

	val := Empty()

	val.Bind(func(i types.Any) Maybe {
		return Unit(i.(int) + 2)
	})

	val.Bind(func(i types.Any) Maybe {
		assert.Fail(t, "Nothing monad shouldn't be called")
		return Empty()
	})

}
