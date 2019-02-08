package main

import "fmt"

// slice type:
//   construct or idiomatic indexed data and more flexible than an array
//   slices do not declare size thus size is not fixed
//   slices can grow unbounded
//   slice types can be nested to create multi-dimensional slices
//   syntax: var <name identifier> []<element_type>
// slice initialization:
//   an uninitialized slice has a value of nil, any attempt to access an uninit'ed slice will cause a panic
//   initial a slice with a composite literal value
//   syntax: var <name identifier> = []<slice type> {comma-seperated list of element values}
// slice representation:
//   - pointer:
//       memory address of the 1st element of the slice.
//       it is a static value.
//       uninit'ed slice ptr is nil
//   - length:
//       number of contigious elements in slice beyond 1st element
//       it is a dynamic value.
//       lenght is always less than capacity.
//       accessing elements beyond length will cause a panic.
//   - capacity:
//       max elements that can be stored starting from 1st element
//       bounded by size of underlying array
// slicing:
//   can slice an existing array or slice using indexing format to express
//   syntax: <slice or array value>[<low index>:<high index>]
//   [:] - omit high and low values slices all elements
//   [:3] - omit low value equals first 3 elements
//   [3:] - omit high value equals omitting first 3 elements
//   [2:4] - omit 1st 2 elements and include element indexes 3,4
// slicing a slice:
//   slicing an existing slice or array doesn't create a new array or slice
//   new slice creates a pointer to underlying array or slice
var (
	image  []byte
	ids    []string = []string{"fe225", "ac144", "3b12"}
	vector          = []float64{12.4, 44, 126, 2, 11.5}
	months          = []string{
		"Jan", "Feb", "Mar", "Apr",
		"May", "Jun", "Jul", "Aug",
		"Sep", "Oct", "Nov", "Dec",
	}
	tables []map[string][]int
	board  = [][]int{
		{712, 5, 344},
		{},
	} // multi-dimensional slice
	graph = [][][][]int{
		{{{44}, {3, 5}}, {{55, 12, 3}, {22, 4}}},
		{{{22, 12, 9, 19}, {7, 9}}, {{43, 0, 44, 12}, {7}}},
	} // multi-dimensional slice
)

// slicing a slice:
//   slicing an existing slice or array doesn't create a new array or slice
//   new slice creates a pointer to underlying array or slice
var (
	halfyr = []string{
		"Jan", "Feb", "Mar",
		"Apr", "May", "Jun",
	}

	q1 = halfyr[:3]
	q2 = halfyr[3:]
)

// slicing an array:
//   existing array becomes the underlying array
var (
	monthsArray [12]string = [12]string{
		"Jan", "Feb", "Mar",
		"Apr", "May", "Jun",
		"Jul", "Aug", "Sep",
		"Oct", "Nov", "Dec",
	}

	halfyrArray = monthsArray[:6]
	q1Array     = halfyrArray[:3]
	q2Array     = halfyrArray[3:6]
	q3Array     = monthsArray[6:9]
	q4Array     = monthsArray[9:]

	// slice expressions with capacity:
	//   includes max capacity of slice in expr
	//   max attribute specifies index value to be used as max capacity of slice
	//   when max isn't specified it last position of underlying array
	//   syntax: <slice_or_array_value[<low index>:<high index>:<max>]
	summer1Array = monthsArray[5:8:8]
)

func print(strs []string) {
	for _, str := range strs {
		fmt.Println(str)
	}
}

// slices as parameters:
//   when func receives a slice as its parameter, the internal ptr of that slice pts to the underlying array of slice
//   any updates of slice within func will be seen by func caller
func scale(factor float64, vector []float64) []float64 {
	// slice traversal using range
	for i, _ := range vector {
		vector[i] *= factor
	}
	return vector
}

func scale2(factor float64, vector []float64) {
	// slice traversal using range
	for i, _ := range vector {
		vector[i] *= factor
	}
}

func join(v1, v2 []float64) []float64 {
	return append(v1, v2...)
}

// copying slices:
//   builtin copy() returns copy of slice with new underlying array
func clone(v []float64) (result []float64) {
	result = make([]float64, len(v), cap(v))
	copy(result, v)
	return
}

func add(v1, v2 []float64) []float64 {
	if len(v1) != len(v2) {
		panic("size mismatch")
	}
	result := make([]float64, len(v1))
	// slice traversal using range
	for i := range v1 {
		result[i] = v1[i] + v2[i]
	}
	return result
}

