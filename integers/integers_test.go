package integers

import "testing"

func TestAdder(t *testing.T) {

	t.Run("should add the 2 given values correctly", func(t *testing.T) {
		actual := Add(2, 2)
		expected := 4

		if actual != expected {
			t.Errorf("expected '%d' but got '%d'", expected, actual)
		}
	})
}
