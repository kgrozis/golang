package main

import "fmt"

// the map type:
//   stores unordered elements of same type.  indexed by an arbitrary key value
//   syntax: map[<key type>]<element type>
//   - key is index type, can be any type
// map initialization:
//   uninit map has value of nil.  can't insert without a panic.  can access elements from a nil map (returns 0)
//   syntax: <map_type>{<comman-separated list of key:value pairs>}
var (
	legends   map[int]string
	histogram map[string]int = map[string]int{ // long map init
		"Jan": 100, "Feb": 445, "Mar": 514, "Apr": 233,
		"May": 321, "Jun": 644, "Jul": 113, "Aug": 734,
		"Sep": 553, "Oct": 344, "Nov": 831, "Dec": 312,
	}
	calibration map[float64]bool
	matrix      map[[2][2]int]bool  // map with array key type
	table       = map[string][]int{ // short map init
		"Men":   []int{32, 55, 12, 55, 42, 53},
		"Women": []int{44, 42, 23, 41, 65, 44},
	} // map of string slices
	log map[struct{ name string }]map[string]string // map (with struct key) of map of string
)

// saves into map, blows up upon dup key
func save(store map[string]int, key string, value int) {
	// comm-ok idiom:
	//   boolean value stored in ok
	//   false when value is not there
	val, ok := store[key]
	if !ok {
		store[key] = value
	} else {
		panic(fmt.Sprintf("Slot %d taken", val))
	}
}

// removes an entry, error if not found
func remove(store map[string]int, key string) error {
	_, ok := store[key] // comma-ok idiom
	if !ok {
		return fmt.Errorf("Key not found")
	}
	delete(store, key) // builtin func() that deletes elements of map with given key
	return nil
}

func main() {
	fmt.Printf("%T\n", legends)
	fmt.Printf("%T\n", histogram)
	fmt.Printf("%T\n", calibration)
	fmt.Printf("%T\n", matrix)
	fmt.Printf("%T\n", table)
	fmt.Printf("%T\n", log)

	fmt.Println(histogram)
	fmt.Println(table)

	// making maps:
	//  map can be initialized using make func
	//  inits underlying storage and allows values to be inputed
	//  2nd parameter is capacity of map, but will continue to grow as needed
	hist := make(map[string]int, 6) // make func() to init map
	hist["Jan"] = 100               // inputting/updating a key, value pair
	hist["Feb"] = 445
	hist["Mar"] = 514
	hist["Apr"] = 233
	hist["May"] = 321
	hist["Jun"] = 644
	// maps as parameters:
	//   a map maintains an interanl ptr to it backing storage struct, all updates in a func will be seen in main
	save(hist, "Aug", 734)
	save(hist, "Sep", 553)
	save(hist, "Oct", 344)
	save(hist, "Nov", 831)
	save(hist, "Dec", 312)
	save(hist, "Dec0", 332)
	remove(hist, "Dec0")

	// map traversal:
	//  for...range loop used to walk map values
	//  emits values for both key and elements value
	//  iteration order not guaranteed
	for key, val := range hist {
		adjVal := int(float64(val) * 0.100)
		fmt.Printf("%s (%d):", key, val)
		for i := 0; i < adjVal; i++ {
			fmt.Printf(".")
		}
		fmt.Println()
	}
	fmt.Println(hist)
	fmt.Println(len(hist)) // returns number of entries in map and 0 for uninit map

}
