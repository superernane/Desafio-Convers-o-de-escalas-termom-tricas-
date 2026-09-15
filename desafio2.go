package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			// Caso o número seja múltiplo de 3 e 5 ao mesmo tempo (ex: 15)
			fmt.Println("PinPan")
		} else if i%3 == 0 {
			// Múltiplos de 3 mudam para "Pin"
			fmt.Println("Pin")
		} else if i%5 == 0 {
			// Múltiplos de 5 mudam para "Pan"
			fmt.Println("Pan")
		} else {
			// Caso contrário, exibe o próprio número
			fmt.Println(i)
		}
	}
}
