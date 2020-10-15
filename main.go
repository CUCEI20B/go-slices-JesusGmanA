package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var valores string
	fmt.Scan(&valores)
	split := strings.Split(valores, "\r\n")
	s := split[1:]
	sum := 0
	for _, v := range s {
		aux, _ := strconv.Atoi(v)
		sum += aux
	}
	fmt.Println(sum)
}
