package main

import (
	"fmt"
	"math"
)

func main() {
	const pi = math.Pi
	var radio float64 = 5.0
	var area = pi * radio * radio

	fmt.Println("El area del circulo es ", area)
}
