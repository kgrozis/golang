package resistor

// R:
//  returns resistance using Ohm's law when volt v and current i are known
//   Exported: captilized so public
func R(v, i float64) float64 {
	return v / i
}
