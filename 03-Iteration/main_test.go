package iteration

import (
	"testing"

	"github.com/matryer/is"
)

func TestRepeat(t *testing.T) {
	is := is.New(t)
	t.Run("Repeat x 10 times", func(t *testing.T) {
		is.Equal("xxxxxxxxxx", Repeat('x', 10))
	})
}
