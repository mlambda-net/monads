package monad

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Future_NotError(t *testing.T) {

	f := NewFuture()
	go func() {
		u, e := f.Result()
		assert.Equal(t, u, "good")
		assert.Nil(t, e)
	}()

	f.SetResult("good")
}

func Test_Future_Error(t *testing.T) {

	f := NewFuture()
	go func() {
		u, e := f.Result()
		assert.Nil(t, u)
		assert.Error(t, e)
	}()

	f.SetResult(errors.New("bad"))
}
