package main

import (
	"errors"
	"fmt"
)

// use error data type to define a new error

func returnError(a, b int) error { // func definition
	if a == b {
		err := errors.New("error in returnError() function") // create error
		return err
	} else {
		return nil // no error to report (aka nil)
	}
}

func main() {
	err := returnError(1, 2)
	if err == nil { // error test
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}

	err = returnError(10, 10)
	if err == nil { // error test
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}

	if err.Error() == "error in returnError() function" { // converts error variable int a string var; can compare an error with a string
		fmt.Println("!!")
	}

}
