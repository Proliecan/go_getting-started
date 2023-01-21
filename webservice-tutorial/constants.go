package main

import (
	"fmt"
)

// constant at package level
const (
	pi     = 3.14 // evaluated at compile time
	first  = iota + 6
	second = 2 << iota // bit shift left by 2 = multiply by 4
	third              // reuses the previous expression (2 << iota)
)

const ( // iota is reset to 0 for each const block
	_      = iota // ignore first value by assigning to blank identifier
	fourth = iota
)

func main() {
	fmt.Println(pi, first, second, third)
	fmt.Println("reset value...", fourth)

	const c = 3 // c is an implicitly typed constant
	fmt.Println(c + 3)
	fmt.Println(c + 1.2) // flexible type interpretation

	const i int = 3 // i is an explicitly typed constant
	fmt.Println(i + 3)
	fmt.Println(float32(i) + 1.2) // type conversion
}
