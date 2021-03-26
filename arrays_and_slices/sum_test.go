package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("should sum up all entries in array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		actual := Sum(numbers)
		expected := 15

		if actual != expected {
			t.Errorf("wanted %d but get %d with given: %v", expected, actual, numbers)
		}

	})

	t.Run("should sum up all entries in array of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

		actual := Sum(numbers)
		expected := 36

		if actual != expected {
			t.Errorf("wanted %d but get %d with given: %v", expected, actual, numbers)
		}

	})
}

func TestSumAll(t *testing.T) {
	t.Run("should sum up all entries in array", func(t *testing.T) {
		actual := SumAll([]int{1, 2, 3, 4, 5}, []int{1, 2, 3})
		expected := []int{15, 6}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("wanted %v but get %v", expected, actual)
		}

	})
}

func TestSumTails(t *testing.T) {
	t.Run("should sum up all entries in array", func(t *testing.T) {
		actual := SumTails([]int{1, 2, 3, 4, 5}, []int{1, 2, 3})
		expected := []int{14, 5}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("wanted %v but get %v", expected, actual)
		}

	})

	t.Run("should sum up all entries in array", func(t *testing.T) {
		actual := SumTails([]int{1, 2, 3, 4, 5}, []int{})
		expected := []int{14, 0}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("wanted %v but get %v", expected, actual)
		}

	})
}
