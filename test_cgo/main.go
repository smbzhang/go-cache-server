// Package main provides ...
package main

// #include "test.h"
// #cgo LDFLAGS: -L./ -ltest -lstdc++ 
import "C"

import (
	"fmt"
)

func main() {
	C.test()
	fmt.Println("Cgo test !")
}
