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

type hasName interface {
	printNames()
}

type dynamicStruct struct {
	names []string
}

func (ds *dynamicStruct) printNames() {
	fmt.Println(ds.names)
}

func deepCopyDynamicStruct(ds *dynamicStruct) *dynamicStruct {
	// Create a new slice and copy the data from the original slice
	newNames := make([]string, len(ds.names))
	copy(newNames, ds.names)

	// Return a new instance of dynamicStruct with the copied slice
	return &dynamicStruct{names: newNames}
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

	fmt.Println()

	// var newDyn hasName = &dynamicStruct{[]string{"Narek"}}
	// newAssertedDyn := newDyn.(*dynamicStruct)
	// newAssertedDyn.names[0] = "Alo"
	// newDyn = newDyn.(*dynamicStruct)
	// newAssertedDyn.printNames()
	// newDyn.printNames()

	var newDyn hasName = &dynamicStruct{[]string{"Narek"}}

	newAssertedDyn := deepCopyDynamicStruct(newDyn.(*dynamicStruct))

	newAssertedDyn.names[0] = "Alo"

	newAssertedDyn.printNames()
	newDyn.printNames()
}
