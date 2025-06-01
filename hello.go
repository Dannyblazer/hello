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
	}

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

type kelvin float64

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}
func realSensor() kelvin {
	return 0
}
func main() {
	var sensor func() kelvin
	sensor = fakeSensor
	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())

}

type kelvin float64

func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%vº K\n", k)
		time.Sleep(time.Second)
	}
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

type getRowFn func(row int) (string, string)

func drawTable(rows int, getrow getRowFn) kelvin {
	name := g
	fmt.Println(name)
	return kelvin(rows)
}

var test getRowFn

func main() {
	measureTemperature(3, fakeSensor) // LEARN DECLARING FUNCTION TYPES
	fmt.Println(test)
	fmt.Println(drawTable(2, getRowFn(test)))
}

// Anonymous Function
var f = func() {
	fmt.Println("Masquerade Done for Annanimos")
}

func main() {
	f()
	func(messages string) {
		fmt.Println(messages)
	}("Welcome My Annanimos people!")
}

type kelvin float64

// sensor function type
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func caliberate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}
func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}
func main() {
	sensor := caliberate(fakeSensor, 5)
	fmt.Println(sensor())
	fmt.Println(sensor())
}


func main() {
	list := [5]string{"Daniel", "Dozie", "Alucard", "Danny", "Debian"}
	for _, lis := range list {
		fmt.Println(lis)
	}
	var board [8][8]string
	board[0][0] = "r"
	board[0][7] = "r"
	for column := range board[1] {
		board[1][column] = "p"
	}

	fmt.Print(board)
}*/

// func main() {
// 	fmt.Println("Welcome to our Pizza App")
// 	fmt.Println("Please Rate Us between 1-5")
// 	reader := bufio.NewReader(os.Stdin)
// 	input, _ := reader.ReadString('\n')
// 	input = strings.TrimSpace(input)

// 	numRating, err := strconv.ParseFloat(input, 64)

// 	if err != nil {
// 		fmt.Println(err)
// 		//panic(err) // Stops Code Exec
// 	} else if numRating <= 5 {
// 		fmt.Println("Thanks for rating ", int(numRating)+1)
// 	} else {
// 		fmt.Println("Rating should be between 1-5")
// 	}

// }
type User struct {
	Name   string
	Email  string
	Gender string
	Age    int
}

func main() {
	lists := [3]string{"Daniel", "Blaze", "Dozie"}
	fmt.Println(lists)

	var slice = []string{"Yello", "Broski", "Thanks"} //Slice
	fmt.Println(slice)

	var news = make([]int, 4) // Slices

	news[0] = 2
	news[1] = 4
	news[2] = 5
	news[3] = 7

	slice = append(slice, "Test", "Okay")
	fmt.Println(slice)

	news = append(news, 23, 12)
	fmt.Println(news)
	index := 4

	news = append(news[:index], news[index+1:]...) // POP using double index
	fmt.Println(news)

	data := make(map[string]string)

	data["py"] = "Python"
	data["js"] = "Javascript"
	data["go"] = "Golang"
	data["jv"] = "Java"

	// Delete data from maps
	fmt.Println(data)

	delete(data, "jv")
	fmt.Println(data)

	me := User{"Daniel", "test@gmail.com", "Male", 27}

	fmt.Printf("Name: %v and Age: %v \n", me.Name, me.Age)
	fmt.Printf("Data: %+v \n", me)
}
