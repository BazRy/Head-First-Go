package paint

import "fmt"

func PaintNeeded(width float64, height float64) (float64, error) {

	if width <= 0 {
		return 0, fmt.Errorf("Width must be greater that zero,  cannot accept %.2f\n", width)
	}
	if height <= 0 {
		return 0, fmt.Errorf("Height must be greater that zero,  cannot accept %.2f\n", height)
	}
	area := width * height
	return area / 10.0, nil
}
