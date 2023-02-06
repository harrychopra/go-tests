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

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < 10000; i++ {
		Repeat('x', 10)
	}
}

func BenchmarkRepeatByConcatenation(b *testing.B) {
	for i := 0; i < 10000; i++ {
		RepeatByConcatenation('x', 10)
	}
}
