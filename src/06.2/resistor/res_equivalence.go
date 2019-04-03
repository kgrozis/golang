package resistor

// Rser:
//   return equivalent resistance for ressitors arrange in series in a circuit
//   Exported: captilized so public
func Rser(resists ...float64) (Rtotal float64) {
	for _, r := range resists {
		Rtotal = Rtotal + r
	}
	return
}

// Rpara:
//   return equivalent resistance for resistors arranged in parallel in a circuit
//   Exported: captilized so public
func Rpara(resists ...float64) (Rtotal float64) {
	for _, r := range resists {
		Rtotal = Rtotal + recip(r)
	}
	return
}
