package structs_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("should return correct perimeter for positive values", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		actual := Perimeter(rectangle)
		expected := 40.0

		if actual != expected {
			t.Errorf("got %.2f but wanted %.2f", actual, expected)
		}
	})
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10, Height: 10}, expected: 100.0},
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, expected: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, expected: 314.1592653589793},
		{name: "Circle", shape: Triangle{Base: 12, Height: 6}, expected: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.shape.Area()
			if actual != tt.expected {
				t.Errorf("%#v got %g want %g", tt.shape, actual, tt.expected)
			}
		})
	}
}
