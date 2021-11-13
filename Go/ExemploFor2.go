package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			//roda somente quando o número for impar
			continue
		}
		fmt.Println(i)
	}

}
