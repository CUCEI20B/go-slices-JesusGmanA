package main

import (
	"fmt"
)

func main() {
	var tamanio, aux int
	fmt.Scan(&tamanio)
	s := make([]int, 0, tamanio)
	for i := 0; i < tamanio; i++ {
		fmt.Scan(&aux)
		s = append(s, aux)
	}
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Println(sum)
}
