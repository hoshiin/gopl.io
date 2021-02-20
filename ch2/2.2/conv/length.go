package conv

import "fmt"

// Feet .
type Feet float64

// Meter .
type Meter float64

// FToM converts feet to meters
func FToM(f Feet) Meter {
	return Meter(0.3048 * f)
}

// MToF converts meters to feet
func MToF(m Meter) Feet {
	return Feet(m / 0.3048)
}

func (f Feet) String() string {
	return fmt.Sprintf("%gft", f)
}

func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}
