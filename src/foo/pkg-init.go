package foo

// resolved var declaration order is a, y, b, x
var x = a + b(a)
var a = 2
var b = func(i int) int { return y * i }
var y = 3
