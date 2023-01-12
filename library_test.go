package library_test

import (
	"library"
	"testing"
)

func TestMutex(t *testing.T) {
	b := library.CreateBook()

	t.Run("Change books's name", func(t *testing.T) {
		t.Parallel()

		for i := 0; i < 1000000; i++ {
			b.ChangeName("test")
		}
	})

	t.Run("Get Book's name", func(t *testing.T) {
		t.Parallel()

		for i := 0; i < 1000000; i++ {
			b.Name()
		}
	})
}
