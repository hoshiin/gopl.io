package conv

import "fmt"

// Pond .
type Pond float64

// Kilogram .
type Kilogram float64

// PToK converts pond to kilogram.
func PToK(p Pond) Kilogram {
	return Kilogram(1 / 2.2046 * p)
}

// KToP converts kilogram to pond.
func KToP(k Kilogram) Pond {
	return Pond(k * 2.20462234)
}

func (f Pond) String() string {
	return fmt.Sprintf("%glb", f)
}

func (m Kilogram) String() string {
	return fmt.Sprintf("%gkg", m)
}
