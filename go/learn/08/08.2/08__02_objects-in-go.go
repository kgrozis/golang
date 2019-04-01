package main

import "fmt"

// objects:
//   object idioms:
//   - data type that stores state: types achieve this
//   - composition: use struct or interface type to express polymorphism
//   - subtype: type defines set of behaviors (methods) that other types may implement
//   - modularity / encapsulation: packages and extensible types
//   - type inheritence: no inheritance thru polymorphism
//   - classes: no notion of a class type

// struct as object:
//   - ability to host methods
//   - ability to be extended via composition
//   - ability to be sub-typed

// object composition:
//   composition over inheritance principle to achieve polymorphism: uses type embedding mechanism supported by struct type
type fuel int

const (
	GASOLINE fuel = iota
	BIO
	ELECTRIC
	JET
)

type vehicle struct {
	make  string
	model string
}

type engine struct {
	fuel   fuel
	thrust int
}

func (e *engine) start() {
	fmt.Println("Engine started.")
}

type truck struct {
	vehicle
	engine
	axels  int
	wheels int
	class  int
}

func (t *truck) drive() {
	fmt.Printf("Truck %s %s, on the go!\n", t.make, t.model)
}

// init function:
//   encapsulates repeatable init logic and enforce validation requirements
func newTruck(mk, mdl string) *truck {
	return &truck{vehicle: vehicle{mk, mdl}}
}

type plane struct {
	vehicle
	engine
	engineCount int
	fixedWings  bool
	maxAltitude int
}

func (p *plane) fly() {
	fmt.Printf("Aircraft %s %s clear for takeoff!\n", p.make, p.model)
}

// init function
//   encapsulates repeatable init logic and enforce validation requirements
func newPlane(mk, mdl string) *plane {
	p := &plane{}
	p.make = mk
	p.model = mdl
	return p
}

func main() {
	t := &truck{ // init var t of type truck using struct literal
		vehicle: vehicle{"Ford", "F150"},
		engine:  engine{GASOLINE + BIO, 700},
		axels:   2,
		wheels:  4,
		class:   1,
	}
	t.start()
	t.drive()

	t2 := newTruck("Chevy", "Silverado 3500")
	t2.axels = 2
	t2.wheels = 6
	t2.class = 3
	t2.start()
	t2.drive()

	p := &plane{}         // init var p of type plane, update struct using dot notation
	p.make = "HondaJet"   // promote vehicle.make field with dot notation
	p.model = "HA-420"    // promote vehicle.model field with dot notation
	p.fuel = JET          // promote engine.fuel field with dot notation
	p.thrust = 2050       // promote engine.thrust field with dot notation
	p.engineCount = 2     // promote engine.engineCount field with dot notation
	p.fixedWings = true   // promote engine.fixedWings field with dot notation
	p.maxAltitude = 43000 // promote engine.maxAltitude field with dot notation
	p.start()             // promote engine.start() method with dot notation
	p.fly()

	p2 := newPlane("Boeing", "747")
	p2.fuel = JET
	p2.thrust = 10000
	p2.engineCount = 4
	p2.fixedWings = true
	p2.maxAltitude = 6000
	p2.start()
	p2.fly()
}
