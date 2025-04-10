package main

import (
	"fmt"
	"math/rand"
)

func main() {

	if num := rand.Intn(3); num == 0 {
		fmt.Println("Yeah, 0")
	} else if num == 1 {
		fmt.Println("Okay, 1")
	} else {
		fmt.Println("Alright, 3")
	}
}
