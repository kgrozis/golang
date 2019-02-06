package main

// creating programs:
//   a program is also known as a command
//   within a pkg define an entrypoint of execution
//   one source file must be a main package
//   one function must be main() and will be the entry point
//   - main takes no args and returns no values

import (
	"flag"
	"fmt"
	"os"

	"06.2/current"
	"06.2/power"
	"06.2/power/ir"
	"06.2/power/vr"
	res "06.2/resistor"
	"06.2/volt"
)

var (
	op string
	v  float64
	r  float64
	i  float64
	p  float64

	usage = "Usage: ./circ <command> [arguments]\n" +
		"Valid command { V | Vpi | R | Rvp | I | Ivp | P | Pir | Pvr }"
)

// init:
//   used to parse and initialize expected flags
func init() {
	flag.Float64Var(&v, "v", 0.0, "Voltage value (volt)")
	flag.Float64Var(&r, "r", 0.0, "Resistance value (ohms)")
	flag.Float64Var(&i, "i", 0.0, "Current value (amp)")
	flag.Float64Var(&p, "p", 0.0, "Electrical power (watt)")
	flag.StringVar(&op, "op", "V", "Command - one of { V | Vpi | R | Rvp | I | Ivp | P | Pir | Pvr }")
}

// Simple voltage, resistance, current calculator.
// Usage ./circ -op=<operation> arugments
// Example ./circ -op=V -r 12 -i 0.5
// Example ./circ -op=I -v 15 -r 10000
func main() {
	flag.Parse()
	switch op {
	case "V", "v":
		val := volt.V(i, r)
		fmt.Printf("V = %0.2f * %0.2f = %0.2f volts\n", i, r, val)
	case "Vpi", "vpi":
		fmt.Printf("Vpi = %0.2f / %0.2f = %0.2f volts\n", p, i, volt.Vpi(p, i))
	case "R", "r":
		fmt.Printf("R = %0.2f / %0.2f = %0.2f Ohms\n", v, i, res.R(v, i))
	case "I", "i":
		fmt.Printf("I = %0.2f / %0.2f = %0.2f amps\n", v, r, current.I(v, r))
	case "Ivp", "ivp":
		fmt.Printf("Ipv = %0.2f / %0.2f = %0.2f amps\n", p, v, current.Ivp(v, p))
	case "P", "p":
		fmt.Printf("P = %0.2f * %0.2f = %0.2f watts\n", v, i, power.P(v, i))
	case "Pir", "pir":
		fmt.Printf("Pir = %0.2f\u00B2 * %0.2f = %0.2f watts\n", i, r, ir.P(i, r))
	case "Pvr", "pvr":
		fmt.Printf("Pvr = %0.2f\u00B2 / %0.2f = %0.2f watts\n", v, r, vr.P(v, r))
	default:
		fmt.Println(usage)
		os.Exit(1)
	}
}
