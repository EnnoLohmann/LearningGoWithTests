package iteration

import "testing"

func TestRepeatIt(t *testing.T) {

	t.Run("should repeat the letter 5 times", func(t *testing.T) {
		actual := Repeat("a", 5)
		expected := "aaaaa"

		if actual != expected {
			t.Errorf("exptected %q but got %q", expected, actual)
		}
	})

	t.Run("should repeat the letter 10 times", func(t *testing.T) {
		actual := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		if actual != expected {
			t.Errorf("exptected %q but got %q", expected, actual)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
