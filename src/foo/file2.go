package foo

// illegal: bar identifier clashes with file1.go
var bar struct {
	x, y int
}

func quux() {
	bar = bar * bar
}
