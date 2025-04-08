package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("You're about to enter the cave and there's a sign")
	var command = "Check out the sign"
	var read = strings.Contains(command, "sign")
	fmt.Println("Need to read sign:", read)
	var check = "zax" > "zap"
	fmt.Println(check)
}
