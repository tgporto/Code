package main

import "fmt"

func main() {
	for horas := 0; horas <= 12; horas++ {
		fmt.Println("\nHora: ", horas, "\n")
		for x := 0; x < 60; x++ {
			fmt.Println(" ", x)
		}
	}
}
