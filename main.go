package main

import (
	"fmt"

	"nt.com/user"
)

func main() {
	fmt.Println("START!\n")
	x, _ := user.New("Narek", "Tamazyan", 0)
	var newName string
	fmt.Scanln(&newName)
	x.SetFirstName(&newName)
	x.PrintDetails()
}
