package main

import (
	"fmt"

	"github.com/bsandusky/xcompile-rpi-cgo-musl/funcs"
)

func main() {
	fmt.Printf("main():\tHello from Go!\n")
	defer fmt.Printf("main():\tI am done now! Bye-bye.\n")

	funcs.CallGoCode()
	funcs.CallCCode()
}
