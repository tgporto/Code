/*- Crie uma variável de tipo string utilizando uma raw string literal.
- Demonstre-a. */

package main

import (
	"fmt"
)

func main() {
	x := `qualquer
				coisa
		vale`

	fmt.Printf(x)
}
