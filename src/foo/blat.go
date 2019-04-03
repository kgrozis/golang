// package name
package foo

import (
	"fmt"
	"foo/bazz"
)

func fooIt() {
	fmt.Println("Foo!")
	bazz.Qux()
}
