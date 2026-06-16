package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
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

type location struct {
	Name string  `json:"name"xml:"name"`
	Lat  float64 `json:"latitude"xml:"latitude"`
	Long float64 `json:"longitude"xml:"longitude"`
}

type User struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

type coordinate struct {
	name    string
	d, m, s float64
	h       rune
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func (c coordinate) String() string {
	return fmt.Sprintf("%vº%v'%.1f\" %c", c.d, c.m, c.s, c.h)
}

func newLocation(lat, long coordinate) location {
	return location{Lat: lat.decimal(), Long: long.decimal()}
}

func (u User) Str() {
	fmt.Println("My name is: ", u.Name)
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type world struct {
	radius float64
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

// distance calculation using the Spherical Law of Cosines.
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.Lat))
	s2, c2 := math.Sincos(rad(p2.Lat))
	// Uses the world’s
	clong := math.Cos(rad(p1.Long - p2.Long))
	// radius field
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func distanceBetween(locations []location) float64 {
	if len(locations) < 2 {
		return 0
	}
	var mars = world{radius: 3389.5}
	//var distanceBetween float64
	for i := 0; i < len(locations)-1; i++ {
		for j := i + 1; j < len(locations); j++ {
			distanceBetween := mars.distance(locations[i], locations[j])
			fmt.Printf("Distance between %s and %s is %.2f km\n", locations[i].Name, locations[j].Name, distanceBetween)
		}
	}
	return 0
}

type report struct {
	sol
	temperature
	location
}

type temperature struct {
	low, high celcius
}

type celcius float64

func (t temperature) average() celcius {
	return (t.low + t.high) / 2
}

func (r report) average() celcius {
	return r.temperature.average()
}

type sol int

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		return -days
	}
	return days
}

func (l location) days(l2 location) int {
	// To-do: complicated distance calculation
	return 5
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("Pew\n", int(l))
}

func (l laser) canTalk() bool {
	return true
}

func (loc location) description() string {
	return fmt.Sprintf("Name: %v, Latitude: %v, Longitude: %v", loc.Name, loc.Lat, loc.Long)
}

type check struct {
	laser
}

type gps struct {
	current     location
	destination location
	world
}

func (g gps) distanceToDestination() float64 {
	return g.distance(g.current, g.destination)
}

func (g gps) message() string {
	return fmt.Sprintf("Distance remains %1.2fkm", g.distanceToDestination())
}

type rover struct {
	gps
}

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	p.age++
}

func (p *person) birthday() {
	p.age++
}

func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool { return len(s[i]) < len(s[j]) }
	}
	sort.Slice(s, less)
}

func reclassify(planets *[]string) {
	*planets = (*planets)[0:8]
}

type turtle struct {
	X, Y float64
}

func (p *turtle) moveRight(n float64) {
	fmt.Println("increment X")
	p.X += n
}

func (p *turtle) moveUp(n float64) {
	fmt.Println("increment Y")
	p.Y += n
}

func (p *turtle) moveDown(n float64) {
	fmt.Println("decrement Y")
	p.Y += n
}

func (p *turtle) moveLeft(n float64) {
	fmt.Println("decrement X")
	p.X -= n
}

type laser1 int

