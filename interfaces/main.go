package main

import "fmt"

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

func makeSound(animal Animal) {
	animal.getSound()
}

func main() {
	animals := []Animal{&Dog{}, &Cat{}}
	for _, val := range animals {
		makeSound(val)
	}

	var newAnimal Animal = Dog{}
	var newDog any = newAnimal

	assertedType, ok := newDog.(Animal)
	if ok {
		fmt.Printf("My new type is - %T and value is %v\n", assertedType, assertedType)
	} else {
		fmt.Println("Its not dog")
	}
}
