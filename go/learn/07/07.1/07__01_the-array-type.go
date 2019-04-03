package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	size     = 1024 * 1024
	variance = 100000
)

// array type:
//   containers for storing seq values of same type that numerically indexed
//   syntax: var <name identifier> [<length>]<element type>
// array initialization:
//   when array variable not explicitly initialized, all elements will be assigned the zero-value for the elements
//   array can be initialized with a composite literal value
//   number of elements in literal must be less than or equal to size declared in the array type
//   multi-dimensional arrays use additional brackets ({}) to nest dimensions
//   [...] can be used to omit length of array
//   syntax: <array_type>{<comma-separated list of element value}
// using arrays:
//   static entities that cannot grow or shrink in size once they are declared with a specified length
//   can reassign values within array
//   use when need to allocate a block of seq memory of predefined size
//   brackets ([]) are used to access values in array
// array as parameters:
//   treated as a single unit and block of memory containing elements.  not a pointer
//   use pointer type when acceessing in a func so value does not need to be copied
var val [100]int = [100]int{44, 72, 12, 55, 64, 1, 4, 90, 13, 54} // array type int and length 100
var msg = [12]rune{0: 'H', 2: 'E', 4: 'L', 6: '0', 8: '!'}
var days [7]string = [7]string{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
}                           // array type string and length 7
var weekdays = [...]string{ // omitting length with [...].  will be 5
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
}
var truth = [256]bool{true}
var histogram = [5]map[string]int{
	map[string]int{"A": 12, "B": 1, "D": 15},
	map[string]int{"man": 1344, "women": 844, "children": 577, "pets": 150},
} // 1 dimension array
var board = [4][2]int{
	{33, 23},
	{62, 2},
	{23, 4},
	{51, 88},
} // multi-dimensional array

// declaring named array types:
//   each declaration must repeat type, use alias array types for declarations
type matrix [2][2][2][2]byte // can reuse complex type without redeclaring
type numbers [size]int
type galaxies [14]string

func printDays(days [7]string) {
	for i, val := range days {
		fmt.Printf("Day %d = %s\n", i, val)
	}
}

func initMat() matrix {
	return matrix{
		{{{4, 4}, {3, 5}}, {{55, 12}, {22, 4}}},
		{{{2, 2}, {7, 9}}, {{43, 0}, {88, 7}}},
	}
}

func max(nums *numbers) int {
	// array traversal:
	//   can be done using the traditional for statement or for...range statement
	//   array iterator stored in val value
	temp := nums[0]
	for _, val := range nums {
		if val > temp {
			temp = val
		}
	}
	return temp
}

func initialize(nums *numbers) {
	rand.Seed(time.Now().UnixNano())
	// array traversal:
	//   can be done using the traditional for statement or for...range statement
	//   array iterator stored in i value
	for i := 0; i < size; i++ {
		nums[i] = rand.Intn(variance)
	}
}

func printGalaxies(names *galaxies) {
	for _, g := range names {
		fmt.Println(g)
	}
}

func main() {
	fmt.Printf("%T: %v\n", val, val)
	fmt.Printf("%T: %v\n", days, days)
	fmt.Printf("%T: %v\n", weekdays, weekdays)
	fmt.Printf("%T: %v\n", truth, truth)
	fmt.Printf("%T: %v\n", histogram, histogram)

	printDays(days)

	var mat1 matrix // using matrix type
	mat1 = initMat()
	fmt.Println(mat1)

	// array length and capacity:
	//   builtin len function return the declard length of an array type
	//   builtin cap function can be used on array to return its capacity
	//   in arrays cap() and len() always return the same value
	seven := [7]string{"grumpy", "sleepy", "bashful"}
	fmt.Println(len(seven), cap(seven))

	// use builtin func new() to init array elements with 0 value and obtain pointer to array
	var nums *numbers = new(numbers)
	// send copy of address (nums -> ptr) instead of 100K array
	initialize(nums)
	fmt.Println(nums)
	fmt.Println("Max num is: ", max(nums))

	// init using & returns a pointer (*galaxies)
	namedGalaxies := &galaxies{
		"Andromeda",
		"Black Eye",
		"Bode's",
		"Cartwheel",
		"Cigar",
		"Comet",
		"Hoag's",
		"Magellanic",
		"Mayall's",
		"Pinwheel",
		"Sombrero",
		"Sunflower",
		"Tadpole",
		"Whirpool",
	}
	fmt.Println("Named Galaxies")
	fmt.Println("--------------")
	printGalaxies(namedGalaxies)
}