func (l *laser1) talk() string {
	return strings.Repeat("pew", int(*l))
}

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main1() {
	// lists := [3]string{"Daniel", "Blaze", "Dozie"}
	// fmt.Println(lists)

	// var slice = []string{"Yello", "Broski", "Thanks"} //Slice
	// fmt.Println(slice)

	// var news = make([]int, 4) // Slices

	// news[0] = 2
	// news[1] = 4
	// news[2] = 5
	// news[3] = 7

	// slice = append(slice, "Test", "Okay")
	// fmt.Println(slice)

	// news = append(news, 23, 12)
	// fmt.Println(news)
	// index := 4

	// news = append(news[:index], news[index+1:]...) // POP using double index
	// fmt.Println(news)

	// data := make(map[string]string)

	// data["py"] = "Python"
	// data["js"] = "Javascript"
	// data["go"] = "Golang"
	// data["jv"] = "Java"

	// // Delete data from maps
	// fmt.Println(data)

	// delete(data, "jv")
	// fmt.Println(data)

	// me := User{"Daniel", "test@gmail.com", "Male", 27}

	// fmt.Printf("Name: %v and Age: %v \n", me.Name, me.Age)
	// fmt.Printf("Data: %+v \n", me)
	// me.Str()
	// fmt.Println("Testing Golang http Request/Reponse")
	// url := "http://20.119.80.237/swagger/"

	// response, err := http.Get(url)

	// if err != nil {
	// 	panic(err)
	// }
	// defer response.Body.Close()
	// number := response.StatusCode
	// character := "Hello there and here's the status code: " + strconv.Itoa(number)
	// fmt.Println(character)
	planets := [...]string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	//fmt.Println(len(planets))
	for i := 0; i < len(planets); i++ {
		fmt.Println(i, planets[i])
	}
	// for _, planet := range planets {
	// fmt.Println(planets)
	// }
	// terrestial := planets[0:4]
	// gasGiants := planets[4:6]
	// iceGiants := planets[6:8]

	// giants := planets[4:8]
	// gas := giants[0:2]
	// ice := giants[2:4]
	// fmt.Println(giants, gas, ice)
	// iceGiantsMarkII := iceGiants
	// iceGiants[1] = "Poseidon"
	// fmt.Println(planets)
	// fmt.Println(iceGiants, iceGiantsMarkII, ice)
	//fmt.Printf("Response Body: %v \n", response.Body)
	dwarfArray := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfSlice := dwarfArray[:]
	//fmt.Printf("Dwarf types: %T \n %T \n", dwarfArray, dwarfSlice)
	//dwarfSlice = append(dwarfArray, "Tinubu", "Toronto")
	fmt.Printf("len: %v and capacity %v", len(dwarfSlice), cap(dwarfSlice))
	fmt.Println("")
	dwarfs1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfs2 := append(dwarfs1, "Orcus")
	dwarfs3 := append(dwarfs2, "Salacia")
	dwarfs4 := append(dwarfs3, "Orcus", "Quaoar", "Sedna")
	dwarfs4[1] = "Pluto!"

	terrestial := planets[0:2:7]
	fmt.Println(cap(terrestial))

	temperatures := map[string]float64{
		"celcius":    0,
		"fahrenheit": 32.0,
		"kelvin":     273.15,
	}
	temperatures["toyota"] = 100.0
	if result, ok := temperatures["toyotas"]; ok {
		fmt.Printf("The value of toyota is %v \n", result)
	}
	g := math.Trunc(20.9/10) * 10
	fmt.Printf("The value of g is %v \n", g)

	user := User{Name: "Daniel", Gender: "Male", Age: 28, Email: "test@gmail.com"}
	user.Name = "Nnaji Daniel"

	user1 := user
	user1.Name = "Changed"

	//fmt.Println("Here's my user struct: ", user)
	// fmt.Printf("%+v\n", user)
	// fmt.Printf("%+v\n", user1)

	users := []User{
		{Name: "John Doe", Age: 30, Gender: "Male", Email: "john@example"},
		{Name: "Jane Smith", Age: 25, Gender: "Female", Email: "jane@example"},
		{Name: "Alice Johnson", Age: 25, Gender: "Female", Email: "alice@gmail.com"},
	}

	for _, user := range users {
		fmt.Printf("%+v\n", user)
		bytes, err := json.Marshal(user)

		exitOnError(err)
		fmt.Println(string(bytes))
	}
	fmt.Println("NEW LINE")

	// locations := []location{
	// 	{Name: "Bradbury Landing", Lat: -4.5895, Long: 137.4417},
	// 	{Name: "Columbia Memorial Station", Lat: -14.5684, Long: 175.472636},
	// 	{Name: "Challenger Memorial Station", Lat: -1.9462, Long: 354.4734},
	// }

	// for _, location := range locations {
	// 	//fmt.Printf("%+v\n", location)
	// 	bytes, err := json.MarshalIndent(location, "", "  ")
	// 	exitOnError(err)
	// 	fmt.Println(string(bytes))
	// }

	// lat := coordinate{4, 35, 22.2, 'S'}
	// long := coordinate{137, 26, 30.1, 'E'}

	// fmt.Printf("Latitude: %v, Longitude: %v\n", lat.decimal(), long.decimal())

	//fmt.Println("%+v\n", users)

	spirit := location{Name: "Spirit", Lat: -14.5684, Long: 175.472636}
	opportunity := location{Name: "Opportunity", Lat: -1.9462, Long: 354.4734}
	var mars = world{radius: 3389.5}
	//var earth = world{radius: 6371.0}

	dist := mars.distance(spirit, opportunity)
	fmt.Printf("%.2f km\n", dist)

	// latitudes := []coordinate{
	// 	{name: "Spirit", d: 14, m: 34, s: 6.2, h: 'S'},
	// 	{name: "Opportunity", d: 1, m: 56, s: 46.3, h: 'S'},
	// 	{name: "Curiosity", d: 4, m: 35, s: 22.2, h: 'S'},
	// 	{name: "Insight", d: 4, m: 30, s: 0.0, h: 'N'},
	// }

	// longitudes := []coordinate{
	// 	{name: "Spirit", d: 175, m: 28, s: 21.5, h: 'E'},
	// 	{name: "Opportunity", d: 354, m: 28, s: 24.4, h: 'E'},
	// 	{name: "Curiosity", d: 137, m: 26, s: 30.1, h: 'E'},
	// 	{name: "Insight", d: 135, m: 54, s: 0.0, h: 'E'},
	// }

	// latitudes := []coordinate{
	// 	{name: "London", d: 14, m: 30, s: 0.0, h: 'N'},
	// 	{name: "Paris", d: 48, m: 51, s: 0.0, h: 'N'},
	// 	{name: "Sharp", d: 5, m: 4, s: 48.0, h: 'S'},
	// 	{name: "Olympus Mons", d: 18, m: 39, s: 0.0, h: 'N'},
	// }

	// longitudes := []coordinate{
	// 	{name: "London", d: 0, m: 8, s: 0.0, h: 'W'},
	// 	{name: "Paris", d: 2, m: 21, s: 0.0, h: 'E'},
	// 	{name: "Sharp", d: 137, m: 51, s: 0.0, h: 'E'},
	// 	{name: "Olympus Mons", d: 226, m: 12, s: 0.0, h: 'E'},
	// }

	// var locations []location

	// for i, lat := range latitudes {
	// 	long := longitudes[i]

	// 	fmt.Printf("%s: Latitude(radius): %.6f, Longitude(radius): %.6f\n", lat.name, lat.decimal(), long.decimal())

	// 	locations = append(locations, location{Name: lat.name, Lat: lat.decimal(), Long: long.decimal()})

	// 	// function to calculate distance between two locations

	// }
	// //fmt.Println(distanceBetween(locations))
	// fmt.Println("Distance between London and Paris:", earth.distance(locations[0], locations[1]), "km")
	// fmt.Println("Distance between Sharp and Olympus Mons:", mars.distance(locations[2], locations[3]), "km")

	report := report{
		sol:         15,
		location:    location{Name: "Bradbury", Lat: -4.5895, Long: 137.4417},
		temperature: temperature{low: -78, high: -1},
	}

	fmt.Println("latitude: ", report.location.Lat)

	fmt.Println("Report average: ", report.average())

	var talker interface {
		canTalk() bool
		talk() string
	}
	var laser laser = 3

	talker = laser
	fmt.Println("Can talk? ", talker.canTalk())

	// Exercise
	gps := gps{
		current:     location{Name: "Bradbury", Lat: -4.5895, Long: 137.4417},
		destination: location{Name: "Elysium", Lat: 4.5, Long: 135.9},
		world:       world{radius: 3389.5},
	}

	fmt.Println(gps.message())

	// Quick check

	// POINTERS
	// var administrator *string
	// pointer := "Hello Pointer"
	// fmt.Println(pointer)
	// administrator = &pointer
	// fmt.Println(&administrator)

	// new_data := administrator

	// *administrator = "Changed data"
	// fmt.Println(pointer)
	// *new_data = "ultimate change"
	// fmt.Println(*&pointer)
	// lighfoot := "New User"
	// administrator = &lighfoot
	// fmt.Println(administrator == new_data)

	worldling := &world{
		radius: 1234,
	}
	worldling.radius = 4321
	fmt.Println(worldling.radius)

	array := &[3]string{"Hello", "Mei", "Chan"}
	fmt.Println(array[0:2])

	slice := &[]string{"test", "exam", "result"}
	fmt.Println((*slice)[:2])

	personal := person{name: "Doxie", superpower: "Speed", age: 27}

	birthday(&personal)
	fmt.Println(personal)
	personal.birthday()
	fmt.Println(personal)

	planetx := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
		"Pluto",
	}

	reclassify(&planetx)
	fmt.Println(planetx)

	pew := laser1(2)
	shout(&pew) // different case here when pointer is used

	// TURTLE PROGRAM
	position := &turtle{X: 0, Y: 0}
	position.moveDown(2)
	position.moveRight(3)
	position.moveLeft(4)
	position.moveUp(5)
	fmt.Println(*position)

	// NIL
	food := []string{"onion", "carrot", "celery"}
	sortStrings(food, nil)
	fmt.Println(food)

	// Note: Empty slice and Nil behave the same way when passed to a method/function
	// Operations/keywords available on both empty slice and Nil slice are append, range, len and cap.

	var soup map[string]int
	fmt.Println(soup == nil)

	measurement, ok := soup["onion"]
	if ok {
		fmt.Println(measurement)
	}

	for ingredients, measurements := range soup {
		fmt.Println(ingredients, measurements)
	}

	var s fmt.Stringer
	fmt.Printf("%#v\n", s)

	//var check bool
	// if !check {

	// }
	e := number{}
	fmt.Println(e)

	arthur := character{
		name:     "Arthur",
		lefthand: nil,
	}

	knight := &character{
		name:     "Merlin",
		lefthand: nil,
	}

	// item1 := item{
	// 	name:        "Excaliber",
	// 	description: "Most powerful weapon in the King Arthur Franchise",
	// }

	item2 := &item{
		name:        "Wand",
		description: "Most powerful wizard weapon",
	}

	arthur.pickup(item2)
	arthur.give(knight)
	fmt.Printf("\nFinal State:\n")
	fmt.Printf("%s is now holding: %s\n", knight.name, knight.lefthand.name)
	if arthur.lefthand == nil {
		fmt.Printf("%s's hand is now empty.\n", arthur.name)
	}

	// ERR Handling
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	// var zero int
	// _ = 10 / zero
	panic("Help me!")
}

type item struct {
	name        string
	description string
}

type character struct {
	name     string
	lefthand *item
}

func (c *character) pickup(to *item) {
	if to == nil {
		fmt.Println("Can't take a nil item")
		return
	}
	c.lefthand = to
}

func (c *character) give(to *character) {
	if c.lefthand == nil {
		fmt.Println("Can't give what you don't have")
		return
	}

	to.lefthand = c.lefthand
	c.lefthand = nil
}

func (n number) String() string {
	if !n.valid {
		return "not set"
	}
	return fmt.Sprintf("%d", n.value)
}

type number struct {
	value int
	valid bool
}
