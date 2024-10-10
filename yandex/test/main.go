package main

import (
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dat, err := os.ReadFile("input.txt")
	check(err)

	nums := strings.Fields(string(dat))
	num1, _ := strconv.Atoi(nums[0])
	num2, _ := strconv.Atoi(nums[1])

	write, err := os.Create("output.txt")
	check(err)

	write.Write([]byte(strconv.Itoa(num1 + num2)))
}
