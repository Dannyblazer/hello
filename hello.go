package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Coin struct {
	name  string
	value float64
}

func main() {

	bank := 0.0
	money := []Coin{
		{name: "nickel", value: 0.05},
		{name: "dime", value: 0.10},
		{name: "quarter", value: 0.25},
	}
	check := math.Abs(20.0-20.05) > 0.0001
	fmt.Println(check)

	for bank < 20.0 {
		money := money[rand.Intn(len(money))]
		bank += money.value
		//time.Sleep(1 * time.Second)
		fmt.Printf("You Added %s to your bank with new balance %5.2f\n", money.name, bank)
		fmt.Println(check)
	}
	fmt.Printf("You have %5.2f in you bank\n", bank)

}
