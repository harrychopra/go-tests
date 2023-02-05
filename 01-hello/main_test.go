package hello

import (
	"testing"

	"github.com/matryer/is"
)

func TestHello(t *testing.T) {
	is := is.New(t)
	t.Run("Name is empty, Language is English", func(t *testing.T) {
		is.Equal("Hello, World!", Hello("", "English"))
	})
	t.Run("Language is French", func(t *testing.T) {
		is.Equal("Bonjour, Harry!", Hello("Harry", "French"))
	})
	t.Run("Language is Spanish", func(t *testing.T) {
		is.Equal("Hola, John!", Hello("John", "Spanish"))
	})
}
