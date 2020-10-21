package monad

import (
	"errors"
	types "github.com/mlambda-net/monads"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Task(t *testing.T) {

	f := RunTask(func() types.Any {
		time.Sleep(10 * time.Millisecond)
		return 2
	}).Continue(func(any types.Any) Task {
		return RunTask(func() types.Any {
			return 2 * any.(int)
		})
	})

	r, e := f.Result()

	assert.Equal(t, r, 4)
	assert.Nil(t, e)
}

func Test_Task_Error(t *testing.T) {

	f := RunTask(func() types.Any {
		return errors.New("no data")
	})
	r, e := f.Result()
	assert.Nil(t, r)
	assert.Error(t, e)
}
