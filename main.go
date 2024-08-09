package main

import "fmt"

func main() {
	var input interface{} = 12
	number, ok := input.(int)
	fmt.Println(number, ok)
}
