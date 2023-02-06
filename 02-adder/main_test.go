package adder

import (
	"testing"

	"github.com/matryer/is"
)

func TestAdd(t *testing.T) {
	is := is.New(t)
	t.Run("Add 3 and 5", func(t *testing.T) {
		is.Equal(8, Add(3, 5))
	})
}
