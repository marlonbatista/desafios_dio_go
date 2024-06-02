package main

import "fmt"

// desafio: imprima todos os números que são divisíveis por 3 de 0 a 100

func main() {

	for i := 1; i <= 100; i++ {

		if i%5 ==0 && i%3==0 {
			fmt.Println("PimPam")
		} else if i%3 == 0 {
			fmt.Println("Pim")
		} else if i%5 == 0 {
			fmt.Println("Pam")
		} else  {
			fmt.Println(i);
		}
	}
}
