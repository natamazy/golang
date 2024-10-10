package main

import (
	"os"
	"strconv"
	"strings"
)

func longestCommonPrefix(a, b string) int {
	minLength := len(a)
	if len(b) < minLength {
		minLength = len(b)
	}

	prefixLen := 0
	for i := 0; i < minLength; i++ {
		if a[i] == b[i] {
			prefixLen++
		} else {
			break
		}
	}
	return prefixLen
}

func main() {
	dat, _ := os.ReadFile("input.txt")

	lines := strings.Fields(string(dat))

	dbSize, _ := strconv.Atoi(lines[0])
	db := lines[1 : dbSize+1]

	searchSize, _ := strconv.Atoi(lines[dbSize+1])
	search := lines[dbSize+2 : dbSize+2+searchSize]
	
	results := make([]int, searchSize)

	for qIndex, searchVal := range search {
		i := 0
		for _, dbVal := range db {
			longestPrefix := longestCommonPrefix(searchVal, dbVal)
			i += longestPrefix + 1

			if dbVal == searchVal {
				break
			}
		}
		results[qIndex] = i
	}

	write, _ := os.Create("output.txt")
	defer write.Close()

	for _, val := range results {
		write.Write([]byte(strconv.Itoa(val) + "\n"))
	}
}
