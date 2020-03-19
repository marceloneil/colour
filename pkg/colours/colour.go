package colours

import "math"

type Colour struct {
	Name    string
	Article string
	RGB     RGB
	HSV     HSV
}

type RGB struct {
	R uint8 // 8-bit val: [0-255]
	G uint8 // 8-bit val: [0-255]
	B uint8 // 8-bit val: [0-255]
}

type HSV struct {
	H float64 // degrees: [0-1]
	S float64 // percent: [0-1]
	V float64 // percent: [0-1]
}

func (c Colour) Luminance() float64 {
	// https://stackoverflow.com/a/596243
	return math.Sqrt(0.299*math.Pow(float64(c.RGB.R), 2) +
		0.587*math.Pow(float64(c.RGB.G), 2) +
		0.114*math.Pow(float64(c.RGB.B), 2))
}

func (c Colour) HLV(repetitions float64) (uint, uint, uint) {
	// Step sorting
	// https://www.alanzucconi.com/2015/09/30/colour-sorting/
	steppedHue := uint(repetitions * c.HSV.H)
	steppedLuminosity := uint(repetitions * c.Luminance())
	steppedValue := uint(repetitions * c.HSV.V)
	if steppedHue%2 == 1 {
		steppedLuminosity = uint(repetitions) - steppedLuminosity
		steppedValue = uint(repetitions) - steppedValue
	}
	return steppedHue, steppedLuminosity, steppedValue
}
