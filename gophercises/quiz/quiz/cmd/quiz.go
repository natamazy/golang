package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

var answer string

func questionAnswer(qa []string, count int) int {
	fmt.Printf("Question %d: %s - ", count, qa[0])
	fmt.Scanln(&answer)

	if strings.EqualFold(strings.TrimSpace(answer), strings.TrimSpace(qa[1])) {
		return 1
	}
	return 0
}

func readFromFile() {
	f, err := os.Open(csvFile)
	if err != nil {
		fmt.Printf("%s: no such file, or user don't have permission to read it\n", csvFile)
		return
	}

	csvReader := csv.NewReader(f)
	var qa []string
	count := 1
	for qa, _ = csvReader.Read(); len(qa) != 0; qa, _ = csvReader.Read() {
		answerCount += questionAnswer(qa, count)
		count++
	}

	f.Close()
}
