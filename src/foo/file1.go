package foo

var (
	// illegal: bar identifier clashes with file2.go
	bar int = 12
)

// illegal: func name qux already used in pkg
func qux() {
	bar += bar
}