func sort(str string) string {
	bytes := []byte(str)
	var temp byte
	for i := range bytes {
		for j := i + 1; j < len(bytes); j++ {
			if bytes[j] < bytes[i] {
				temp = bytes[i]
				bytes[i], bytes[j] = bytes[j], temp
			}
		}
	}
	return string(bytes)
}

func printBytes(str string) {
	for i := range str {
		fmt.Println(str[i])
	}
}

func main() {
	// print all variables
	fmt.Printf("Image %T : %v\n", image, image)
	fmt.Printf("Ids %T : %v\n", ids, ids)
	fmt.Printf("Vector %T : %v\n", vector, vector)
	fmt.Printf("Months %T : %v\n", months, months)
	fmt.Printf("Q1 %T : %v\n", q1, q1)
	fmt.Printf("Tables %T : %v\n", tables, tables)
	fmt.Printf("Board %T : %v\n", board, board)
	fmt.Printf("Graph %T : %v\n", graph, graph)

	print(months)
	print(q1)

	fmt.Printf("halfyr [%d,%d]-- %v\n", len(halfyr), cap(halfyr), halfyr)
	fmt.Printf("q1 [%d,%d]-- %v\n", len(q1), cap(q1), q1)
	fmt.Printf("q2 [%d,%d]-- %v\n", len(q2), cap(q2), q2)

	fmt.Printf("monthsArray [%d,%d]--> %v\n", len(monthsArray), cap(monthsArray), monthsArray)
	fmt.Printf("halfyrArray [%d,%d]--> %v\n", len(halfyrArray), cap(halfyrArray), halfyrArray)
	fmt.Printf("q1Array [%d,%d]--> %v\n", len(q1Array), cap(q1Array), q1Array)
	fmt.Printf("q2Array [%d,%d]--> %v\n", len(q2Array), cap(q2Array), q2Array)
	fmt.Printf("q3Array [%d,%d]--> %v\n", len(q3Array), cap(q3Array), q3Array)
	fmt.Printf("q4Array [%d,%d]--> %v\n", len(q4Array), cap(q4Array), q4Array)
	fmt.Printf("summer1Array[%d,%d]--> %v\n", len(summer1Array), cap(summer1Array), summer1Array)

	// making a slice:
	//   slice can be init at runtime using builtin func make()
	//   inits with 0 value and element type
	//   uninit'ed accessing slice will cause a panic
	//   make() takes args - array type, length, capacity
	cal := make([]string, 6, 12)
	cal[0] = "Jan"
	cal[1] = "Feb"
	cal[2] = "Mar"
	cal[3] = "Apr"
	cal[4] = "May"
	cal[5] = "Jun"

	qtr1 := cal[:3]
	qtr2 := cal[3:6]
	fmt.Printf("cal [%d,%d]--> %v\n", len(cal), cap(cal), cal)
	fmt.Printf("qtr1 [%d,%d]--> %v\n", len(qtr1), cap(qtr1), qtr1)
	fmt.Printf("qtr2 [%d,%d]--> %v\n", len(qtr2), cap(qtr2), qtr2)

	// using slices:
	//   access elements by using index notation
	h := []float64{12.5, 18.4, 7.0}
	h[0] = 15 // access index 0 and change to 15
	fmt.Println(h[0])

	h10 := scale(2.0, h)
	fmt.Println(h10)

	h2 := clone(h)
	fmt.Println(add(h, h2))
	fmt.Println("h2", h2)
	scale2(0.5, h2)
	fmt.Println("h2", h2)

	var vector2 []float64     // value nil
	fmt.Println(len(vector2)) // prints 0, no panic
	h3 := make([]float64, 4, 10)
	fmt.Println(len(h3), ",", cap(h3)) // len 4, cap 10

	// appending to slices
	//   builtin append allows slices to grow dynamically
	mths := make([]string, 3, 3)                                     // create slice with cap 3, len 3
	mths = append(mths, "Jan", "Feb", "March", "Apr", "May", "June") // out grows original underlying array and creates new array
	mths = append(mths, []string{"Jul", "Aug", "Sep"}...)
	mths = append(mths, "Oct", "Nov", "Dec")
	fmt.Println(len(mths), cap(mths), mths)

	// strings as slices:
	//   string type is implemented as slice using an underlying array of rune
	//   because strings are array, can use index expressions
	msg := "Bobsayshelloworld!"
	fmt.Println(msg[:3], msg[3:7], msg[7:12], msg[12:17], msg[len(msg)-1:])
	fmt.Println(sort(msg))
	printBytes(msg)
}
