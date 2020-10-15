package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var valores string
	fmt.Scanln(valores)
	s := strings.Split(valores, "\n")
	sum := 0
	for _, v := range s {
		aux, _ := strconv.Atoi(v)
		sum += aux
	}
	fmt.Println(sum)
}
