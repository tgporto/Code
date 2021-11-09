/* - Crie um programa que:
   - Atribua um valor int a uma variável
   - Demonstre este valor em decimal, binário e hexadecimal */

package main

import (
	"fmt"
)

var x int

func main() {
	x := 10
	fmt.Printf("%d\n%b\n%#x\n", x, x, x)
}
