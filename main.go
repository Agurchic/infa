package main

import (
	"fmt"
	"math"
)

func main() {
	a := 1.35
	b := 0.98
	var x float64
	fmt.Scanln(&x)
	y := math.Sqrt(math.Pow(a*float64(x)+b, 2)) / math.Pow(float64(math.Log10(float64(x))), 2)
	fmt.Printf("%f", y)

}
