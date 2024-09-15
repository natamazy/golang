package main

import (
	"fmt"
)

type Animal interface {
	getSound()
}

type Dog struct{}

type Cat struct{}

func (dog Dog) getSound() {
	fmt.Println("woof!")
}

func (cat Cat) getSound() {
	fmt.Println("meow!")
}

func main() {
	someVariables := []interface{}{Dog{}, Cat{}, 15, 6.6}

	for _, val := range someVariables {
		switch typeToCheck := val.(type) {
		case int:
			fmt.Printf("int\n")
		case Dog:
			fmt.Printf("dog\n")
		case Cat:
			fmt.Printf("cat\n")
		default:
			fmt.Printf("%T\n", typeToCheck)
		}
	}
}
