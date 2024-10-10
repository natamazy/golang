package main

import (
	"os"
	"sort"
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

	lines := strings.Split(string(dat), "\n")

	nums := strings.Fields(lines[0])
	rows, _ := strconv.Atoi(nums[0])
	cols, _ := strconv.Atoi(nums[1])

	lines = lines[1:]

	words := make([]string, 0)

	for _, word := range lines {
		wordsInEach := strings.Split(word, "#")
		for _, wordInWord := range wordsInEach {
			if len(wordInWord) > 1 {
				words = append(words, wordInWord)
			}
		}
	}

	for i := 0; i < cols; i++ {
		word := make([]byte, 0)
		for j := 0; j < rows; j++ {
			word = append(word, lines[j][i])
		}

		wordsInEach := strings.Split(string(word), "#")
		for _, valInVal := range wordsInEach {
			if len(valInVal) > 1 {
				words = append(words, valInVal)
			}
		}
	}

	sort.Strings(words)

	write, err := os.Create("output.txt")
	check(err)

	write.Write([]byte(words[0]))
}
