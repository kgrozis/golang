package main

import (
	"fmt"
	"math/rand"
)

type Curr struct {
	Currency string
	Name     string
	Country  string
	Number   int
}

var currencies = []Curr{
	Curr{"NOK", "Norwegian Krone", "Norwary", 578},
	Curr{"KES", "Kenyan Shilling", "Kenya", 404},
	Curr{"DZD", "Algerian Dinar", "Algeria", 12},
	Curr{"AUD", "Australian Dollar", "Australia", 36},
	Curr{"MXN", "Mexican Peso", "Mexico", 484},
	Curr{"EUR", "Euro", "Greece", 978},
	Curr{"KHR", "Riel", "Cambodia", 116},
	Curr{"SZL", "Lilangeni", "Swaziland", 748},
	Curr{"GBP", "Pound Sterling", "Isle of Man", 826},
	Curr{"HTG", "Gourde", "Haiti", 332},
	Curr{"BWP", "Pula", "Botswana", 72},
	Curr{"CLP", "Chilean Peso", "Chile", 152},
	Curr{"HKD", "Hong Kong Dollar", "Hong Koong", 344},
	Curr{"HTG", "Gourde", "Haiti", 332},
	Curr{"TRY", "Turkish Lira", "Turkey", 949},
	Curr{"EUR", "Euro", "Belgium", 978},
	Curr{"JMD", "Jamaican Dollar", "Jamaica", 388},
	Curr{"ALL", "Lek", "Albania", 8},
	Curr{"GEL", "Lari", "Georgia", 981},
	Curr{"KFM", "Coromo Franc", "Comoros", 174},
	Curr{"NZD", "New Zeland Dollar", "Tokelau", 554},
}

var list1 = []string{"break", "lake", "go", "right", "strong", "kite", "hello"}

var list2 = []string{"fix", "river", "stop", "left", "weak", "flight", "bye"}

func listCurrs() {
	// For condition:
	//   syntactically similar to a while loop
	//   for followed by a boolean expression
	//   loop continues as long as expression is true
	i := 0
	for i < len(currencies) {
		// Infinite Loop:
		// Excluding boolean statement will make loop run indefinitely
		// for {
		//   statement(s)...
		// }
		fmt.Println(currencies[i])
		// must update value of i or loop becomes infinite
		i++
	}
}

func sortByCurrency() {
	N := len(currencies)
	// traditional for statement
	//   initialization statement, conditional expression, and update statement
	//   omit initialization: for ;  k < 10; ++ {...}
	//   omit update: for k := 0; k < 10; {...}
	for i := 0; i < N-1; i++ {
		currMin := i
		for k := i + 1; k < N; k++ {
			if currencies[k].Number < currencies[currMin].Number {
				currMin = k
			}
		}
		// swap
		if currMin != i {
			temp := currencies[i]
			currencies[i] = currencies[currMin]
			currencies[currMin] = temp
		}
	}
}

func nextPair() (w1, w2 string) {
	pos := rand.Intn(len(list1))
	return list1[pos], list2[pos]
}

func main() {
	fmt.Println("Currencies")
	fmt.Println("----------")
	listCurrs()

	fmt.Println("\nSorted by Currency")
	fmt.Println("------------------")
	sortByCurrency()
	listCurrs()

	rand.Seed(31)
	// Example of for multiple vars in initialization and update statements using
	//   nextPair() func
	for w1, w2 := nextPair(); w1 != "go" && w2 != "stop"; w1, w2 = nextPair() {
		fmt.Printf("Word Pair -> [%s, %s]\n", w1, w2)
	}

	vals := []int{4, 2, 6}
	// the for range
	//   iterates over an array, slice, map, string or channel
	//   range returns 2 values: index and value
	//   index set to nil
	for _, v := range vals {
		v--
		fmt.Println(v)
	}
	fmt.Println(vals)

	for i, v := range vals {
		// update list using index var i
		vals[i] = v - 1
	}
	fmt.Println(vals)

	// omit value iterator in range statement
	for i := range currencies {
		fmt.Printf("%d: %v\n", i, currencies[i])
	}

	// omit index and value iterator in range statement
	for range []int{1, 1, 1, 1, 1} {
		fmt.Println("looping")
	}
}
