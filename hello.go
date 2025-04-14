package main

import (
	"fmt"
	"math"
)

func main() {
	var pi64 = math.Pi
	var pi32 float32 = math.Pi

	fmt.Printf("%f\n", pi64)
	fmt.Printf("%.2f\n", pi64)
	fmt.Printf("%5.2f\n", pi64)
	fmt.Println(pi32)
}
