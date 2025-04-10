package main

import (
	"fmt"
)

func main() {
	var checkString = "go west"

	if checkString == "go east" {
		fmt.Println("Head up to the mountains")
	} else if checkString == "go west" {
		fmt.Println("Go to the beach")
	} else {
		fmt.Println("Nah I didn't get that!")
	}
}
