/*- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
- Demonstre estes valores.*/

package main

import (
	"fmt"
)

const (
	_ = 2021 + iota
	x
	y
	w
	z
)

func main() {
	fmt.Println(x, y, w, z)

}
