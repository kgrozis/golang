// constants is a value with a literal representation (string, boolean, numbers)
//   can be treated as type or untyped (unlike variables)

package main

import (
	"fmt"
	"math"
)

// typed constants:
//   const <name identifier> type = <value list or initializer expression>
//   must include value
//   cannot have runtime dependency.  must be able to resolve at compile timeconst
const a1, a2 string = "Mastering", "Go"
const b rune = 'G'
const c bool = false
const d int32 = 111009
const e float32 = 2.71828
const f float64 = math.Pi * 2.0e+3
const g complex64 = 5.0i

// untyped constants:
//   const <name identifier> = <value list or initializer expression>
//   no typing percission on storage of constant's value
//   limited use until assigned to var, func, expr
//   can be assigned to many different types
const i = "G is" + " for Go "
const j = 'V'
const k1, k2 = true, !k1
const l = 111*100000 + 9
const m1 = math.Pi * 3.141592
const m2 = 1.414213562373095048801688724209698078569671875376
const m3 = m2 * m2
const m4 = m3 * 1.0e+400
const n = -5.0i * 3

// assign untyped to differing types
var u1 float32 = m2
var u2 float64 = m2

// constant declaration block
// increases readability
const (
	a3, a4 string = "Mastering", "Go"
	b2     rune   = 'G'
	c2     bool   = false
)
const (
	// constant enumeration
	//   Each member becomes untyped integer
	//   iota initializes at zero
	//   remaining member increment by 1
	StarHyperGiant = iota // identifies enumeration
	StarSuperGiant
	StarBrightGiant
	StarGiant
	StarSubGiant
)
const (
	// resets constant enumeration to 0
	// Can declare a type to override default enumeration type of int
	// Can specify any numeric type
	StarDwarf float32 = iota
	StarSubDwarf
	StarWhiteDwarf
)
const (
	// use expression to set value of iota
	// increments base on even numbers
	StarRedDwarf   = 2.0 * iota // value 0 [float64]
	StarBrownDwarf              // value 2 [float64]
)
const (
	_ = iota // value 0, skip value
	square
	triangle
	circle
	_ // value 32, skip value
	rectangle
)

func main() {
	// print default values
	fmt.Println(a1, a2, b)
	fmt.Println(c)
	fmt.Println(d, e, f, g)
	fmt.Println(u1, u2)

	u3 := m2
	fmt.Println(u3)

	u4 := square
	u5 := triangle
	u6 := circle
	u7 := rectangle
	fmt.Println(u4, u5, u6, u7)

}
