package main

import "fmt"

// binding a type to an identifier to create a new named type
//   can be used and referenced where needed
//   syntax: type <name identifier> underllying type name
//   underlying typer can be a builtin (string, numeric ...)
//   can also use composite types ( array, slice, ...)

type fahrenheit float64 // con't enclose in var()
type celsius float64    // all underlying types are float64
type kelvin float64     // can convert between types because floats

func fharToCel(f fahrenheit) celsius { // pass type fharenheit, return type celsius
	return celsius((f - 32) * 5 / 9)
}

func fharToKel(f fahrenheit) kelvin {
	return kelvin((f-32)*5/9 + 273.15)
}

func celToFahr(c celsius) fahrenheit {
	return fahrenheit(c*5/9 + 32)
}

func celToKel(c celsius) kelvin {
	return kelvin(c + 273.15)
}

func main() {
	var c celsius = 32.0
	f := fahrenheit(122)
	fmt.Printf("%.2f \u00b0C = %.2f \u00b0K\n", c, celToKel(c))
	fmt.Printf("%.2f \u00b0F = %.2f \u00b0C\n", f, fharToCel(f))
}
