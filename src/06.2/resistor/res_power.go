package resistor

var Rpi func(float64, float64) float64

//   Not Exported: not captilized so private
// init:
//   special func that takes no args and returns no results
//   encapsulates custom initialization logic that is invoked when pkg is imported
//   pkg can have more than one init func
//   can't access init func at runtime
//   accessed in ordered they appear within each source file
func init() {
	Rpi = func(p, i float64) float64 {
		return p / (i * i)
	}
}

// Rvp:
//   claculates resistance in a circuit when volt v and power p are known
//   Exported: captilized so public
func Rvp(v, p float64) float64 {
	return (v * v) / p
}
