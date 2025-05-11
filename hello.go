package main

import (
	"fmt"
	"unicode/utf8"
)

/*
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

	func main() {
		test := `Nice \n`
		large := new(big.Int)
		large.SetString("24000000000000000", 10)
		secondary := big.NewInt(240030302830802380)

		fmt.Println(large, secondary, test)
	}

	func main() {
		code := '*'
		smile := ''
		last := 'é'

		fmt.Printf("%v %T\n %c %[2]v\n %c %[3]v\n", code, smile, last)
		fmt.Printf("%v is a %[1]T\n", code)
	}

	func main() {
		words := "shalom"
		for i := 0; i < len(words); i++ {
			fmt.Printf("%c \n", words[i])
		}

}

	func main() {
		c := 'g'
		c = c - 'a' + 'A'
		fmt.Printf("%c \n", c)
	}
*/
// refactored to handle int16 and int32 (spanish and russian)
func main() {
	message := "uv vagreangvbany fcnpr fgngvba"

	for _, c := range message {

		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
	words := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(len(words), "bytes")
	fmt.Println(utf8.RuneCountInString(words), "runes")
}

/*
func main() {
	words := "*okay"
	c := words[0]
	fmt.Printf("%c %[1]T \n", c)
}

func main() {
	question := "¿Cómo estás?"
	/*fmt.Println(len(question), "bytes")
	fmt.Println(utf8.RuneCountInString(question), "runes")

	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes \n", c, size)
	for i, c := range question {
		fmt.Printf("%v %c\n", i, c)
	}
}
*/
