package main

import "fmt"

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
/*func test1(x, y int64) int64 {
	return x + y
}

func add2(x, y int) int {
	return x + y
}
func main() {
	id := "wow"
	test := (id == "yes" || id == "true" || id == "1")
	fmt.Println(test)

	switch id {
	case "yes", "true", "1":
		test = true
		fmt.Println(test)
	case "no", "false", "0":
		test = false
		fmt.Println(test)
	default:
		fmt.Println("Error: Invalid choice")
	}
	result := test1(2, 4)
	fmt.Println(result)

}

type celcius float32
type check int

func celsiusToFahrenheit(c int) int {
	return (c * 9.0 / 5.0) + 32.0
}
func main() {
	kelvin := 270.0
	result := celsiusToFahrenheit(int(kelvin))
	fmt.Printf("The temperature is %v \n", result)
	var test celcius = 20
	fmt.Println(test)
	}*/

type kelvin float64
type celcius float64
type fahrenheit float64

func celciusToKelvin(c celcius) kelvin {
	return kelvin(c + 273.15)
}

// Converter Methods
func (k kelvin) kelvinToCelcius() celcius { // kelvinToCelsius
	return celcius(k - 273.15)
}

func (c celcius) celciusToKelvin() kelvin { // CelsiusToKelvin
	return kelvin(c + 273.15)
}

func (f fahrenheit) fahrenheitToCelcius() celcius { // FahrenheitToCelsius
	return celcius((f - 32.0) * 5.0 / 9.0)
}

func (c celcius) celciusToFahrenheit() fahrenheit { // celsiusToFahrenheit
	return fahrenheit(c*9.0/5.0) + 32.0
}

func (k kelvin) kelvinToFahrenheit() fahrenheit { //kelvinToFahrenheit
	return fahrenheit((k-273.15)*9/5 + 32)
}

func (f fahrenheit) fahrenheitToKelvin() kelvin { // farenheitToKelvin
	return kelvin((f-32)*5/9 + 273.15)
}

func main() {
	var c celcius = 127
	cTk := c.celciusToKelvin()
	fmt.Printf("Here's the result: %vk\n", cTk)
	kTc := cTk.kelvinToCelcius()
	cTf := kTc.celciusToFahrenheit()
	fTc := cTf.fahrenheitToCelcius()
	kTf := cTk.kelvinToFahrenheit()
	fTk := cTf.fahrenheitToKelvin()

	fmt.Println(kTf, cTk, kTc, cTf, fTc, fTk)

}
