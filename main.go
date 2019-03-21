package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Question struct {
	question string
	answer   string
}

func main() {
	fmt.Println("Test Check")

	//Opening CSV
	file := "problems.csv"
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Reading the file
	//
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}
	//Looping through lines
	counter := 0
	for _, line := range lines {
		data := Question{
			question: line[0],
			answer:   line[1],
		}
		fmt.Println("What is ", data.question, " , sir?")
		var storageVariable string
		fmt.Scan(&storageVariable)
		if storageVariable == data.answer {
			continue
		} else {
			counter++

		}

	}
	fmt.Println("You had ", counter, " mistakes")
	correctCounter := len(lines) - counter
	fmt.Println("You had ", correctCounter, " correct answers")
}
