// declare package:
//   package identifier can be set any arbitrary value
//   name of pkg is not taken into consideration
//   best practice name same as dir
package volt

// V:
//   returns the voltage of a circuit using Ohm's law
func V(i, r float64) float64 {
	return i * r
}

// Vser:
//   returns the total voltage across a series circuit
func Vser(volts ...float64) (Vtotal float64) {
	for _, v := range volts {
		Vtotal = Vtotal + v
	}
	return
}

// Vpi:
//   returns the voltage when power p and current i are known
func Vpi(p, i float64) float64 {
	return p / i
}
