package main

import (
	"encoding/json"
	"fmt"
)

// struct type:
//   composite type, serves as a container for other named types known as fields
//   when declaring a struct, must incude all declared field
//   syntax: var <name id> struct {<field declaration set>}
// accessing struct fields:
//   use selector expression (dot notation) to access stored values
//   can chain nested structs into a single selector expression
// struct init
//   pure values, with no additioanal underlying storage structure
//   fields of uninit struct are assinged zero values
//   can use a composite literal to assign:
//     syntax: <struct_type>{<positional or named field values>}
var (
	empty    = struct{}{}
	car      = struct{ make, model string }{make: "Ford", model: "F150"} // composite literal def
	currency = struct {
		name, country string
		code          int
	}{
		"USD", "united States",
		840,
	}
	node = struct {
		edges  []string
		weight int
	}{
		edges: []string{"north", "south", "west"},
	}
	person = struct {
		name    string
		address struct {
			street      string
			city, state string
			postal      string
		}
	}{
		name: "Time the Robot",
		address: struct {
			street      string
			city, state string
			postal      string
		}{street: "Main street"},
	}
)

type address2 struct {
	street      string
	city, state string
	postal      string
}
type person2 struct {
	name     string
	address2 address2
}

type diameter int

type name struct {
	long   string
	short  string
	symbol rune
}

// Field tags:
//   optional string values added to each field declaration
//   hints to tools or other API that consume
//   tags ignored by normal code execution
//   tags wrapped in single quotes ('')
type Person struct {
	Name    string `json:"person_name"`  // json tag annotation
	Title   string `json:"person_title"` // json tag annotation
	Address `json:"person_address_obj"`
}

type Address struct {
	Street string `json:"person_addr_street"` // json tag annotation
	City   string `json:"person_city"`        // json tag annotation
	State  string `json:"person_state"`
	Postal string `json:"person_postal_code"`
}

// anonymous field:
//   define a field with only its type and omitt identifer
//   embeds type directly into struct
//   rules:
//   - name of type become name of field
//   - name of anonymous field cannot clash with other fields
//   - use name when accessing an anonymous field
type planet struct {
	diameter // anonymous field
	name     // anonymous field
	desc     string
}

func makePerson() person2 {
	addr := address2{
		city:   "Goville",
		state:  "Go",
		postal: "12345",
	}
	return person2{
		name:     "kgrozis",
		address2: addr,
	}
}

func nameStr(n name) string {
	return fmt.Sprintf(
		"long name %s short name: %s symbol: %v",
		n.long, n.short, string(n.symbol),
	)
}

func updateName(p person2, name string) {
	p.name = name
}

func updateName2(p *person2, name string) {
	p.name = name
}
func main() {
	fmt.Printf("emtpy %T\n", empty)
	fmt.Printf("car %T\n", car)
	fmt.Printf("currency %T\n", currency)
	fmt.Printf("node %T\n", node)
	fmt.Printf("person %T\n", person)

	p := makePerson()
	fmt.Println(p)

	earth := planet{
		diameter: 7926, // update anonymous field
		name: name{ // update anonymous field
			long:   "Earth",
			short:  "E",
			symbol: '\u2641',
		},
		desc: "Third rock from the Sun",
	}

	jupiter := planet{}
	jupiter.diameter = 88846      // access anonymous field
	jupiter.name.long = "Jupiter" // access anonymous field
	jupiter.name.short = "J"
	jupiter.name.symbol = '\u2643'
	jupiter.desc = "A ball of gas"

	saturn := planet{}
	saturn.diameter = 120536
	// promoted fields
	//   fields in an embedded struct can be promoted to enclosing type
	saturn.long = "Saturn" // promoted, omitt name from selector expr
	saturn.short = "S"     // promoted, omitt name from selector expr
	saturn.symbol = '\u2644'
	saturn.desc = "Slow mover"

	fmt.Printf("Planet %v, diam %d, desc: %s\n", nameStr(earth.name), earth.diameter, earth.desc)
	fmt.Printf("Planet %v, diam %d, desc: %s\n", nameStr(jupiter.name), jupiter.diameter, jupiter.desc)
	fmt.Printf("Planet %v, diam %d, desc: %s\n", nameStr(saturn.name), saturn.diameter, saturn.desc)

	p2 := person2{}
	p2.name = "unknown"
	p2.address2.street = "12345 Main street"
	p2.address2.city, p2.address2.state = "Goville", "Go"
	p2.address2.postal = "12345"
	fmt.Println(p2)
	// struct as paramters
	//   new copy of struct is created and passed in as func paramter
	//   thus value will not update after func call
	updateName(p2, "kgrozis")
	fmt.Println(p2)

	p3 := new(person2)
	p3.name = "unknown"
	p3.address2.street = "12345 Main street"
	p3.address2.city, p2.address2.state = "Goville", "Go"
	p3.address2.postal = "12345"
	fmt.Println(p3)
	// struct as paramters
	//   pass ptr to struct value and name will now update
	updateName2(p3, "kgrozis")
	fmt.Println(p3)

	P := Person{
		Name:  "kgrozis",
		Title: "Author",
		Address: Address{
			Street: "1234 Main street",
			City:   "Goville",
			State:  "Go",
			Postal: "12345",
		},
	}

	b, err := json.Marshal(P)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(P)
	fmt.Println(string(b))
}
