package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	in, err := os.Open("problems.csv")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	reader := csv.NewReader(in)

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	var correct int
	for _, line := range records {
		fmt.Println(line[0])
		scanner.Scan()
		if scanner.Text() == line[1] {
			correct++
		}
	}

	fmt.Printf("\nYou got %d out of %d answers correctly\n", correct, len(records))
}
