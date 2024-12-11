package structs // Variável receptora
import (
	"testing"
)

func TestArea(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		received := rectangle.Area()
		expected := 72.0

		if received != expected {
			t.Errorf("expected '%g' but received '%g'", expected, received)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		received := circle.Area()
		expected := 314.1592653589793

		if received != expected {
			t.Errorf("expected '%g' but received '%g'", expected, received)
		}
	})

}
