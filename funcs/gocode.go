package funcs

// #include "../ccode/implementation.c"
import "C"

import "fmt"

// CallGoCode is your typical go func that will be called from another package
func CallGoCode() {
	fmt.Println("CallGoCode():\tThis is go code that lives in another package called by main.")
}

// CallCCode is a wrapper function for a C call
func CallCCode() {
	fmt.Printf("CallCCode():\tThis is a go wrapper function for a call to C code that is defined elsewhere.\n")
	C.c_function_call() // Call C function
	fmt.Printf("CallCCode():\tThis is Go, again.\n")
}
