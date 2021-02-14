package main

import (
	"encoding/json"
	"fmt"
	"github.com/nickhalden/mynicceprogram/helpers"
	"log"
	"sort"
	"time"
)

var b1 = "seven"

// available to other packages/publicly using Caps
type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	Birthdate   time.Time
}

// receiver attached to a type
func (m *User) printFirstName() string {
	return m.FirstName
}

type Animal interface {
	//signature list of functions every type of animal must have
	Says() string
	NumberofLegs() int
}

type Dog struct {
	Name string
}

type Gorrila struct {
	Name  string
	Breed string
}

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	whatToSay, world := saySomething("Hello World")
	log.Println(whatToSay)
	log.Println(saySomething("Something new in this world"))
	var i int
	log.Println(i)
	log.Println(world)

	//Pointers
	var a string
	a = "Green"
	changeUsingPointer(&a)
	log.Println(a)

	var b2 = "six"
	//scope
	log.Println(b1)
	log.Println(b2)

	user := User{
		FirstName:   "Nipun",
		LastName:    "Chawla",
		PhoneNumber: "+1999922233",
	}

	log.Println(user.FirstName, user.LastName, user.Birthdate)

	log.Println(user.printFirstName())

	// maps are immutable and order of insertion
	myMap := make(map[string]string)
	myMap["dog"] = "Oreo"
	log.Println(myMap["dog"])
	log.Println(myMap["nokey"]) //nil

	myOtherMap := make(map[string]int)
	myOtherMap["First"] = 1
	myOtherMap["Second"] = 2
	myOtherMap["Third"] = 3
	myOtherMap["Fourth"] = 4

	log.Println(myOtherMap["First"])
	log.Println(myOtherMap["Nokey"]) //0

	myUserMap := make(map[string]User)
	myUserMap[user.FirstName] = user
	log.Println(myUserMap["Nipun"].FirstName) // maps are immutable and order changes

	// slices
	var fruits []string

	fruits = append(fruits, "orange")
	fruits = append(fruits, "apple")

	log.Println(fruits)
	sort.Strings(fruits) //sorted object

	log.Println(fruits)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers)
	log.Println(numbers[:5])
	log.Println(numbers[:])

	var state bool

	state = true

	if state == true {
		log.Println("found you")
	} else if state == false {
		log.Println("something")
	} else {
		log.Println("false flag")
	}

	// switch
	newVar := "dang"

	switch newVar {

	case "nodang":
		log.Println("")
	case "dang":
		log.Println("dang")
	default:
		log.Println("something else")
	}

	// for loop
	for i := 0; i < 10; i++ {
		log.Println(i)
	}

	for _, x := range fruits {
		log.Println(x)
	}

	for k, v := range myOtherMap {
		log.Println("hello user")
		log.Print(k)
		log.Println(v)
	}

	//interfaces
	d := Dog{Name: "Oreo"}
	PrintInfo(d)

	g := Gorrila{Name: "Chadd"}
	PrintInfo(g)

	var packageTest = helpers.SomeType{SomveVariable: 10}

	log.Println(packageTest)

	// passing info from one part to another or one package to another
	intChan := make(chan int)
	defer close(intChan)
	go calcluateValue(intChan)

	num := <-intChan

	log.Println(num)

	myJson := `
	[
		{
		"first_name": "Nipun",
		"last_name": "Chawla",
		"hair_color": "black",
		"has_dog": true
		},
		{
		"first_name": "Ravi",
		"last_name": "Chawla",
		"hair_color": "black",
		"has_dog": true
		}
	]
	`

	log.Println(myJson)

	var unMarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unMarshalled)
	if err != nil {
		log.Println("Error Unmarshalling", err)
	}
	log.Println(unMarshalled)

	var mySlice []Person
	var m1 Person
	m1.FirstName = "Sangita"
	m1.LastName = "Chawla"
	m1.HairColor = "Orange"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	log.Println(mySlice)
	newJson, err := json.MarshalIndent(mySlice, "", "	")

	log.Println("printing marshalled document")
	if err != nil {
		log.Println(newJson)
	}

	fmt.Println(mySlice)

}

func saySomething(s string) (string, string) {
	return s, "world"
}

func changeUsingPointer(s *string) {
	log.Println(s)
	newValue := "Red"
	*s = newValue
}

// interfaces tied Dog to animal
func (d Dog) Says() string {
	return "woof"
}

func (g Gorrila) NumberofLegs() int {
	return 4
}

func (g Gorrila) Says() string {
	return "grunt"
}

func (d Dog) NumberofLegs() int {
	return 4
}

func PrintInfo(a Animal) {
	log.Println(a.Says())
}

func calcluateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(10)
	intChan <- randomNumber
}
