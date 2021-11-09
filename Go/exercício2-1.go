//- Escreva um programa que mostre um número em decimal, binário, e hexadecimal.

package main

import (
	"fmt"
)

func main() {
	x := 200
	fmt.Printf("%d, %#x, %b", x, x, x)
}
